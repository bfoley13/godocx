package ctypes

import (
	"encoding/xml"
	"testing"

	"github.com/bfoley13/godocx/internal"
	"github.com/bfoley13/godocx/wml/stypes"
	"github.com/stretchr/testify/assert"
)

func TestSDTFullRoundTripMarshaling(t *testing.T) {
	// Create a comprehensive SDT with all major components
	original := &StructuredDocumentTag{
		Properties: &SdtProperties{
			Alias:     NewCTString("Full Test Control"),
			Tag:       NewCTString("fulltest"),
			ID:        NewDecimalNum(999999),
			Lock:      NewGenSingleStrVal(stypes.SdtLockSdtContentLocked),
			Temporary: OnOffFromBool(false),
			ComboBox: &SdtComboBox{
				LastValue: NewCTString("value1"),
				ListItems: []SdtListItem{
					{DisplayText: "First Option", Value: "value1"},
					{DisplayText: "Second Option", Value: "value2"},
					{DisplayText: "Third Option", Value: "value3"},
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
										{Text: TextFromString("Test content in paragraph")},
									},
								},
							},
						},
					},
				},
				{
					Run: &Run{
						Children: []RunChild{
							{Text: TextFromString("Direct run content")},
						},
					},
				},
			},
		},
	}

	// Marshal to XML
	xmlData, err := xml.Marshal(original)
	assert.NoError(t, err)
	
	xmlStr := string(xmlData)
	
	// Verify expected content is in XML
	assert.Contains(t, xmlStr, "w:sdt")
	assert.Contains(t, xmlStr, "w:sdtPr")
	assert.Contains(t, xmlStr, "w:sdtContent")
	assert.Contains(t, xmlStr, "Full Test Control")
	assert.Contains(t, xmlStr, "fulltest")
	assert.Contains(t, xmlStr, "999999")
	assert.Contains(t, xmlStr, "sdtContentLocked")
	assert.Contains(t, xmlStr, "w:comboBox")
	assert.Contains(t, xmlStr, "value1")
	assert.Contains(t, xmlStr, "First Option")
	assert.Contains(t, xmlStr, "Test content in paragraph")
	assert.Contains(t, xmlStr, "Direct run content")

	// Unmarshal back to struct
	var unmarshaled StructuredDocumentTag
	err = xml.Unmarshal(xmlData, &unmarshaled)
	assert.NoError(t, err)

	// Verify structure integrity
	assert.NotNil(t, unmarshaled.Properties)
	assert.Equal(t, "Full Test Control", unmarshaled.Properties.Alias.Val)
	assert.Equal(t, "fulltest", unmarshaled.Properties.Tag.Val)
	assert.Equal(t, 999999, unmarshaled.Properties.ID.Val)
	assert.Equal(t, stypes.SdtLockSdtContentLocked, unmarshaled.Properties.Lock.Val)
	
	// Verify combo box properties
	assert.NotNil(t, unmarshaled.Properties.ComboBox)
	assert.Equal(t, "value1", unmarshaled.Properties.ComboBox.LastValue.Val)
	assert.Len(t, unmarshaled.Properties.ComboBox.ListItems, 3)
	assert.Equal(t, "First Option", unmarshaled.Properties.ComboBox.ListItems[0].DisplayText)
	assert.Equal(t, "value1", unmarshaled.Properties.ComboBox.ListItems[0].Value)
	
	// Verify content structure
	assert.NotNil(t, unmarshaled.Content)
	assert.Len(t, unmarshaled.Content.Children, 2)
	
	// Check paragraph content
	assert.NotNil(t, unmarshaled.Content.Children[0].Paragraph)
	paragraphRun := unmarshaled.Content.Children[0].Paragraph.Children[0].Run
	assert.NotNil(t, paragraphRun)
	assert.Equal(t, "Test content in paragraph", paragraphRun.Children[0].Text.Text)
	
	// Check direct run content
	assert.NotNil(t, unmarshaled.Content.Children[1].Run)
	directRun := unmarshaled.Content.Children[1].Run
	assert.Equal(t, "Direct run content", directRun.Children[0].Text.Text)
}

func TestSDTTextControlRoundTrip(t *testing.T) {
	original := &StructuredDocumentTag{
		Properties: &SdtProperties{
			Alias: NewCTString("Text Input"),
			Tag:   NewCTString("textinput"),
			ID:    NewDecimalNum(12345),
			Text: &SdtText{
				MultiLine: OnOffFromBool(true),
			},
		},
		Content: &SdtContent{
			Children: []SdtContentChild{
				{
					Run: &Run{
						Children: []RunChild{
							{Text: TextFromString("Sample multiline\ntext input")},
						},
					},
				},
			},
		},
	}

	// Marshal to XML
	xmlData, err := xml.Marshal(original)
	assert.NoError(t, err)
	
	xmlStr := string(xmlData)
	assert.Contains(t, xmlStr, "w:text")
	assert.Contains(t, xmlStr, "w:multiLine")
	assert.Contains(t, xmlStr, "true")

	// Unmarshal back
	var unmarshaled StructuredDocumentTag
	err = xml.Unmarshal(xmlData, &unmarshaled)
	assert.NoError(t, err)

	// Verify text control properties
	assert.NotNil(t, unmarshaled.Properties.Text)
	assert.NotNil(t, unmarshaled.Properties.Text.MultiLine)
	assert.Equal(t, stypes.OnOffTrue, *unmarshaled.Properties.Text.MultiLine.Val)
}

