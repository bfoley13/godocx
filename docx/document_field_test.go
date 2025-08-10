package docx

import (
	"testing"

	"github.com/bfoley13/godocx/wml/ctypes"
	"github.com/bfoley13/godocx/wml/stypes"
	"github.com/stretchr/testify/assert"
)

func TestRootDoc_ReplaceFields(t *testing.T) {
	t.Run("Empty field map returns 0", func(t *testing.T) {
		doc := NewRootDoc()
		doc.Document = &Document{Body: &Body{}}
		
		result := doc.ReplaceFields(map[string]string{})
		assert.Equal(t, 0, result)
	})

	t.Run("Nil document returns 0", func(t *testing.T) {
		doc := NewRootDoc()
		
		result := doc.ReplaceFields(map[string]string{"test": "value"})
		assert.Equal(t, 0, result)
	})

	t.Run("Replace simple field in paragraph", func(t *testing.T) {
		doc := createDocumentWithField("PAGE", "1")
		
		fieldMap := map[string]string{
			"PAGE": "42",
		}
		
		result := doc.ReplaceFields(fieldMap)
		assert.Equal(t, 1, result)
		
		// Verify the field result was replaced
		para := doc.Document.Body.Children[0].Para
		run := para.ct.Children[0].Run
		
		// Find the text element that should contain "42"
		var resultText string
		for _, child := range run.Children {
			if child.Text != nil && child.Text.Text != "" {
				resultText = child.Text.Text
				break
			}
		}
		assert.Equal(t, "42", resultText)
	})

	t.Run("Replace MERGEFIELD", func(t *testing.T) {
		doc := createDocumentWithField("MERGEFIELD Name", "«Name»")
		
		fieldMap := map[string]string{
			"MERGEFIELD Name": "John Doe",
		}
		
		result := doc.ReplaceFields(fieldMap)
		assert.Equal(t, 1, result)
		
		// Verify the field result was replaced
		para := doc.Document.Body.Children[0].Para
		run := para.ct.Children[0].Run
		
		var resultText string
		for _, child := range run.Children {
			if child.Text != nil && child.Text.Text != "" {
				resultText = child.Text.Text
				break
			}
		}
		assert.Equal(t, "John Doe", resultText)
	})

	t.Run("Replace multiple fields", func(t *testing.T) {
		doc := NewRootDoc()
		doc.Document = &Document{
			Root: doc,
			Body: &Body{
				Children: []DocumentChild{
					{Para: createParagraphWithField("PAGE", "1")},
					{Para: createParagraphWithField("MERGEFIELD Name", "«Name»")},
				},
			},
		}
		
		fieldMap := map[string]string{
			"PAGE":           "42",
			"MERGEFIELD Name": "John Doe",
		}
		
		result := doc.ReplaceFields(fieldMap)
		assert.Equal(t, 2, result)
	})

	t.Run("Field not in map is not replaced", func(t *testing.T) {
		doc := createDocumentWithField("DATE", "1/1/2024")
		
		fieldMap := map[string]string{
			"PAGE": "42",
		}
		
		result := doc.ReplaceFields(fieldMap)
		assert.Equal(t, 0, result)
		
		// Verify the original result text is unchanged
		para := doc.Document.Body.Children[0].Para
		run := para.ct.Children[0].Run
		
		var resultText string
		for _, child := range run.Children {
			if child.Text != nil && child.Text.Text != "" {
				resultText = child.Text.Text
				break
			}
		}
		assert.Equal(t, "1/1/2024", resultText)
	})
}

func TestRootDoc_ReplaceFields_InTable(t *testing.T) {
	t.Run("Replace field in table cell", func(t *testing.T) {
		doc := createDocumentWithFieldInTable("PAGE", "1")
		
		fieldMap := map[string]string{
			"PAGE": "42",
		}
		
		result := doc.ReplaceFields(fieldMap)
		assert.Equal(t, 1, result)
	})
}

// Helper functions for creating test documents with fields

