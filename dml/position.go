package dml

import (
	"encoding/xml"
	"errors"

	"github.com/bfoley13/godocx/dml/dmlst"
)

type PoistionH struct {
	RelativeFrom dmlst.RelFromH `xml:"relativeFrom,attr"`
	PosOffset    int            `xml:"posOffset"`
}

type PoistionV struct {
	RelativeFrom dmlst.RelFromV `xml:"relativeFrom,attr"`
	PosOffset    int            `xml:"posOffset"`
}

func (p PoistionH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if p.RelativeFrom == "" {
		return errors.New("Invalid RelativeFrom in PoistionH")
	}

	start.Name.Local = "wp:positionH"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeFrom"}, Value: string(p.RelativeFrom)})

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	offsetElem := xml.StartElement{Name: xml.Name{Local: "wp:posOffset"}}
	if err = e.EncodeElement(p.PosOffset, offsetElem); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (p PoistionV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if p.RelativeFrom == "" {
		return errors.New("Invalid RelativeFrom in PoistionV")
	}

	start.Name.Local = "wp:positionV"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeFrom"}, Value: string(p.RelativeFrom)})

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	offsetElem := xml.StartElement{Name: xml.Name{Local: "wp:posOffset"}}
	if err = e.EncodeElement(p.PosOffset, offsetElem); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (p *PoistionH) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "relativeFrom" {
			p.RelativeFrom = dmlst.RelFromH(attr.Value)
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
			case "posOffset":
				if err := d.DecodeElement(&p.PosOffset, &elem); err != nil {
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

func (p *PoistionV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "relativeFrom" {
			p.RelativeFrom = dmlst.RelFromV(attr.Value)
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
			case "posOffset":
				if err := d.DecodeElement(&p.PosOffset, &elem); err != nil {
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
