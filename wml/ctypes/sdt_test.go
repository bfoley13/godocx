package ctypes

import (
	"encoding/xml"
	"testing"

	"github.com/bfoley13/godocx/internal"
	"github.com/bfoley13/godocx/wml/stypes"
	"github.com/stretchr/testify/assert"
)

func TestStructuredDocumentTagMarshalXML(t *testing.T) {
	sdt := &StructuredDocumentTag{
		Properties: &SdtProperties{
			Alias: NewCTString("Test Control"),
			Tag:   NewCTString("test"),
			ID:    NewDecimalNum(12345),
			Text:  &SdtText{MultiLine: OnOffFromBool(true)},
		},
		Content: &SdtContent{
			Children: []SdtContentChild{
				{
					Run: &Run{
						Children: []RunChild{
							{Text: TextFromString("Hello World")},
						},
					},
				},
			},
		},
	}

	// Marshal to XML
	xmlData, err := xml.Marshal(sdt)
	assert.NoError(t, err)

	xmlStr := string(xmlData)
	assert.Contains(t, xmlStr, "w:sdt")
	assert.Contains(t, xmlStr, "w:sdtPr")
	assert.Contains(t, xmlStr, "w:sdtContent")
	assert.Contains(t, xmlStr, "Test Control")
	assert.Contains(t, xmlStr, "Hello World")
}

func TestStructuredDocumentTagUnmarshalXML(t *testing.T) {
	xmlInput := `<w:sdt xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
		<w:sdtPr>
			<w:alias w:val="Test Control"/>
			<w:tag w:val="test"/>
			<w:id w:val="12345"/>
			<w:text>
				<w:multiLine w:val="true"/>
			</w:text>
		</w:sdtPr>
		<w:sdtContent>
			<w:r>
				<w:t>Hello World</w:t>
			</w:r>
		</w:sdtContent>
	</w:sdt>`

	var sdt StructuredDocumentTag
	err := xml.Unmarshal([]byte(xmlInput), &sdt)
	assert.NoError(t, err)

	assert.NotNil(t, sdt.Properties)
	assert.NotNil(t, sdt.Properties.Alias)
	assert.Equal(t, "Test Control", sdt.Properties.Alias.Val)
	assert.Equal(t, "test", sdt.Properties.Tag.Val)
	assert.Equal(t, 12345, sdt.Properties.ID.Val)

	assert.NotNil(t, sdt.Content)
	assert.Len(t, sdt.Content.Children, 1)
	assert.NotNil(t, sdt.Content.Children[0].Run)
}

func TestSdtComboBoxMarshal(t *testing.T) {
	comboBox := &SdtComboBox{
		LastValue: NewCTString("option1"),
		ListItems: []SdtListItem{
			{DisplayText: "Option 1", Value: "option1"},
			{DisplayText: "Option 2", Value: "option2"},
		},
	}

	xmlData, err := xml.Marshal(comboBox)
	assert.NoError(t, err)

	xmlStr := string(xmlData)
	assert.Contains(t, xmlStr, "option1")
	assert.Contains(t, xmlStr, "Option 1")
	assert.Contains(t, xmlStr, "Option 2")
}

func TestSdtDropDownListMarshal(t *testing.T) {
	dropDown := &SdtDropDownList{
		LastValue: NewCTString("choice2"),
		ListItems: []SdtListItem{
			{DisplayText: "Choice 1", Value: "choice1"},
			{DisplayText: "Choice 2", Value: "choice2"},
		},
	}

	xmlData, err := xml.Marshal(dropDown)
	assert.NoError(t, err)

	xmlStr := string(xmlData)
	assert.Contains(t, xmlStr, "choice2")
	assert.Contains(t, xmlStr, "Choice 1")
	assert.Contains(t, xmlStr, "Choice 2")
}

func TestSdtDateMarshal(t *testing.T) {
	date := &SdtDate{
		FullDate:      NewCTString("2024-12-25T00:00:00Z"),
		DateFormat:    NewCTString("yyyy-MM-dd"),
		Calendar:      NewGenSingleStrVal(stypes.CalendarTypeGregorian),
		StorageFormat: NewCTString("date"),
	}

	xmlData, err := xml.Marshal(date)
	assert.NoError(t, err)

	xmlStr := string(xmlData)
	assert.Contains(t, xmlStr, "2024-12-25")
	assert.Contains(t, xmlStr, "yyyy-MM-dd")
	assert.Contains(t, xmlStr, "gregorian")
}

