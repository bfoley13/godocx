package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/bfoley13/godocx/wml/stypes"
)

// BookmarkStart represents the start of a bookmark (w:bookmarkStart)
type BookmarkStart struct {
	// Bookmark ID
	ID *DecimalNum `xml:"id,attr"`

	// Bookmark Name
	Name *CTString `xml:"name,attr"`

	// Column First
	ColFirst *DecimalNum `xml:"colFirst,attr,omitempty"`

	// Column Last
	ColLast *DecimalNum `xml:"colLast,attr,omitempty"`

	// Disable Column Spanning for Revision
	DisplacedByCustomXml *stypes.DisplacedByCustomXml `xml:"displacedByCustomXml,attr,omitempty"`
}

// BookmarkEnd represents the end of a bookmark (w:bookmarkEnd)
type BookmarkEnd struct {
	// Bookmark ID
	ID *DecimalNum `xml:"id,attr"`

	// Disable Column Spanning for Revision
	DisplacedByCustomXml *stypes.DisplacedByCustomXml `xml:"displacedByCustomXml,attr,omitempty"`
}

// MarshalXML implements xml.Marshaler for BookmarkStart
func (b BookmarkStart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:bookmarkStart"

	if b.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(b.ID.Val)})
	}

	if b.Name != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:name"}, Value: b.Name.Val})
	}

	if b.ColFirst != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:colFirst"}, Value: strconv.Itoa(b.ColFirst.Val)})
	}

	if b.ColLast != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:colLast"}, Value: strconv.Itoa(b.ColLast.Val)})
	}

	if b.DisplacedByCustomXml != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:displacedByCustomXml"}, Value: string(*b.DisplacedByCustomXml)})
	}

	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for BookmarkStart
func (b *BookmarkStart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				b.ID = NewDecimalNum(val)
			}
		case "name":
			b.Name = NewCTString(attr.Value)
		case "colFirst":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				b.ColFirst = NewDecimalNum(val)
			}
		case "colLast":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				b.ColLast = NewDecimalNum(val)
			}
		case "displacedByCustomXml":
			val := stypes.DisplacedByCustomXml(attr.Value)
			b.DisplacedByCustomXml = &val
		}
	}

	// Skip to end element
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

// MarshalXML implements xml.Marshaler for BookmarkEnd
func (b BookmarkEnd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:bookmarkEnd"

	if b.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(b.ID.Val)})
	}

	if b.DisplacedByCustomXml != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:displacedByCustomXml"}, Value: string(*b.DisplacedByCustomXml)})
	}

	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for BookmarkEnd
func (b *BookmarkEnd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				b.ID = NewDecimalNum(val)
			}
		case "displacedByCustomXml":
			val := stypes.DisplacedByCustomXml(attr.Value)
			b.DisplacedByCustomXml = &val
		}
	}

	// Skip to end element
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