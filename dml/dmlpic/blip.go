package dmlpic

import "encoding/xml"

// Binary large image or picture
type Blip struct {
	EmbedID string `xml:"embed,attr,omitempty"`
}

func (b Blip) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:blip"

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "r:embed"}, Value: b.EmbedID},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (b *Blip) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "embed" || (attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "embed") {
			b.EmbedID = attr.Value
		}
	}

	// Skip any content and read to end
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