func TestSdtCheckboxMarshal(t *testing.T) {
	checkbox := &SdtCheckbox{
		Checked: NewGenSingleStrVal(stypes.OnOffTrue),
		CheckedState: &SdtCheckboxSymbol{
			Font: NewCTString("Wingdings"),
			Val:  internal.ToPtr(stypes.HexChar("2713")),
		},
		UncheckedState: &SdtCheckboxSymbol{
			Font: NewCTString("Wingdings"),
			Val:  internal.ToPtr(stypes.HexChar("2717")),
		},
	}

	xmlData, err := xml.Marshal(checkbox)
	assert.NoError(t, err)

	xmlStr := string(xmlData)
	assert.Contains(t, xmlStr, "true")
	assert.Contains(t, xmlStr, "Wingdings")
	assert.Contains(t, xmlStr, "2713")
	assert.Contains(t, xmlStr, "2717")
}

func TestSdtContentMarshal(t *testing.T) {
	content := &SdtContent{
		Children: []SdtContentChild{
			{
				Run: &Run{
					Children: []RunChild{
						{Text: TextFromString("Run text")},
					},
				},
			},
			{
				Paragraph: &Paragraph{
					Children: []ParagraphChild{
						{
							Run: &Run{
								Children: []RunChild{
									{Text: TextFromString("Paragraph text")},
								},
							},
						},
					},
				},
			},
		},
	}

	xmlData, err := xml.Marshal(content)
	assert.NoError(t, err)

	xmlStr := string(xmlData)
	assert.Contains(t, xmlStr, "w:sdtContent")
	assert.Contains(t, xmlStr, "Run text")
	assert.Contains(t, xmlStr, "Paragraph text")
}

func TestSdtContentUnmarshal(t *testing.T) {
	xmlInput := `<w:sdtContent xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
		<w:r>
			<w:t>Run content</w:t>
		</w:r>
		<w:p>
			<w:r>
				<w:t>Paragraph content</w:t>
			</w:r>
		</w:p>
	</w:sdtContent>`

	var content SdtContent
	err := xml.Unmarshal([]byte(xmlInput), &content)
	assert.NoError(t, err)

	assert.Len(t, content.Children, 2)
	assert.NotNil(t, content.Children[0].Run)
	assert.NotNil(t, content.Children[1].Paragraph)
}

func TestComplexSdtStructure(t *testing.T) {
	// Test a complex SDT with nested content
	sdt := &StructuredDocumentTag{
		Properties: &SdtProperties{
			Alias:     NewCTString("Complex Control"),
			Tag:       NewCTString("complex"),
			ID:        NewDecimalNum(99999),
			Lock:      NewGenSingleStrVal(stypes.SdtLockContentLocked),
			Temporary: OnOffFromBool(false),
			ComboBox: &SdtComboBox{
				LastValue: NewCTString("val1"),
				ListItems: []SdtListItem{
					{DisplayText: "Value 1", Value: "val1"},
					{DisplayText: "Value 2", Value: "val2"},
				},
			},
		},
		Content: &SdtContent{
			Children: []SdtContentChild{
				{
					Paragraph: &Paragraph{
						Children: []ParagraphChild{
							{
								Run: &Run{
									Children: []RunChild{
										{Text: TextFromString("Complex content")},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(sdt)
	assert.NoError(t, err)
	
	xmlStr := string(xmlData)
	assert.Contains(t, xmlStr, "Complex Control")
	assert.Contains(t, xmlStr, "complex")
	assert.Contains(t, xmlStr, "contentLocked")
	assert.Contains(t, xmlStr, "Value 1")
	assert.Contains(t, xmlStr, "Complex content")

	// Test unmarshaling back
	var unmarshaledSdt StructuredDocumentTag
	err = xml.Unmarshal(xmlData, &unmarshaledSdt)
	assert.NoError(t, err)
	
	assert.NotNil(t, unmarshaledSdt.Properties)
	assert.Equal(t, "Complex Control", unmarshaledSdt.Properties.Alias.Val)
	assert.NotNil(t, unmarshaledSdt.Content)
}