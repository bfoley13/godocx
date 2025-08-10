package dml

import (
	"encoding/xml"
)

type NonVisualGraphicFrameProp struct {
	GraphicFrameLocks *GraphicFrameLocks `xml:"graphicFrameLocks,omitempty"`
}

func (n NonVisualGraphicFrameProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:cNvGraphicFramePr"
	start.Attr = []xml.Attr{}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if n.GraphicFrameLocks != nil {
		if err := e.EncodeElement(n.GraphicFrameLocks, xml.StartElement{Name: xml.Name{Local: "a:graphicFrameLocks"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (n *NonVisualGraphicFrameProp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse child elements
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "graphicFrameLocks":
				n.GraphicFrameLocks = &GraphicFrameLocks{}
				if err := d.DecodeElement(n.GraphicFrameLocks, &elem); err != nil {
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
