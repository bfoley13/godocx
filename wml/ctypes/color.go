package ctypes

import (
	"encoding/xml"

	"github.com/bfoley13/godocx/wml/stypes"
)

// Color represents the color of a text or element.
type Color struct {
	//Run Content Color
	Val string `xml:"val,attr"`

	//Run Content Theme Color
	ThemeColor *stypes.ThemeColor `xml:"themeColor,attr,omitempty"`

	//Run Content Theme Color Tint
	ThemeTint *string `xml:"themeTint,attr,omitempty"`

	//Run Content Theme Color Shade
	ThemeShade *string `xml:"themeShade,attr,omitempty"`
}

// NewColor creates a new Color instance with the specified color value.
func NewColor(value string) *Color {
	return &Color{Val: value}
}

// MarshalXML implements the xml.Marshaler interface for the Color type.
func (c Color) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:color"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: c.Val})

	if c.ThemeColor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeColor"}, Value: string(*c.ThemeColor)})
	}

	if c.ThemeTint != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeTint"}, Value: *c.ThemeTint})
	}

	if c.ThemeShade != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:themeShade"}, Value: *c.ThemeShade})
	}

	return e.EncodeElement("", start)
}

func (c *Color) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "val":
			c.Val = attr.Value
		case "themeColor":
			themeColor := stypes.ThemeColor(attr.Value)
			c.ThemeColor = &themeColor
		case "themeTint":
			value := attr.Value
			c.ThemeTint = &value
		case "themeShade":
			value := attr.Value
			c.ThemeShade = &value
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
