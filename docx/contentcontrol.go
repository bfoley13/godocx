package docx

import (
	"time"

	"github.com/bfoley13/godocx/internal"
	"github.com/bfoley13/godocx/wml/ctypes"
	"github.com/bfoley13/godocx/wml/stypes"
)

// ContentControl represents a wrapper around the structured document tag
type ContentControl struct {
	root *RootDoc
	sdt  *ctypes.StructuredDocumentTag
}

// newContentControl creates a new ContentControl wrapper
func newContentControl(root *RootDoc, sdt *ctypes.StructuredDocumentTag) *ContentControl {
	return &ContentControl{
		root: root,
		sdt:  sdt,
	}
}

// AddContentControl adds a basic content control to the document
func (rd *RootDoc) AddContentControl(alias, tag string, controlType ContentControlType) *ContentControl {
	sdt := &ctypes.StructuredDocumentTag{
		Properties: &ctypes.SdtProperties{
			Alias: ctypes.NewCTString(alias),
			Tag:   ctypes.NewCTString(tag),
			ID:    ctypes.NewDecimalNum(rd.generateContentControlID()),
		},
		Content: &ctypes.SdtContent{},
	}

	// Set control type-specific properties
	switch controlType {
	case ContentControlTypeText:
		sdt.Properties.Text = &ctypes.SdtText{}
	case ContentControlTypeRichText:
		sdt.Properties.RichText = &ctypes.Empty{}
	case ContentControlTypePicture:
		sdt.Properties.Picture = &ctypes.Empty{}
	case ContentControlTypeGroup:
		sdt.Properties.Group = &ctypes.Empty{}
	}

	// Add to document body
	para := rd.AddEmptyParagraph()
	para.ct.Children = append(para.ct.Children, ctypes.ParagraphChild{Sdt: sdt})

	return newContentControl(rd, sdt)
}

// AddTextContentControl adds a text content control with initial content
func (rd *RootDoc) AddTextContentControl(alias, tag, initialText string, multiline bool) *ContentControl {
	sdt := &ctypes.StructuredDocumentTag{
		Properties: &ctypes.SdtProperties{
			Alias: ctypes.NewCTString(alias),
			Tag:   ctypes.NewCTString(tag),
			ID:    ctypes.NewDecimalNum(rd.generateContentControlID()),
			Text: &ctypes.SdtText{
				MultiLine: ctypes.OnOffFromBool(multiline),
			},
		},
		Content: &ctypes.SdtContent{},
	}

	// Add initial text if provided
	if initialText != "" {
		run := ctypes.NewRun()
		run.Children = append(run.Children, ctypes.RunChild{
			Text: ctypes.TextFromString(initialText),
		})
		sdt.Content.Children = append(sdt.Content.Children, ctypes.SdtContentChild{Run: run})
	}

	// Add to document body
	para := rd.AddEmptyParagraph()
	para.ct.Children = append(para.ct.Children, ctypes.ParagraphChild{Sdt: sdt})

	return newContentControl(rd, sdt)
}

// AddComboBoxContentControl adds a combo box content control
func (rd *RootDoc) AddComboBoxContentControl(alias, tag string, items []ContentControlListItem, lastValue string) *ContentControl {
	listItems := make([]ctypes.SdtListItem, len(items))
	for i, item := range items {
		listItems[i] = ctypes.SdtListItem{
			DisplayText: item.DisplayText,
			Value:       item.Value,
		}
	}

	sdt := &ctypes.StructuredDocumentTag{
		Properties: &ctypes.SdtProperties{
			Alias: ctypes.NewCTString(alias),
			Tag:   ctypes.NewCTString(tag),
			ID:    ctypes.NewDecimalNum(rd.generateContentControlID()),
			ComboBox: &ctypes.SdtComboBox{
				LastValue: ctypes.NewCTString(lastValue),
				ListItems: listItems,
			},
		},
		Content: &ctypes.SdtContent{},
	}

	// Add initial value as text if provided
	if lastValue != "" {
		run := ctypes.NewRun()
		run.Children = append(run.Children, ctypes.RunChild{
			Text: ctypes.TextFromString(lastValue),
		})
		sdt.Content.Children = append(sdt.Content.Children, ctypes.SdtContentChild{Run: run})
	}

	// Add to document body
	para := rd.AddEmptyParagraph()
	para.ct.Children = append(para.ct.Children, ctypes.ParagraphChild{Sdt: sdt})

	return newContentControl(rd, sdt)
}