func createDocumentWithField(fieldCode, resultText string) *RootDoc {
	doc := NewRootDoc()
	doc.Document = &Document{
		Root: doc,
		Body: &Body{
			Children: []DocumentChild{
				{Para: createParagraphWithField(fieldCode, resultText)},
			},
		},
	}
	return doc
}

func createParagraphWithField(fieldCode, resultText string) *Paragraph {
	para := &Paragraph{
		ct: ctypes.Paragraph{
			Children: []ctypes.ParagraphChild{
				{Run: createRunWithField(fieldCode, resultText)},
			},
		},
	}
	return para
}

func createRunWithField(fieldCode, resultText string) *ctypes.Run {
	run := &ctypes.Run{
		Children: []ctypes.RunChild{
			// Field begin
			{
				FldChar: &ctypes.FieldChar{
					FldCharType: &ctypes.GenSingleStrVal[stypes.FldCharType]{
						Val: stypes.FldCharTypeBegin,
					},
				},
			},
			// Field instruction text
			{
				InstrText: &ctypes.Text{
					Text: fieldCode,
				},
			},
			// Field separate
			{
				FldChar: &ctypes.FieldChar{
					FldCharType: &ctypes.GenSingleStrVal[stypes.FldCharType]{
						Val: stypes.FldCharTypeSeparate,
					},
				},
			},
			// Field result text
			{
				Text: &ctypes.Text{
					Text: resultText,
				},
			},
			// Field end
			{
				FldChar: &ctypes.FieldChar{
					FldCharType: &ctypes.GenSingleStrVal[stypes.FldCharType]{
						Val: stypes.FldCharTypeEnd,
					},
				},
			},
		},
	}
	return run
}

func createDocumentWithFieldInTable(fieldCode, resultText string) *RootDoc {
	doc := NewRootDoc()
	
	cell := &ctypes.Cell{
		Contents: []ctypes.TCBlockContent{
			{
				Paragraph: &ctypes.Paragraph{
					Children: []ctypes.ParagraphChild{
						{Run: createRunWithField(fieldCode, resultText)},
					},
				},
			},
		},
	}
	
	row := &ctypes.Row{
		Contents: []ctypes.TRCellContent{
			{Cell: cell},
		},
	}
	
	table := &Table{
		ct: ctypes.Table{
			RowContents: []ctypes.RowContent{
				{Row: row},
			},
		},
	}
	
	doc.Document = &Document{
		Root: doc,
		Body: &Body{
			Children: []DocumentChild{
				{Table: table},
			},
		},
	}
	
	return doc
}

func TestReplaceFieldResult(t *testing.T) {
	t.Run("Replace single text element", func(t *testing.T) {
		doc := NewRootDoc()
		run := &ctypes.Run{
			Children: []ctypes.RunChild{
				{Text: &ctypes.Text{Text: "original"}},
				{Text: &ctypes.Text{Text: "other"}},
			},
		}
		
		doc.replaceFieldResult(run, []int{0}, "replacement")
		
		assert.Equal(t, "replacement", run.Children[0].Text.Text)
		assert.Equal(t, "other", run.Children[1].Text.Text)
	})

	t.Run("Replace multiple text elements", func(t *testing.T) {
		doc := NewRootDoc()
		run := &ctypes.Run{
			Children: []ctypes.RunChild{
				{Text: &ctypes.Text{Text: "first"}},
				{Text: &ctypes.Text{Text: "second"}},
				{Text: &ctypes.Text{Text: "third"}},
			},
		}
		
		doc.replaceFieldResult(run, []int{0, 1, 2}, "replacement")
		
		assert.Equal(t, "replacement", run.Children[0].Text.Text)
		assert.Equal(t, "", run.Children[1].Text.Text) // Should be cleared
		assert.Equal(t, "", run.Children[2].Text.Text) // Should be cleared
	})

	t.Run("Empty indices does nothing", func(t *testing.T) {
		doc := NewRootDoc()
		run := &ctypes.Run{
			Children: []ctypes.RunChild{
				{Text: &ctypes.Text{Text: "original"}},
			},
		}
		
		doc.replaceFieldResult(run, []int{}, "replacement")
		
		assert.Equal(t, "original", run.Children[0].Text.Text)
	})
}