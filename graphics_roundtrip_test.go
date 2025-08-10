package godocx

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bfoley13/godocx/docx"
	"github.com/bfoley13/godocx/wml/ctypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGraphicsRoundtrip tests that graphics formatting is preserved during
// document load and save operations
func TestGraphicsRoundtrip(t *testing.T) {
	// Use complexdoc.docx which should contain graphics
	inputFile := filepath.Join("testdata", "complexdoc.docx")
	outputFile := filepath.Join("testdata", "complexdoc_roundtrip_test.docx")
	
	// Ensure the test file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		t.Skip("complexdoc.docx test file not found")
		return
	}
	
	// Clean up output file at the end
	defer func() {
		if err := os.Remove(outputFile); err != nil && !os.IsNotExist(err) {
			t.Logf("Failed to clean up output file: %v", err)
		}
	}()

	// Load the document
	doc, err := OpenDocument(inputFile)
	require.NoError(t, err, "Failed to open document")
	require.NotNil(t, doc, "Document should not be nil")

	// Count graphics elements in the original document
	originalGraphicsCount := countGraphicsElements(doc)
	t.Logf("Original document contains %d graphics elements", originalGraphicsCount)

	// Save the document
	err = doc.SaveTo(outputFile)
	require.NoError(t, err, "Failed to save document")

	// Verify output file was created
	_, err = os.Stat(outputFile)
	require.NoError(t, err, "Output file should exist")

	// Load the saved document
	savedDoc, err := OpenDocument(outputFile)
	require.NoError(t, err, "Failed to open saved document")
	require.NotNil(t, savedDoc, "Saved document should not be nil")

	// Count graphics elements in the saved document
	savedGraphicsCount := countGraphicsElements(savedDoc)
	t.Logf("Saved document contains %d graphics elements", savedGraphicsCount)

	// Verify graphics count is preserved
	assert.Equal(t, originalGraphicsCount, savedGraphicsCount, 
		"Graphics count should be preserved during roundtrip")
	
	// If we have graphics, verify their properties are preserved
	if originalGraphicsCount > 0 {
		verifyGraphicsProperties(t, doc, savedDoc)
	}
}

// countGraphicsElements counts the number of graphics elements in a document
func countGraphicsElements(doc *docx.RootDoc) int {
	count := 0
	
	if doc.Document == nil || doc.Document.Body == nil {
		return count
	}
	
	// Count drawings in paragraphs
	for _, child := range doc.Document.Body.Children {
		if child.Para != nil {
			count += countGraphicsInParagraph(child.Para)
		} else if child.Table != nil {
			count += countGraphicsInTable(child.Table)
		}
	}
	
	return count
}

// countGraphicsInParagraph counts graphics elements in a paragraph
func countGraphicsInParagraph(para *docx.Paragraph) int {
	count := 0
	
	for _, child := range para.GetCT().Children {
		if child.Run != nil {
			for _, runChild := range child.Run.Children {
				if runChild.Drawing != nil {
					count += len(runChild.Drawing.Inline) + len(runChild.Drawing.Anchor)
				}
			}
		}
	}
	
	return count
}

// countGraphicsInTable counts graphics elements in a table
func countGraphicsInTable(table *docx.Table) int {
	count := 0
	
	for _, rowContent := range table.GetCT().RowContents {
		if rowContent.Row != nil {
			for _, cellContent := range rowContent.Row.Contents {
				if cellContent.Cell != nil {
					for _, content := range cellContent.Cell.Contents {
						if content.Paragraph != nil {
							count += countGraphicsInCTypeParagraph(content.Paragraph)
						} else if content.Table != nil {
							count += countGraphicsInNestedTable(content.Table)
						}
					}
				}
			}
		}
	}
	
	return count
}

// countGraphicsInCTypeParagraph counts graphics in a ctypes.Paragraph
func countGraphicsInCTypeParagraph(para *ctypes.Paragraph) int {
	count := 0
	
	for _, child := range para.Children {
		if child.Run != nil {
			for _, runChild := range child.Run.Children {
				if runChild.Drawing != nil {
					count += len(runChild.Drawing.Inline) + len(runChild.Drawing.Anchor)
				}
			}
		}
	}
	
	return count
}

