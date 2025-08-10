package ctypes

import (
	"encoding/xml"

	"github.com/bfoley13/godocx/wml/stypes"
)

// StructuredDocumentTag represents a content control (w:sdt)
type StructuredDocumentTag struct {
	// Structured Document Tag Properties
	Properties *SdtProperties `xml:"sdtPr,omitempty"`

	// Structured Document Tag Content
	Content *SdtContent `xml:"sdtContent,omitempty"`
}

// SdtProperties represents structured document tag properties (w:sdtPr)
type SdtProperties struct {
	// Friendly Name
	Alias *CTString `xml:"alias,omitempty"`

	// Tag
	Tag *CTString `xml:"tag,omitempty"`

	// Unique ID
	ID *DecimalNum `xml:"id,omitempty"`

	// Lock Setting
	Lock *GenSingleStrVal[stypes.SdtLock] `xml:"lock,omitempty"`

	// Placeholder Document Part Reference
	Placeholder *CTString `xml:"placeholder,omitempty"`

	// Temporary
	Temporary *OnOff `xml:"temporary,omitempty"`

	// Content control type-specific properties
	Text         *SdtText         `xml:"text,omitempty"`
	RichText     *Empty           `xml:"richText,omitempty"`
	Picture      *Empty           `xml:"picture,omitempty"`
	ComboBox     *SdtComboBox     `xml:"comboBox,omitempty"`
	DropDownList *SdtDropDownList `xml:"dropDownList,omitempty"`
	Date         *SdtDate         `xml:"date,omitempty"`
	Checkbox     *SdtCheckbox     `xml:"checkbox,omitempty"`
	Group        *Empty           `xml:"group,omitempty"`
	Citation     *Empty           `xml:"citation,omitempty"`
}

// SdtContent represents structured document tag content (w:sdtContent)
type SdtContent struct {
	// Can contain paragraphs, runs, tables, etc.
	Children []SdtContentChild
}

// SdtContentChild represents possible children of structured document tag content
type SdtContentChild struct {
	Paragraph *Paragraph `xml:"p,omitempty"`
	Run       *Run       `xml:"r,omitempty"`
	Table     *Table     `xml:"tbl,omitempty"`
	// Can add more content types as needed
}

// SdtText represents text content control properties
type SdtText struct {
	// Multi-line
	MultiLine *OnOff `xml:"multiLine,omitempty"`
}

// SdtComboBox represents combo box content control
type SdtComboBox struct {
	// Last Saved Value
	LastValue *CTString `xml:"lastValue,attr,omitempty"`

	// Combo Box List Items
	ListItems []SdtListItem `xml:"listItem,omitempty"`
}

// SdtDropDownList represents drop-down list content control
type SdtDropDownList struct {
	// Last Saved Value
	LastValue *CTString `xml:"lastValue,attr,omitempty"`

	// Drop-Down List Items
	ListItems []SdtListItem `xml:"listItem,omitempty"`
}

// SdtListItem represents a list item in combo box or drop-down list
type SdtListItem struct {
	// Display Text
	DisplayText string `xml:"displayText,attr"`

	// Value
	Value string `xml:"value,attr"`
}

// SdtDate represents date picker content control
type SdtDate struct {
	// Full Date
	FullDate *CTString `xml:"fullDate,attr,omitempty"`

	// Date Format
	DateFormat *CTString `xml:"dateFormat,omitempty"`

	// Language ID
	Lid *CTString `xml:"lid,omitempty"`

	// Storage Format
	StorageFormat *CTString `xml:"storageFormat,omitempty"`

	// Calendar Type
	Calendar *GenSingleStrVal[stypes.CalendarType] `xml:"calendar,omitempty"`
}

// SdtCheckbox represents checkbox content control
type SdtCheckbox struct {
	// Checked State
	Checked *GenSingleStrVal[stypes.OnOff] `xml:"checked,omitempty"`

	// Checked Symbol
	CheckedState *SdtCheckboxSymbol `xml:"checkedState,omitempty"`

	// Unchecked Symbol
	UncheckedState *SdtCheckboxSymbol `xml:"uncheckedState,omitempty"`
}

