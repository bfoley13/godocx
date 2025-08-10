package ctypes

import (
	"encoding/xml"

	"github.com/bfoley13/godocx/wml/stypes"
)

// FieldChar represents a field character (w:fldChar)
type FieldChar struct {
	// Field Character Type
	FldCharType *GenSingleStrVal[stypes.FldCharType] `xml:"fldCharType,attr"`

	// Field Should Not Be Recalculated
	Dirty *OnOff `xml:"dirty,attr,omitempty"`

	// Field Result Invalidated
	Lock *OnOff `xml:"lock,attr,omitempty"`

	// Form Field Properties
	FFData *FFData `xml:"ffData,omitempty"`

	// Numbering Level Associated with Field
	NumId *DecimalNum `xml:"numId,omitempty"`
}

// FieldCode represents field code (w:instrText)
type FieldCode struct {
	// Preserve Whitespace
	Space *stypes.Space `xml:"space,attr,omitempty"`

	// Field instruction text
	Text string `xml:",chardata"`
}

// FFData represents form field data (w:ffData)
type FFData struct {
	// Form Field Name
	Name *FFName `xml:"name,omitempty"`

	// Form Field Status Text
	StatusText *FFStatusText `xml:"statusText,omitempty"`

	// Form Field Help Text
	HelpText *FFHelpText `xml:"helpText,omitempty"`

	// Form Field Properties
	Enabled *OnOff `xml:"enabled,omitempty"`

	// Calculate on Exit
	CalcOnExit *OnOff `xml:"calcOnExit,omitempty"`

	// Entry Macro
	EntryMacro *FFMacro `xml:"entryMacro,omitempty"`

	// Exit Macro
	ExitMacro *FFMacro `xml:"exitMacro,omitempty"`

	// Text Form Field Properties
	TextInput *FFTextInput `xml:"textInput,omitempty"`

	// Checkbox Form Field Properties
	CheckBox *FFCheckBox `xml:"checkBox,omitempty"`

	// Drop-Down List Form Field Properties
	DDList *FFDDList `xml:"ddList,omitempty"`
}

// FFName represents form field name (w:name)
type FFName struct {
	// Form Field Name Value
	Val *CTString `xml:"val,attr,omitempty"`
}

// FFStatusText represents form field status text (w:statusText)
type FFStatusText struct {
	// Status Text Type
	Type *GenSingleStrVal[stypes.InfoTextType] `xml:"type,attr,omitempty"`

	// Status Text Value
	Val *CTString `xml:"val,attr,omitempty"`
}

// FFHelpText represents form field help text (w:helpText)
type FFHelpText struct {
	// Help Text Type
	Type *GenSingleStrVal[stypes.InfoTextType] `xml:"type,attr,omitempty"`

	// Help Text Value
	Val *CTString `xml:"val,attr,omitempty"`
}

// FFMacro represents form field macro (w:entryMacro, w:exitMacro)
type FFMacro struct {
	// Macro Name
	Val *CTString `xml:"val,attr,omitempty"`
}

// FFTextInput represents text form field properties (w:textInput)
type FFTextInput struct {
	// Text Form Field Type
	Type *GenSingleStrVal[stypes.TextFormFieldType] `xml:"type,omitempty"`

	// Default Text Form Field String
	Default *CTString `xml:"default,omitempty"`

	// Text Form Field Maximum Length
	MaxLength *DecimalNum `xml:"maxLength,omitempty"`

	// Text Form Field Formatting
	Format *CTString `xml:"format,omitempty"`
}

// FFCheckBox represents checkbox form field properties (w:checkBox)
type FFCheckBox struct {
	// Checkbox Size Type
	Size *DecimalNum `xml:"size,omitempty"`

	// Checkbox Size Automatic
	SizeAuto *OnOff `xml:"sizeAuto,omitempty"`

	// Default Checkbox Form Field State
	Default *OnOff `xml:"default,omitempty"`

	// Checkbox Form Field State
	Checked *OnOff `xml:"checked,omitempty"`
}

// FFDDList represents drop-down list form field properties (w:ddList)
type FFDDList struct {
	// Drop-Down List Selection
	Result *DecimalNum `xml:"result,omitempty"`

	// Default Drop-Down List Selection
	Default *DecimalNum `xml:"default,omitempty"`

	// Drop-Down List Entries
	ListEntries []FFListEntry `xml:"listEntry,omitempty"`
}

// FFListEntry represents drop-down list entry (w:listEntry)
type FFListEntry struct {
	// List Entry Value
	Val *CTString `xml:"val,attr,omitempty"`
}