// AddDropDownContentControl adds a drop-down list content control
func (rd *RootDoc) AddDropDownContentControl(alias, tag string, items []ContentControlListItem, lastValue string) *ContentControl {
	listItems := make([]ctypes.SdtListItem, len(items))
	for i, item := range items {
		listItems[i] = ctypes.SdtListItem{
			DisplayText: item.DisplayText,
			Value:       item.Value,
		}
	}

	sdt := &ctypes.StructuredDocumentTag{
		Properties: &ctypes.SdtProperties{
			Alias: ctypes.NewCTString(alias),
			Tag:   ctypes.NewCTString(tag),
			ID:    ctypes.NewDecimalNum(rd.generateContentControlID()),
			DropDownList: &ctypes.SdtDropDownList{
				LastValue: ctypes.NewCTString(lastValue),
				ListItems: listItems,
			},
		},
		Content: &ctypes.SdtContent{},
	}

	// Add initial value as text if provided
	if lastValue != "" {
		run := ctypes.NewRun()
		run.Children = append(run.Children, ctypes.RunChild{
			Text: ctypes.TextFromString(lastValue),
		})
		sdt.Content.Children = append(sdt.Content.Children, ctypes.SdtContentChild{Run: run})
	}

	// Add to document body
	para := rd.AddEmptyParagraph()
	para.ct.Children = append(para.ct.Children, ctypes.ParagraphChild{Sdt: sdt})

	return newContentControl(rd, sdt)
}

// AddDateContentControl adds a date picker content control
func (rd *RootDoc) AddDateContentControl(alias, tag string, fullDate *time.Time, dateFormat string) *ContentControl {
	sdt := &ctypes.StructuredDocumentTag{
		Properties: &ctypes.SdtProperties{
			Alias: ctypes.NewCTString(alias),
			Tag:   ctypes.NewCTString(tag),
			ID:    ctypes.NewDecimalNum(rd.generateContentControlID()),
			Date: &ctypes.SdtDate{
				DateFormat: ctypes.NewCTString(dateFormat),
				Calendar:   ctypes.NewGenSingleStrVal(stypes.CalendarTypeGregorian),
			},
		},
		Content: &ctypes.SdtContent{},
	}

	if fullDate != nil {
		dateStr := fullDate.Format(time.RFC3339)
		sdt.Properties.Date.FullDate = ctypes.NewCTString(dateStr)
		
		// Add formatted date as initial content
		displayDate := fullDate.Format("2006-01-02")
		if dateFormat != "" {
			// You could implement custom date formatting here
			displayDate = fullDate.Format("2006-01-02") // Simplified for now
		}
		
		run := ctypes.NewRun()
		run.Children = append(run.Children, ctypes.RunChild{
			Text: ctypes.TextFromString(displayDate),
		})
		sdt.Content.Children = append(sdt.Content.Children, ctypes.SdtContentChild{Run: run})
	}

	// Add to document body
	para := rd.AddEmptyParagraph()
	para.ct.Children = append(para.ct.Children, ctypes.ParagraphChild{Sdt: sdt})

	return newContentControl(rd, sdt)
}