// verifyGraphicsProperties performs a more detailed comparison of graphics properties
func verifyGraphicsProperties(t *testing.T, original, saved *docx.RootDoc) {
	// This is a simplified verification - in a real scenario you'd want to
	// compare specific properties like sizes, positions, etc.
	
	originalGraphics := extractGraphicsInfo(original)
	savedGraphics := extractGraphicsInfo(saved)
	
	assert.Equal(t, len(originalGraphics), len(savedGraphics), 
		"Number of graphics should match")
		
	for i := 0; i < len(originalGraphics) && i < len(savedGraphics); i++ {
		orig := originalGraphics[i]
		saved := savedGraphics[i]
		
		// Compare basic properties
		assert.Equal(t, orig.Type, saved.Type, "Graphics type should match")
		assert.Equal(t, orig.Width, saved.Width, "Graphics width should match")
		assert.Equal(t, orig.Height, saved.Height, "Graphics height should match")
	}
}

// GraphicsInfo holds basic information about a graphics element
type GraphicsInfo struct {
	Type   string // "inline" or "anchor"
	Width  uint64
	Height uint64
}

// extractGraphicsInfo extracts basic graphics information from a document
func extractGraphicsInfo(doc *docx.RootDoc) []GraphicsInfo {
	var graphics []GraphicsInfo
	
	if doc.Document == nil || doc.Document.Body == nil {
		return graphics
	}
	
	for _, child := range doc.Document.Body.Children {
		if child.Para != nil {
			graphics = append(graphics, extractGraphicsFromParagraph(child.Para)...)
		} else if child.Table != nil {
			graphics = append(graphics, extractGraphicsFromTable(child.Table)...)
		}
	}
	
	return graphics
}

// extractGraphicsFromParagraph extracts graphics info from a paragraph
func extractGraphicsFromParagraph(para *docx.Paragraph) []GraphicsInfo {
	var graphics []GraphicsInfo
	
	for _, child := range para.GetCT().Children {
		if child.Run != nil {
			for _, runChild := range child.Run.Children {
				if runChild.Drawing != nil {
					// Extract inline graphics
					for _, inline := range runChild.Drawing.Inline {
						graphics = append(graphics, GraphicsInfo{
							Type:   "inline",
							Width:  inline.Extent.Width,
							Height: inline.Extent.Height,
						})
					}
					
					// Extract anchor graphics
					for _, anchor := range runChild.Drawing.Anchor {
						graphics = append(graphics, GraphicsInfo{
							Type:   "anchor",
							Width:  anchor.Extent.Width,
							Height: anchor.Extent.Height,
						})
					}
				}
			}
		}
	}
	
	return graphics
}

// extractGraphicsFromTable extracts graphics info from a table
func extractGraphicsFromTable(table *docx.Table) []GraphicsInfo {
	var graphics []GraphicsInfo
	
	for _, rowContent := range table.GetCT().RowContents {
		if rowContent.Row != nil {
			for _, cellContent := range rowContent.Row.Contents {
				if cellContent.Cell != nil {
					for _, content := range cellContent.Cell.Contents {
						if content.Paragraph != nil {
							graphics = append(graphics, extractGraphicsFromCTypeParagraph(content.Paragraph)...)
						} else if content.Table != nil {
							graphics = append(graphics, extractGraphicsFromNestedTable(content.Table)...)
						}
					}
				}
			}
		}
	}
	
	return graphics
}

// extractGraphicsFromCTypeParagraph extracts graphics from a ctypes.Paragraph
func extractGraphicsFromCTypeParagraph(para *ctypes.Paragraph) []GraphicsInfo {
	var graphics []GraphicsInfo
	
	for _, child := range para.Children {
		if child.Run != nil {
			for _, runChild := range child.Run.Children {
				if runChild.Drawing != nil {
					// Extract inline graphics
					for _, inline := range runChild.Drawing.Inline {
						graphics = append(graphics, GraphicsInfo{
							Type:   "inline",
							Width:  inline.Extent.Width,
							Height: inline.Extent.Height,
						})
					}
					
					// Extract anchor graphics
					for _, anchor := range runChild.Drawing.Anchor {
						graphics = append(graphics, GraphicsInfo{
							Type:   "anchor",
							Width:  anchor.Extent.Width,
							Height: anchor.Extent.Height,
						})
					}
				}
			}
		}
	}
	
	return graphics
}