// SdtCheckboxSymbol represents checkbox symbol
type SdtCheckboxSymbol struct {
	// Font
	Font *CTString `xml:"font,attr,omitempty"`

	// Character Code
	Val *stypes.HexChar `xml:"val,attr,omitempty"`
}

// MarshalXML implements xml.Marshaler for SdtProperties
func (props SdtProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:sdtPr"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if props.Alias != nil {
		if err := e.EncodeElement(props.Alias, xml.StartElement{Name: xml.Name{Local: "w:alias"}}); err != nil {
			return err
		}
	}

	if props.Tag != nil {
		if err := e.EncodeElement(props.Tag, xml.StartElement{Name: xml.Name{Local: "w:tag"}}); err != nil {
			return err
		}
	}

	if props.ID != nil {
		if err := e.EncodeElement(props.ID, xml.StartElement{Name: xml.Name{Local: "w:id"}}); err != nil {
			return err
		}
	}

	if props.Lock != nil {
		if err := e.EncodeElement(props.Lock, xml.StartElement{Name: xml.Name{Local: "w:lock"}}); err != nil {
			return err
		}
	}

	if props.Placeholder != nil {
		if err := e.EncodeElement(props.Placeholder, xml.StartElement{Name: xml.Name{Local: "w:placeholder"}}); err != nil {
			return err
		}
	}

	if props.Temporary != nil {
		if err := e.EncodeElement(props.Temporary, xml.StartElement{Name: xml.Name{Local: "w:temporary"}}); err != nil {
			return err
		}
	}

	// Content control type-specific properties
	if props.Text != nil {
		if err := props.Text.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if props.RichText != nil {
		if err := e.EncodeElement(props.RichText, xml.StartElement{Name: xml.Name{Local: "w:richText"}}); err != nil {
			return err
		}
	}

	if props.Picture != nil {
		if err := e.EncodeElement(props.Picture, xml.StartElement{Name: xml.Name{Local: "w:picture"}}); err != nil {
			return err
		}
	}

	if props.ComboBox != nil {
		if err := props.ComboBox.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if props.DropDownList != nil {
		if err := props.DropDownList.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if props.Date != nil {
		if err := props.Date.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if props.Checkbox != nil {
		if err := props.Checkbox.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if props.Group != nil {
		if err := e.EncodeElement(props.Group, xml.StartElement{Name: xml.Name{Local: "w:group"}}); err != nil {
			return err
		}
	}

	if props.Citation != nil {
		if err := e.EncodeElement(props.Citation, xml.StartElement{Name: xml.Name{Local: "w:citation"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for SdtProperties
func (props *SdtProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "alias":
				props.Alias = &CTString{}
				if err := d.DecodeElement(props.Alias, &elem); err != nil {
					return err
				}
			case "tag":
				props.Tag = &CTString{}
				if err := d.DecodeElement(props.Tag, &elem); err != nil {
					return err
				}
			case "id":
				props.ID = &DecimalNum{}
				if err := d.DecodeElement(props.ID, &elem); err != nil {
					return err
				}
			case "lock":
				props.Lock = &GenSingleStrVal[stypes.SdtLock]{}
				if err := d.DecodeElement(props.Lock, &elem); err != nil {
					return err
				}
			case "placeholder":
				props.Placeholder = &CTString{}
				if err := d.DecodeElement(props.Placeholder, &elem); err != nil {
					return err
				}
			case "temporary":
				props.Temporary = &OnOff{}
				if err := d.DecodeElement(props.Temporary, &elem); err != nil {
					return err
				}
			case "text":
				props.Text = &SdtText{}
				if err := props.Text.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "richText":
				props.RichText = &Empty{}
				if err := d.DecodeElement(props.RichText, &elem); err != nil {
					return err
				}
			case "picture":
				props.Picture = &Empty{}
				if err := d.DecodeElement(props.Picture, &elem); err != nil {
					return err
				}
			case "comboBox":
				props.ComboBox = &SdtComboBox{}
				if err := props.ComboBox.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "dropDownList":
				props.DropDownList = &SdtDropDownList{}
				if err := props.DropDownList.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "date":
				props.Date = &SdtDate{}
				if err := props.Date.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "checkbox":
				props.Checkbox = &SdtCheckbox{}
				if err := props.Checkbox.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "group":
				props.Group = &Empty{}
				if err := d.DecodeElement(props.Group, &elem); err != nil {
					return err
				}
			case "citation":
				props.Citation = &Empty{}
				if err := d.DecodeElement(props.Citation, &elem); err != nil {
					return err
				}
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}

// MarshalXML implements xml.Marshaler for StructuredDocumentTag
func (sdt StructuredDocumentTag) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:sdt"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if sdt.Properties != nil {
		if err := e.EncodeElement(sdt.Properties, xml.StartElement{Name: xml.Name{Local: "w:sdtPr"}}); err != nil {
			return err
		}
	}

	if sdt.Content != nil {
		if err := e.EncodeElement(sdt.Content, xml.StartElement{Name: xml.Name{Local: "w:sdtContent"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for StructuredDocumentTag
func (sdt *StructuredDocumentTag) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "sdtPr":
				sdt.Properties = &SdtProperties{}
				if err := d.DecodeElement(sdt.Properties, &elem); err != nil {
					return err
				}
			case "sdtContent":
				sdt.Content = &SdtContent{}
				if err := d.DecodeElement(sdt.Content, &elem); err != nil {
					return err
				}
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}

// MarshalXML implements xml.Marshaler for SdtContent
func (content SdtContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:sdtContent"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, child := range content.Children {
		if child.Paragraph != nil {
			if err := child.Paragraph.MarshalXML(e, xml.StartElement{}); err != nil {
				return err
			}
		} else if child.Run != nil {
			if err := child.Run.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:r"}}); err != nil {
				return err
			}
		} else if child.Table != nil {
			if err := child.Table.MarshalXML(e, xml.StartElement{}); err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for SdtContent
func (content *SdtContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "p":
				para := &Paragraph{}
				if err := d.DecodeElement(para, &elem); err != nil {
					return err
				}
				content.Children = append(content.Children, SdtContentChild{Paragraph: para})
			case "r":
				run := &Run{}
				if err := d.DecodeElement(run, &elem); err != nil {
					return err
				}
				content.Children = append(content.Children, SdtContentChild{Run: run})
			case "tbl":
				table := &Table{}
				if err := d.DecodeElement(table, &elem); err != nil {
					return err
				}
				content.Children = append(content.Children, SdtContentChild{Table: table})
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}

// MarshalXML implements xml.Marshaler for SdtText
func (text SdtText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:text"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if text.MultiLine != nil {
		if err := e.EncodeElement(text.MultiLine, xml.StartElement{Name: xml.Name{Local: "w:multiLine"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for SdtText
func (text *SdtText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "multiLine":
				text.MultiLine = &OnOff{}
				if err := d.DecodeElement(text.MultiLine, &elem); err != nil {
					return err
				}
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}

// MarshalXML implements xml.Marshaler for SdtComboBox
func (combo SdtComboBox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:comboBox"

	// Add lastValue attribute if present
	if combo.LastValue != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastValue"}, Value: combo.LastValue.Val})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Marshal list items
	for _, item := range combo.ListItems {
		if err := item.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for SdtComboBox
func (combo *SdtComboBox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "lastValue" {
			combo.LastValue = NewCTString(attr.Value)
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "listItem":
				var item SdtListItem
				if err := item.UnmarshalXML(d, elem); err != nil {
					return err
				}
				combo.ListItems = append(combo.ListItems, item)
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}

// MarshalXML implements xml.Marshaler for SdtDropDownList
func (dropdown SdtDropDownList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:dropDownList"

	// Add lastValue attribute if present
	if dropdown.LastValue != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastValue"}, Value: dropdown.LastValue.Val})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Marshal list items
	for _, item := range dropdown.ListItems {
		if err := item.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for SdtDropDownList
func (dropdown *SdtDropDownList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "lastValue" {
			dropdown.LastValue = NewCTString(attr.Value)
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "listItem":
				var item SdtListItem
				if err := item.UnmarshalXML(d, elem); err != nil {
					return err
				}
				dropdown.ListItems = append(dropdown.ListItems, item)
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}

// MarshalXML implements xml.Marshaler for SdtListItem
func (item SdtListItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:listItem"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "w:displayText"}, Value: item.DisplayText},
		{Name: xml.Name{Local: "w:value"}, Value: item.Value},
	}

	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for SdtListItem
func (item *SdtListItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "displayText":
			item.DisplayText = attr.Value
		case "value":
			item.Value = attr.Value
		}
	}

	// Skip to end element
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		if _, ok := token.(xml.EndElement); ok {
			break
		}
	}

	return nil
}

// MarshalXML implements xml.Marshaler for SdtDate
func (date SdtDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:date"

	// Add fullDate attribute if present
	if date.FullDate != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fullDate"}, Value: date.FullDate.Val})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if date.DateFormat != nil {
		if err := e.EncodeElement(date.DateFormat, xml.StartElement{Name: xml.Name{Local: "w:dateFormat"}}); err != nil {
			return err
		}
	}

	if date.Lid != nil {
		if err := e.EncodeElement(date.Lid, xml.StartElement{Name: xml.Name{Local: "w:lid"}}); err != nil {
			return err
		}
	}

	if date.StorageFormat != nil {
		if err := e.EncodeElement(date.StorageFormat, xml.StartElement{Name: xml.Name{Local: "w:storageFormat"}}); err != nil {
			return err
		}
	}

	if date.Calendar != nil {
		if err := e.EncodeElement(date.Calendar, xml.StartElement{Name: xml.Name{Local: "w:calendar"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for SdtDate
func (date *SdtDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "fullDate" {
			date.FullDate = NewCTString(attr.Value)
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "dateFormat":
				date.DateFormat = &CTString{}
				if err := d.DecodeElement(date.DateFormat, &elem); err != nil {
					return err
				}
			case "lid":
				date.Lid = &CTString{}
				if err := d.DecodeElement(date.Lid, &elem); err != nil {
					return err
				}
			case "storageFormat":
				date.StorageFormat = &CTString{}
				if err := d.DecodeElement(date.StorageFormat, &elem); err != nil {
					return err
				}
			case "calendar":
				date.Calendar = &GenSingleStrVal[stypes.CalendarType]{}
				if err := d.DecodeElement(date.Calendar, &elem); err != nil {
					return err
				}
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}

// MarshalXML implements xml.Marshaler for SdtCheckbox
func (checkbox SdtCheckbox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:checkbox"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if checkbox.Checked != nil {
		if err := e.EncodeElement(checkbox.Checked, xml.StartElement{Name: xml.Name{Local: "w:checked"}}); err != nil {
			return err
		}
	}

	if checkbox.CheckedState != nil {
		if err := checkbox.CheckedState.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:checkedState"}}); err != nil {
			return err
		}
	}

	if checkbox.UncheckedState != nil {
		if err := checkbox.UncheckedState.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:uncheckedState"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for SdtCheckbox
func (checkbox *SdtCheckbox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "checked":
				checkbox.Checked = &GenSingleStrVal[stypes.OnOff]{}
				if err := d.DecodeElement(checkbox.Checked, &elem); err != nil {
					return err
				}
			case "checkedState":
				checkbox.CheckedState = &SdtCheckboxSymbol{}
				if err := checkbox.CheckedState.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "uncheckedState":
				checkbox.UncheckedState = &SdtCheckboxSymbol{}
				if err := checkbox.UncheckedState.UnmarshalXML(d, elem); err != nil {
					return err
				}
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}

// MarshalXML implements xml.Marshaler for SdtCheckboxSymbol
func (symbol SdtCheckboxSymbol) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if symbol.Font != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:font"}, Value: symbol.Font.Val})
	}

	if symbol.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(*symbol.Val)})
	}

	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for SdtCheckboxSymbol
func (symbol *SdtCheckboxSymbol) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "font":
			symbol.Font = NewCTString(attr.Value)
		case "val":
			val := stypes.HexChar(attr.Value)
			symbol.Val = &val
		}
	}

	// Skip to end element
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		if _, ok := token.(xml.EndElement); ok {
			break
		}
	}

	return nil
}