// AddCheckboxContentControl adds a checkbox content control
func (rd *RootDoc) AddCheckboxContentControl(alias, tag string, checked bool) *ContentControl {
	sdt := &ctypes.StructuredDocumentTag{
		Properties: &ctypes.SdtProperties{
			Alias: ctypes.NewCTString(alias),
			Tag:   ctypes.NewCTString(tag),
			ID:    ctypes.NewDecimalNum(rd.generateContentControlID()),
			Checkbox: &ctypes.SdtCheckbox{
				Checked: func() *ctypes.GenSingleStrVal[stypes.OnOff] {
					if checked {
						return ctypes.NewGenSingleStrVal(stypes.OnOffTrue)
					}
					return ctypes.NewGenSingleStrVal(stypes.OnOffFalse)
				}(),
				CheckedState: &ctypes.SdtCheckboxSymbol{
					Font: ctypes.NewCTString("Wingdings"),
					Val:  internal.ToPtr(stypes.HexChar("2713")), // Checkmark
				},
				UncheckedState: &ctypes.SdtCheckboxSymbol{
					Font: ctypes.NewCTString("Wingdings"),
					Val:  internal.ToPtr(stypes.HexChar("2717")), // X mark
				},
			},
		},
		Content: &ctypes.SdtContent{},
	}

	// Add checkbox symbol as content
	symbol := "☐" // Unchecked box
	if checked {
		symbol = "☑" // Checked box
	}
	
	run := ctypes.NewRun()
	run.Children = append(run.Children, ctypes.RunChild{
		Text: ctypes.TextFromString(symbol),
	})
	sdt.Content.Children = append(sdt.Content.Children, ctypes.SdtContentChild{Run: run})

	// Add to document body
	para := rd.AddEmptyParagraph()
	para.ct.Children = append(para.ct.Children, ctypes.ParagraphChild{Sdt: sdt})

	return newContentControl(rd, sdt)
}

// SetText sets the text content of the content control
func (cc *ContentControl) SetText(text string) *ContentControl {
	// Clear existing content
	cc.sdt.Content.Children = []ctypes.SdtContentChild{}
	
	// Add new text content
	run := ctypes.NewRun()
	run.Children = append(run.Children, ctypes.RunChild{
		Text: ctypes.TextFromString(text),
	})
	cc.sdt.Content.Children = append(cc.sdt.Content.Children, ctypes.SdtContentChild{Run: run})
	
	return cc
}

// SetLock sets the lock setting for the content control
func (cc *ContentControl) SetLock(lock stypes.SdtLock) *ContentControl {
	if cc.sdt.Properties == nil {
		cc.sdt.Properties = &ctypes.SdtProperties{}
	}
	cc.sdt.Properties.Lock = ctypes.NewGenSingleStrVal(lock)
	return cc
}

// SetTemporary marks the content control as temporary
func (cc *ContentControl) SetTemporary(temporary bool) *ContentControl {
	if cc.sdt.Properties == nil {
		cc.sdt.Properties = &ctypes.SdtProperties{}
	}
	cc.sdt.Properties.Temporary = ctypes.OnOffFromBool(temporary)
	return cc
}

// GetTag returns the tag value of the content control
func (cc *ContentControl) GetTag() string {
	if cc.sdt.Properties != nil && cc.sdt.Properties.Tag != nil {
		return cc.sdt.Properties.Tag.Val
	}
	return ""
}

// GetAlias returns the alias (friendly name) of the content control
func (cc *ContentControl) GetAlias() string {
	if cc.sdt.Properties != nil && cc.sdt.Properties.Alias != nil {
		return cc.sdt.Properties.Alias.Val
	}
	return ""
}

// GetText returns the text content of the content control
func (cc *ContentControl) GetText() string {
	if cc.sdt.Content == nil {
		return ""
	}
	
	var text string
	for _, child := range cc.sdt.Content.Children {
		if child.Run != nil {
			for _, runChild := range child.Run.Children {
				if runChild.Text != nil {
					text += runChild.Text.Text
				}
			}
		}
	}
	return text
}

// generateContentControlID generates a unique ID for content controls
func (rd *RootDoc) generateContentControlID() int {
	// Simple implementation - you might want to track this more carefully
	return int(time.Now().UnixNano() % 1000000000)
}

// ContentControlType represents the type of content control
type ContentControlType int

const (
	ContentControlTypeText ContentControlType = iota
	ContentControlTypeRichText
	ContentControlTypePicture
	ContentControlTypeComboBox
	ContentControlTypeDropDownList
	ContentControlTypeDate
	ContentControlTypeCheckbox
	ContentControlTypeGroup
)

// ContentControlListItem represents an item in a combo box or drop-down list
type ContentControlListItem struct {
	DisplayText string
	Value       string
}