func TestSDTCheckboxRoundTrip(t *testing.T) {
	original := &StructuredDocumentTag{
		Properties: &SdtProperties{
			Alias: NewCTString("Agreement Checkbox"),
			Tag:   NewCTString("agree"),
			ID:    NewDecimalNum(54321),
			Checkbox: &SdtCheckbox{
				Checked: NewGenSingleStrVal(stypes.OnOffTrue),
				CheckedState: &SdtCheckboxSymbol{
					Font: NewCTString("Wingdings"),
					Val:  internal.ToPtr(stypes.HexChar("2713")),
				},
				UncheckedState: &SdtCheckboxSymbol{
					Font: NewCTString("Wingdings"),
					Val:  internal.ToPtr(stypes.HexChar("2717")),
				},
			},
		},
		Content: &SdtContent{
			Children: []SdtContentChild{
				{
					Run: &Run{
						Children: []RunChild{
							{Text: TextFromString("☑")},
						},
					},
				},
			},
		},
	}

	// Marshal to XML
	xmlData, err := xml.Marshal(original)
	assert.NoError(t, err)
	
	xmlStr := string(xmlData)
	assert.Contains(t, xmlStr, "w:checkbox")
	assert.Contains(t, xmlStr, "w:checked")
	assert.Contains(t, xmlStr, "w:checkedState")
	assert.Contains(t, xmlStr, "w:uncheckedState")
	assert.Contains(t, xmlStr, "Wingdings")
	assert.Contains(t, xmlStr, "2713")
	assert.Contains(t, xmlStr, "2717")

	// Unmarshal back
	var unmarshaled StructuredDocumentTag
	err = xml.Unmarshal(xmlData, &unmarshaled)
	assert.NoError(t, err)

	// Verify checkbox properties
	checkbox := unmarshaled.Properties.Checkbox
	assert.NotNil(t, checkbox)
	assert.Equal(t, stypes.OnOffTrue, checkbox.Checked.Val)
	assert.Equal(t, "Wingdings", checkbox.CheckedState.Font.Val)
	assert.Equal(t, "2713", string(*checkbox.CheckedState.Val))
	assert.Equal(t, "Wingdings", checkbox.UncheckedState.Font.Val)
	assert.Equal(t, "2717", string(*checkbox.UncheckedState.Val))
}

func TestSDTDateControlRoundTrip(t *testing.T) {
	original := &StructuredDocumentTag{
		Properties: &SdtProperties{
			Alias: NewCTString("Date Picker"),
			Tag:   NewCTString("datepicker"),
			ID:    NewDecimalNum(67890),
			Date: &SdtDate{
				FullDate:      NewCTString("2024-12-25T00:00:00Z"),
				DateFormat:    NewCTString("MM/dd/yyyy"),
				Calendar:      NewGenSingleStrVal(stypes.CalendarTypeGregorian),
				StorageFormat: NewCTString("dateTime"),
			},
		},
		Content: &SdtContent{
			Children: []SdtContentChild{
				{
					Run: &Run{
						Children: []RunChild{
							{Text: TextFromString("12/25/2024")},
						},
					},
				},
			},
		},
	}

	// Marshal to XML
	xmlData, err := xml.Marshal(original)
	assert.NoError(t, err)
	
	xmlStr := string(xmlData)
	assert.Contains(t, xmlStr, "w:date")
	assert.Contains(t, xmlStr, "2024-12-25T00:00:00Z")
	assert.Contains(t, xmlStr, "w:dateFormat")
	assert.Contains(t, xmlStr, "MM/dd/yyyy")
	assert.Contains(t, xmlStr, "gregorian")

	// Unmarshal back
	var unmarshaled StructuredDocumentTag
	err = xml.Unmarshal(xmlData, &unmarshaled)
	assert.NoError(t, err)

	// Verify date control properties
	date := unmarshaled.Properties.Date
	assert.NotNil(t, date)
	assert.Equal(t, "2024-12-25T00:00:00Z", date.FullDate.Val)
	assert.Equal(t, "MM/dd/yyyy", date.DateFormat.Val)
	assert.Equal(t, stypes.CalendarTypeGregorian, date.Calendar.Val)
	assert.Equal(t, "dateTime", date.StorageFormat.Val)
}