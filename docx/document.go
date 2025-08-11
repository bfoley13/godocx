package docx

import (
	"encoding/xml"
	"strings"

	"github.com/bfoley13/godocx/internal"
	"github.com/bfoley13/godocx/wml/ctypes"
	"github.com/bfoley13/godocx/wml/stypes"
)

var docAttrs = map[string]string{
	"xmlns:w":      "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
	"xmlns:o":      "urn:schemas-microsoft-com:office:office",
	"xmlns:r":      "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
	"xmlns:v":      "urn:schemas-microsoft-com:vml",
	"xmlns:w10":    "urn:schemas-microsoft-com:office:word",
	"xmlns:wp":     "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing",
	"xmlns:wps":    "http://schemas.microsoft.com/office/word/2010/wordprocessingShape",
	"xmlns:wpg":    "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup",
	"xmlns:mc":     "http://schemas.openxmlformats.org/markup-compatibility/2006",
	"xmlns:wp14":   "http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing",
	"xmlns:w14":    "http://schemas.microsoft.com/office/word/2010/wordml",
	"xmlns:w15":    "http://schemas.microsoft.com/office/word/2012/wordml",
	"mc:Ignorable": "w14 wp14 w15",
}

// This element specifies the contents of a main document part in a WordprocessingML document.
type Document struct {
	// Reference to the RootDoc
	Root *RootDoc

	// Elements
	Background *Background
	Body       *Body

	// Non elements - helper fields
	DocRels      Relationships // DocRels represents relationships specific to the document.
	RID          int
	relativePath string
}

// IncRelationID increments the relation ID of the document and returns the new ID.
// This method is used to generate unique IDs for relationships within the document.
func (doc *Document) IncRelationID() int {
	doc.RID += 1
	return doc.RID
}

