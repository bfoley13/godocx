package dml

import (
	"encoding/xml"
	"strconv"
)

type DocProp struct {
	ID          uint64 `xml:"id,attr,omitempty"`
	Name        string `xml:"name,attr,omitempty"`
	Description string `xml:"descr,attr,omitempty"`

	//TODO: Remaining attrs & child elements
}

func (d DocProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:docPr"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "id"}, Value: strconv.FormatUint(d.ID, 10)},
		{Name: xml.Name{Local: "name"}, Value: d.Name},
	}

	if d.Description != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "descr"}, Value: d.Description})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (d *DocProp) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			if val, err := strconv.ParseUint(attr.Value, 10, 64); err == nil {
				d.ID = val
			}
		case "name":
			d.Name = attr.Value
		case "descr":
			d.Description = attr.Value
		}
	}

	// Skip to end element since DocProp currently has no child elements
	return decoder.Skip()
}
