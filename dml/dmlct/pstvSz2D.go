package dmlct

import (
	"encoding/xml"
	"strconv"

	"github.com/bfoley13/godocx/common/units"
)

// Complex Type: CT_PositiveSize2D
type PSize2D struct {
	Width  uint64 `xml:"cx,attr,omitempty"`
	Height uint64 `xml:"cy,attr,omitempty"`
}

func NewPostvSz2D(width units.Emu, height units.Emu) *PSize2D {
	return &PSize2D{
		Height: uint64(height),
		Width:  uint64(width),
	}
}

func (p PSize2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cx"}, Value: strconv.FormatUint(p.Width, 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cy"}, Value: strconv.FormatUint(p.Height, 10)})

	return e.EncodeElement("", start)
}

func (p *PSize2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "cx":
			if val, err := strconv.ParseUint(attr.Value, 10, 64); err == nil {
				p.Width = val
			}
		case "cy":
			if val, err := strconv.ParseUint(attr.Value, 10, 64); err == nil {
				p.Height = val
			}
		}
	}

	// Skip to end element since this is self-closing
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