// MarshalXML implements the xml.Marshaler interface for the Document type.
func (doc Document) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:document"

	for key, value := range docAttrs {
		attr := xml.Attr{Name: xml.Name{Local: key}, Value: value}
		start.Attr = append(start.Attr, attr)
	}

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if doc.Background != nil {
		if err = doc.Background.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if doc.Body != nil {
		bodyElement := xml.StartElement{Name: xml.Name{Local: "w:body"}}
		if err = e.EncodeElement(doc.Body, bodyElement); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (d *Document) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) (err error) {

	for {
		currentToken, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "body":
				body := NewBody(d.Root)
				if err := decoder.DecodeElement(body, &elem); err != nil {
					return err
				}
				d.Body = body
			case "background":
				bg := NewBackground()
				if err := decoder.DecodeElement(bg, &elem); err != nil {
					return err
				}
				d.Background = bg
			default:
				if err = decoder.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}

}

// AddPageBreak adds a page break to the document by inserting a paragraph containing only a page break.
//
// Returns:
//   - *Paragraph: A pointer to the newly created Paragraph object containing the page break.
//
// Example:
//
//	document := godocx.NewDocument()
//	para := document.AddPageBreak()
func (rd *RootDoc) AddPageBreak() *Paragraph {
	p := rd.AddEmptyParagraph()
	p.AddRun().AddBreak(internal.ToPtr(stypes.BreakTypePage))

	return p
}

// ReplaceAll replaces all occurrences of oldText with newText throughout the document.
// It searches through all paragraphs in the document body, including text within tables.
//
// Parameters:
//   - oldText: The text to search for and replace.
//   - newText: The text to replace oldText with.
//
// Returns:
//   - int: The total number of replacements made.
//
// Example:
//
//	document := godocx.NewDocument()
//	document.AddParagraph("Hello World! This is a test.")
//	replacements := document.ReplaceAll("World", "Universe")
//	// replacements will be 1, and the text becomes "Hello Universe! This is a test."
func (rd *RootDoc) ReplaceAll(oldText, newText string) int {
	if oldText == "" || rd.Document == nil || rd.Document.Body == nil {
		return 0
	}

	replacements := 0

	for _, child := range rd.Document.Body.Children {
		if child.Para != nil {
			replacements += rd.replaceInParagraph(child.Para, oldText, newText)
		} else if child.Table != nil {
			replacements += rd.replaceInTable(child.Table, oldText, newText)
		}
	}

	return replacements
}

// replaceInParagraph replaces text within a single paragraph
func (rd *RootDoc) replaceInParagraph(para *Paragraph, oldText, newText string) int {
	replacements := 0

	for _, child := range para.ct.Children {
		if child.Run != nil {
			replacements += rd.replaceInRun(child.Run, oldText, newText)
		} else if child.Link != nil {
			if child.Link.Run != nil {
				replacements += rd.replaceInRun(child.Link.Run, oldText, newText)
			}
		} else if child.Sdt != nil {
			replacements += rd.replaceInContentControl(child.Sdt, oldText, newText)
		}
	}

	return replacements
}

// replaceInRun replaces text within a single run
func (rd *RootDoc) replaceInRun(run *ctypes.Run, oldText, newText string) int {
	replacements := 0

	for _, child := range run.Children {
		if child.Text != nil {
			before := child.Text.Text
			after := strings.ReplaceAll(before, oldText, newText)
			if before != after {
				child.Text.Text = after
				replacements += strings.Count(before, oldText)
			}
		}
	}

	return replacements
}

// replaceInTable replaces text within all cells of a table
func (rd *RootDoc) replaceInTable(table *Table, oldText, newText string) int {
	replacements := 0

	for _, rowContent := range table.ct.RowContents {
		if rowContent.Row != nil {
			for _, cellContent := range rowContent.Row.Contents {
				if cellContent.Cell != nil {
					replacements += rd.replaceInCell(cellContent.Cell, oldText, newText)
				}
			}
		}
	}

	return replacements
}

// replaceInCell replaces text within a single table cell
func (rd *RootDoc) replaceInCell(cell *ctypes.Cell, oldText, newText string) int {
	replacements := 0

	for _, content := range cell.Contents {
		if content.Paragraph != nil {
			// For paragraphs within cells, we need to work with ctypes.Paragraph directly
			replacements += rd.replaceInCTypeParagraph(content.Paragraph, oldText, newText)
		} else if content.Table != nil {
			// Handle nested tables within cells
			for _, rowContent := range content.Table.RowContents {
				if rowContent.Row != nil {
					for _, cellContent := range rowContent.Row.Contents {
						if cellContent.Cell != nil {
							replacements += rd.replaceInCell(cellContent.Cell, oldText, newText)
						}
					}
				}
			}
		}
	}

	return replacements
}

// replaceInCTypeParagraph replaces text within a ctypes.Paragraph (used within table cells)
func (rd *RootDoc) replaceInCTypeParagraph(para *ctypes.Paragraph, oldText, newText string) int {
	replacements := 0

	for _, child := range para.Children {
		if child.Run != nil {
			replacements += rd.replaceInRun(child.Run, oldText, newText)
		} else if child.Link != nil {
			if child.Link.Run != nil {
				replacements += rd.replaceInRun(child.Link.Run, oldText, newText)
			}
		}
	}

	return replacements
}

// replaceInContentControl replaces text within a content control
func (rd *RootDoc) replaceInContentControl(sdt *ctypes.StructuredDocumentTag, oldText, newText string) int {
	replacements := 0

	if sdt.Content == nil {
		return replacements
	}

	for _, child := range sdt.Content.Children {
		if child.Run != nil {
			replacements += rd.replaceInRun(child.Run, oldText, newText)
		} else if child.Paragraph != nil {
			replacements += rd.replaceInCTypeParagraph(child.Paragraph, oldText, newText)
		} else if child.Table != nil {
			// Handle nested tables within content controls
			for _, rowContent := range child.Table.RowContents {
				if rowContent.Row != nil {
					for _, cellContent := range rowContent.Row.Contents {
						if cellContent.Cell != nil {
							replacements += rd.replaceInCell(cellContent.Cell, oldText, newText)
						}
					}
				}
			}
		}
	}

	return replacements
}

// ReplaceFields replaces field codes throughout the document based on a map of field codes to replacement values.
// It searches through all paragraphs in the document body, including within tables and content controls,
// to find Word fields and replace their results with the provided values.
//
// Word fields follow the pattern: begin → instrText → separate → result → end
// This function identifies fields by their instruction text and replaces the result text.
//
// Parameters:
//   - fieldMap: A map where keys are field codes (e.g., "MERGEFIELD Name", "PAGE") and values are replacement text.
//
// Returns:
//   - int: The total number of field replacements made.
//
// Example:
//
//	document := godocx.NewDocument()
//	// Document contains fields like: { MERGEFIELD Name } and { PAGE }
//	fieldMap := map[string]string{
//	    "MERGEFIELD Name": "John Doe",
//	    "PAGE": "42",
//	}
//	replacements := document.ReplaceFields(fieldMap)
func (rd *RootDoc) ReplaceFields(fieldMap map[string]string) int {
	if len(fieldMap) == 0 || rd.Document == nil || rd.Document.Body == nil {
		return 0
	}

	replacements := 0

	for _, child := range rd.Document.Body.Children {
		if child.Para != nil {
			replacements += rd.replaceFieldsInParagraph(child.Para, fieldMap)
		} else if child.Table != nil {
			replacements += rd.replaceFieldsInTable(child.Table, fieldMap)
		}
	}

	return replacements
}

// replaceFieldsInParagraph replaces fields within a single paragraph
// This function handles fields that may span across multiple runs
func (rd *RootDoc) replaceFieldsInParagraph(para *Paragraph, fieldMap map[string]string) int {
	replacements := 0

	// Collect all runs from the paragraph for field processing
	var runs []*ctypes.Run
	var runIndices []int

	for i, child := range para.ct.Children {
		if child.Run != nil {
			runs = append(runs, child.Run)
			runIndices = append(runIndices, i)
		} else if child.Link != nil {
			if child.Link.Run != nil {
				runs = append(runs, child.Link.Run)
				runIndices = append(runIndices, i)
			}
		} else if child.Sdt != nil {
			replacements += rd.replaceFieldsInContentControl(child.Sdt, fieldMap)
		}
	}

	// Process fields across all runs in the paragraph
	if len(runs) > 0 {
		replacements += rd.replaceFieldsAcrossRuns(runs, fieldMap)
	}

	return replacements
}

// replaceFieldsInTable replaces fields within all cells of a table
func (rd *RootDoc) replaceFieldsInTable(table *Table, fieldMap map[string]string) int {
	replacements := 0

	for _, rowContent := range table.ct.RowContents {
		if rowContent.Row != nil {
			for _, cellContent := range rowContent.Row.Contents {
				if cellContent.Cell != nil {
					replacements += rd.replaceFieldsInCell(cellContent.Cell, fieldMap)
				}
			}
		}
	}

	return replacements
}

// replaceFieldsInCell replaces fields within a single table cell
func (rd *RootDoc) replaceFieldsInCell(cell *ctypes.Cell, fieldMap map[string]string) int {
	replacements := 0

	for _, content := range cell.Contents {
		if content.Paragraph != nil {
			replacements += rd.replaceFieldsInCTypeParagraph(content.Paragraph, fieldMap)
		} else if content.Table != nil {
			for _, rowContent := range content.Table.RowContents {
				if rowContent.Row != nil {
					for _, cellContent := range rowContent.Row.Contents {
						if cellContent.Cell != nil {
							replacements += rd.replaceFieldsInCell(cellContent.Cell, fieldMap)
						}
					}
				}
			}
		}
	}

	return replacements
}

// replaceFieldsInCTypeParagraph replaces fields within a ctypes.Paragraph (used within table cells)
// This function handles fields that may span across multiple runs
func (rd *RootDoc) replaceFieldsInCTypeParagraph(para *ctypes.Paragraph, fieldMap map[string]string) int {
	replacements := 0

	// Collect all runs from the paragraph for field processing
	var runs []*ctypes.Run

	for _, child := range para.Children {
		if child.Run != nil {
			runs = append(runs, child.Run)
		} else if child.Link != nil {
			if child.Link.Run != nil {
				runs = append(runs, child.Link.Run)
			}
		}
	}

	// Process fields across all runs in the paragraph
	if len(runs) > 0 {
		replacements += rd.replaceFieldsAcrossRuns(runs, fieldMap)
	}

	return replacements
}

// replaceFieldsInContentControl replaces fields within a content control
func (rd *RootDoc) replaceFieldsInContentControl(sdt *ctypes.StructuredDocumentTag, fieldMap map[string]string) int {
	replacements := 0

	if sdt.Content == nil {
		return replacements
	}

	for _, child := range sdt.Content.Children {
		if child.Run != nil {
			replacements += rd.replaceFieldsInRun(child.Run, fieldMap)
		} else if child.Paragraph != nil {
			replacements += rd.replaceFieldsInCTypeParagraph(child.Paragraph, fieldMap)
		} else if child.Table != nil {
			for _, rowContent := range child.Table.RowContents {
				if rowContent.Row != nil {
					for _, cellContent := range rowContent.Row.Contents {
						if cellContent.Cell != nil {
							replacements += rd.replaceFieldsInCell(cellContent.Cell, fieldMap)
						}
					}
				}
			}
		}
	}

	return replacements
}

// replaceFieldsAcrossRuns processes fields that may span across multiple runs within a paragraph
func (rd *RootDoc) replaceFieldsAcrossRuns(runs []*ctypes.Run, fieldMap map[string]string) int {
	replacements := 0

	// Track field state across all runs
	fieldState := "none" // none, begin, instrText, separate, result
	var currentFieldCode strings.Builder
	var resultElements []fieldResultElement // stores references to text elements containing field results

	// Iterate through all runs and their children to find field sequences
	for runIdx, run := range runs {
		for childIdx, child := range run.Children {
			if child.FldChar != nil {
				switch child.FldChar.FldCharType.Val {
				case stypes.FldCharTypeBegin:
					currentFieldCode.Reset()
					resultElements = []fieldResultElement{}
					fieldState = "begin"

					if child.FldChar.FFData != nil && child.FldChar.FFData.TextInput != nil && child.FldChar.FFData.TextInput.Default != nil {
						if _, exists := fieldMap[strings.TrimSpace(child.FldChar.FFData.TextInput.Default.Val)]; exists {
							currentFieldCode.WriteString(child.FldChar.FFData.TextInput.Default.Val)
						}
					}
				case stypes.FldCharTypeSeparate:
					fieldState = "separate"
					resultElements = []fieldResultElement{}
				case stypes.FldCharTypeEnd:
					if fieldState == "separate" {
						// We have a complete field, try to replace it
						fieldCode := strings.TrimSpace(currentFieldCode.String())
						if replacement, exists := fieldMap[fieldCode]; exists {
							// Replace all result text elements with the replacement
							rd.replaceFieldResultAcrossRuns(resultElements, replacement)
							replacements++
						}
					}
					fieldState = "none"
					currentFieldCode.Reset()
					resultElements = []fieldResultElement{}
				}
			} else if child.InstrText != nil && (fieldState == "begin" || fieldState == "instrText") {
				fieldState = "instrText"
				currentFieldCode.WriteString(child.InstrText.Text)
			} else if child.Text != nil && fieldState == "separate" {
				resultElements = append(resultElements, fieldResultElement{
					run:      run,
					runIdx:   runIdx,
					childIdx: childIdx,
				})
			}
		}
	}

	return replacements
}

// fieldResultElement stores a reference to a text element within a field result
type fieldResultElement struct {
	run      *ctypes.Run
	runIdx   int
	childIdx int
}

// replaceFieldResultAcrossRuns replaces field result text elements across multiple runs
func (rd *RootDoc) replaceFieldResultAcrossRuns(elements []fieldResultElement, replacement string) {
	if len(elements) == 0 {
		return
	}

	// Replace the first text element with the replacement text
	if elements[0].childIdx < len(elements[0].run.Children) &&
		elements[0].run.Children[elements[0].childIdx].Text != nil {
		elements[0].run.Children[elements[0].childIdx].Text.Text = replacement
	}

	// Clear the remaining text elements
	for i := 1; i < len(elements); i++ {
		elem := elements[i]
		if elem.childIdx < len(elem.run.Children) &&
			elem.run.Children[elem.childIdx].Text != nil {
			elem.run.Children[elem.childIdx].Text.Text = ""
		}
	}
}

// replaceFieldsInRun replaces fields within a single run by analyzing field patterns
func (rd *RootDoc) replaceFieldsInRun(run *ctypes.Run, fieldMap map[string]string) int {
	replacements := 0

	// Find field sequences within this run
	fieldState := "none" // none, begin, instrText, separate, result
	var currentFieldCode strings.Builder
	var resultIndices []int // indices of text elements that contain field results

	for i, child := range run.Children {
		if child.FldChar != nil {
			switch child.FldChar.FldCharType.Val {
			case stypes.FldCharTypeBegin:
				fieldState = "begin"
				currentFieldCode.Reset()
				resultIndices = []int{}
			case stypes.FldCharTypeSeparate:
				fieldState = "separate"
				resultIndices = []int{}
			case stypes.FldCharTypeEnd:
				if fieldState == "separate" {
					// We have a complete field, try to replace it
					fieldCode := strings.TrimSpace(currentFieldCode.String())
					if replacement, exists := fieldMap[fieldCode]; exists {
						// Replace all result text elements with the replacement
						rd.replaceFieldResult(run, resultIndices, replacement)
						replacements++
					}
				}
				fieldState = "none"
				currentFieldCode.Reset()
				resultIndices = []int{}
			}
		} else if child.InstrText != nil && (fieldState == "begin" || fieldState == "instrText") {
			fieldState = "instrText"
			currentFieldCode.WriteString(child.InstrText.Text)
		} else if child.Text != nil && fieldState == "separate" {
			resultIndices = append(resultIndices, i)
		}
	}

	return replacements
}

// replaceFieldResult replaces the text elements at the given indices with the replacement text
func (rd *RootDoc) replaceFieldResult(run *ctypes.Run, textIndices []int, replacement string) {
	if len(textIndices) == 0 {
		return
	}

	// Replace the first text element with the replacement text
	if textIndices[0] < len(run.Children) && run.Children[textIndices[0]].Text != nil {
		run.Children[textIndices[0]].Text.Text = replacement
	}

	// Clear the remaining text elements (in reverse order to avoid index shifting)
	for i := len(textIndices) - 1; i > 0; i-- {
		idx := textIndices[i]
		if idx < len(run.Children) && run.Children[idx].Text != nil {
			run.Children[idx].Text.Text = ""
		}
	}
}
