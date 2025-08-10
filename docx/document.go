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
