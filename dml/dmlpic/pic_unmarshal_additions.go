package dmlpic

import (
	"encoding/xml"

	"github.com/bfoley13/godocx/dml/geom"
)

func (p *PresetGeometry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "prst" {
			p.Preset = attr.Value
		}
	}

	// Parse child elements
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "avLst":
				p.AdjustValues = &geom.AdjustValues{}
				if err := d.DecodeElement(p.AdjustValues, &elem); err != nil {
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