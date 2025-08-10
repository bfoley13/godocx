package dml

import (
	"encoding/xml"

	"github.com/bfoley13/godocx/dml/dmlpic"
)

func (g *Graphic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse child elements
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "graphicData":
				g.Data = &GraphicData{}
				if err := d.DecodeElement(g.Data, &elem); err != nil {
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

func (gd *GraphicData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "uri" {
			gd.URI = attr.Value
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
			case "pic":
				gd.Pic = &dmlpic.Pic{}
				if err := d.DecodeElement(gd.Pic, &elem); err != nil {
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