// MarshalXML implements xml.Marshaler for FieldChar
func (f FieldChar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:fldChar"

	if f.FldCharType != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fldCharType"}, Value: string(f.FldCharType.Val)})
	}

	if f.Dirty != nil && f.Dirty.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:dirty"}, Value: string(*f.Dirty.Val)})
	}

	if f.Lock != nil && f.Lock.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lock"}, Value: string(*f.Lock.Val)})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if f.FFData != nil {
		if err := f.FFData.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if f.NumId != nil {
		if err := e.EncodeElement(f.NumId, xml.StartElement{Name: xml.Name{Local: "w:numId"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for FieldChar
func (f *FieldChar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "fldCharType":
			f.FldCharType = &GenSingleStrVal[stypes.FldCharType]{Val: stypes.FldCharType(attr.Value)}
		case "dirty":
			val := stypes.OnOff(attr.Value)
			f.Dirty = &OnOff{Val: &val}
		case "lock":
			val := stypes.OnOff(attr.Value)
			f.Lock = &OnOff{Val: &val}
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
			case "ffData":
				f.FFData = &FFData{}
				if err := f.FFData.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "numId":
				f.NumId = &DecimalNum{}
				if err := d.DecodeElement(f.NumId, &elem); err != nil {
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

// MarshalXML implements xml.Marshaler for FieldCode
func (f FieldCode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:instrText"

	if f.Space != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "xml", Local: "space"}, Value: string(*f.Space)})
	}

	return e.EncodeElement(f.Text, start)
}

// UnmarshalXML implements xml.Unmarshaler for FieldCode
func (f *FieldCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "space" && attr.Name.Space == "xml" {
			space := stypes.Space(attr.Value)
			f.Space = &space
		}
	}

	// Get the text content
	var text string
	if err := d.DecodeElement(&text, &start); err != nil {
		return err
	}
	f.Text = text

	return nil
}

// MarshalXML implements xml.Marshaler for FFData
func (f FFData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:ffData"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if f.Name != nil {
		if err := f.Name.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if f.Enabled != nil {
		if err := e.EncodeElement(f.Enabled, xml.StartElement{Name: xml.Name{Local: "w:enabled"}}); err != nil {
			return err
		}
	}

	if f.CalcOnExit != nil {
		if err := e.EncodeElement(f.CalcOnExit, xml.StartElement{Name: xml.Name{Local: "w:calcOnExit"}}); err != nil {
			return err
		}
	}

	if f.StatusText != nil {
		if err := f.StatusText.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if f.HelpText != nil {
		if err := f.HelpText.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if f.EntryMacro != nil {
		if err := f.EntryMacro.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:entryMacro"}}); err != nil {
			return err
		}
	}

	if f.ExitMacro != nil {
		if err := f.ExitMacro.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:exitMacro"}}); err != nil {
			return err
		}
	}

	if f.TextInput != nil {
		if err := f.TextInput.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if f.CheckBox != nil {
		if err := f.CheckBox.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if f.DDList != nil {
		if err := f.DDList.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for FFData
func (f *FFData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "name":
				f.Name = &FFName{}
				if err := f.Name.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "enabled":
				f.Enabled = &OnOff{}
				if err := d.DecodeElement(f.Enabled, &elem); err != nil {
					return err
				}
			case "calcOnExit":
				f.CalcOnExit = &OnOff{}
				if err := d.DecodeElement(f.CalcOnExit, &elem); err != nil {
					return err
				}
			case "statusText":
				f.StatusText = &FFStatusText{}
				if err := f.StatusText.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "helpText":
				f.HelpText = &FFHelpText{}
				if err := f.HelpText.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "entryMacro":
				f.EntryMacro = &FFMacro{}
				if err := f.EntryMacro.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "exitMacro":
				f.ExitMacro = &FFMacro{}
				if err := f.ExitMacro.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "textInput":
				f.TextInput = &FFTextInput{}
				if err := f.TextInput.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "checkBox":
				f.CheckBox = &FFCheckBox{}
				if err := f.CheckBox.UnmarshalXML(d, elem); err != nil {
					return err
				}
			case "ddList":
				f.DDList = &FFDDList{}
				if err := f.DDList.UnmarshalXML(d, elem); err != nil {
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

// MarshalXML implementations for all form field sub-types follow the same pattern...
// For brevity, I'll implement key ones:

// MarshalXML implements xml.Marshaler for FFName
func (f FFName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:name"
	if f.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: f.Val.Val})
	}
	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for FFName
func (f *FFName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			f.Val = NewCTString(attr.Value)
		}
	}
	// Skip to end
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

// MarshalXML implements xml.Marshaler for FFTextInput
func (f FFTextInput) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:textInput"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if f.Type != nil {
		if err := e.EncodeElement(f.Type, xml.StartElement{Name: xml.Name{Local: "w:type"}}); err != nil {
			return err
		}
	}

	if f.Default != nil {
		if err := e.EncodeElement(f.Default, xml.StartElement{Name: xml.Name{Local: "w:default"}}); err != nil {
			return err
		}
	}

	if f.MaxLength != nil {
		if err := e.EncodeElement(f.MaxLength, xml.StartElement{Name: xml.Name{Local: "w:maxLength"}}); err != nil {
			return err
		}
	}

	if f.Format != nil {
		if err := e.EncodeElement(f.Format, xml.StartElement{Name: xml.Name{Local: "w:format"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for FFTextInput
func (f *FFTextInput) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "type":
				f.Type = &GenSingleStrVal[stypes.TextFormFieldType]{}
				if err := d.DecodeElement(f.Type, &elem); err != nil {
					return err
				}
			case "default":
				f.Default = &CTString{}
				if err := d.DecodeElement(f.Default, &elem); err != nil {
					return err
				}
			case "maxLength":
				f.MaxLength = &DecimalNum{}
				if err := d.DecodeElement(f.MaxLength, &elem); err != nil {
					return err
				}
			case "format":
				f.Format = &CTString{}
				if err := d.DecodeElement(f.Format, &elem); err != nil {
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

// MarshalXML implements xml.Marshaler for FFStatusText
func (f FFStatusText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:statusText"
	if f.Type != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(f.Type.Val)})
	}
	if f.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: f.Val.Val})
	}
	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for FFStatusText
func (f *FFStatusText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "type":
			f.Type = &GenSingleStrVal[stypes.InfoTextType]{Val: stypes.InfoTextType(attr.Value)}
		case "val":
			f.Val = NewCTString(attr.Value)
		}
	}
	// Skip to end
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

// MarshalXML implements xml.Marshaler for FFHelpText
func (f FFHelpText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:helpText"
	if f.Type != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(f.Type.Val)})
	}
	if f.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: f.Val.Val})
	}
	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for FFHelpText
