package docx

import (
	"encoding/xml"

	"github.com/bfoley13/godocx/wml/stypes"
)

// Specifies the background information for this document
//
// This background shall be displayed on all pages of the document, behind all other document content.
type Background struct {
	Color      *string            `xml:"color,attr,omitempty"`
	ThemeColor *stypes.ThemeColor `xml:"themeColor,attr,omitempty"`
	ThemeTint  *string            `xml:"themeTint,attr,omitempty"`
	ThemeShade *string            `xml:"themeShade,attr,omitempty"`
}

func NewBackground() *Background {
	return &Background{}
}
func (b Background) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:background"
	if b.Color != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:color"}, Value: *b.Color})
	}
	if b.ThemeColor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeColor"}, Value: string(*b.ThemeColor)})
	}
	if b.ThemeTint != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeTint"}, Value: *b.ThemeTint})
	}
	if b.ThemeShade != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeShade"}, Value: *b.ThemeShade})
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	return e.EncodeToken(start.End())
}

func (b *Background) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes - handle both with and without namespace prefixes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "color":
			value := attr.Value
			b.Color = &value
		case "themeColor":
			themeColor := stypes.ThemeColor(attr.Value)
			b.ThemeColor = &themeColor
		case "themeTint":
			value := attr.Value
			b.ThemeTint = &value
		case "themeShade":
			value := attr.Value
			b.ThemeShade = &value
		}
	}

	// Skip any content and read to end element
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