func (f *FFHelpText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "type":
			f.Type = &GenSingleStrVal[stypes.InfoTextType]{Val: stypes.InfoTextType(attr.Value)}
		case "val":
			f.Val = NewCTString(attr.Value)
		}
	}
	// Skip to end
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

// MarshalXML implements xml.Marshaler for FFMacro
func (f FFMacro) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if f.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: f.Val.Val})
	}
	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for FFMacro
func (f *FFMacro) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			f.Val = NewCTString(attr.Value)
		}
	}
	// Skip to end
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

// MarshalXML implements xml.Marshaler for FFCheckBox
func (f FFCheckBox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:checkBox"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if f.Size != nil {
		if err := e.EncodeElement(f.Size, xml.StartElement{Name: xml.Name{Local: "w:size"}}); err != nil {
			return err
		}
	}

	if f.SizeAuto != nil {
		if err := e.EncodeElement(f.SizeAuto, xml.StartElement{Name: xml.Name{Local: "w:sizeAuto"}}); err != nil {
			return err
		}
	}

	if f.Default != nil {
		if err := e.EncodeElement(f.Default, xml.StartElement{Name: xml.Name{Local: "w:default"}}); err != nil {
			return err
		}
	}

	if f.Checked != nil {
		if err := e.EncodeElement(f.Checked, xml.StartElement{Name: xml.Name{Local: "w:checked"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for FFCheckBox
func (f *FFCheckBox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "size":
				f.Size = &DecimalNum{}
				if err := d.DecodeElement(f.Size, &elem); err != nil {
					return err
				}
			case "sizeAuto":
				f.SizeAuto = &OnOff{}
				if err := d.DecodeElement(f.SizeAuto, &elem); err != nil {
					return err
				}
			case "default":
				f.Default = &OnOff{}
				if err := d.DecodeElement(f.Default, &elem); err != nil {
					return err
				}
			case "checked":
				f.Checked = &OnOff{}
				if err := d.DecodeElement(f.Checked, &elem); err != nil {
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

// MarshalXML implements xml.Marshaler for FFDDList
func (f FFDDList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:ddList"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if f.Result != nil {
		if err := e.EncodeElement(f.Result, xml.StartElement{Name: xml.Name{Local: "w:result"}}); err != nil {
			return err
		}
	}

	if f.Default != nil {
		if err := e.EncodeElement(f.Default, xml.StartElement{Name: xml.Name{Local: "w:default"}}); err != nil {
			return err
		}
	}

	for _, entry := range f.ListEntries {
		if err := entry.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for FFDDList
func (f *FFDDList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "result":
				f.Result = &DecimalNum{}
				if err := d.DecodeElement(f.Result, &elem); err != nil {
					return err
				}
			case "default":
				f.Default = &DecimalNum{}
				if err := d.DecodeElement(f.Default, &elem); err != nil {
					return err
				}
			case "listEntry":
				var entry FFListEntry
				if err := entry.UnmarshalXML(d, elem); err != nil {
					return err
				}
				f.ListEntries = append(f.ListEntries, entry)
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

// MarshalXML implements xml.Marshaler for FFListEntry
func (f FFListEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:listEntry"
	if f.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: f.Val.Val})
	}
	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for FFListEntry
func (f *FFListEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			f.Val = NewCTString(attr.Value)
		}
	}
	// Skip to end
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