package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/bfoley13/godocx/wml/stypes"
)

// Comment represents a comment (w:comment)
type Comment struct {
	// Comment ID
	ID *DecimalNum `xml:"id,attr"`

	// Comment Author
	Author *CTString `xml:"author,attr,omitempty"`

	// Comment Date
	Date *CTString `xml:"date,attr,omitempty"`

	// Annotation Identifier
	Initials *CTString `xml:"initials,attr,omitempty"`

	// Comment content - can contain paragraphs, tables, etc.
	Children []CommentChild
}

// CommentChild represents possible children of comment content
type CommentChild struct {
	Paragraph *Paragraph `xml:"p,omitempty"`
	Table     *Table     `xml:"tbl,omitempty"`
}

// CommentRangeStart represents the start of a comment range (w:commentRangeStart)
type CommentRangeStart struct {
	// Comment ID
	ID *DecimalNum `xml:"id,attr"`

	// Disable Column Spanning for Revision
	DisplacedByCustomXml *stypes.DisplacedByCustomXml `xml:"displacedByCustomXml,attr,omitempty"`
}

// CommentRangeEnd represents the end of a comment range (w:commentRangeEnd)
type CommentRangeEnd struct {
	// Comment ID
	ID *DecimalNum `xml:"id,attr"`

	// Disable Column Spanning for Revision
	DisplacedByCustomXml *stypes.DisplacedByCustomXml `xml:"displacedByCustomXml,attr,omitempty"`
}

// CommentReference represents a comment reference (w:commentReference)
type CommentReference struct {
	// Comment ID
	ID *DecimalNum `xml:"id,attr"`
}

// MarshalXML implements xml.Marshaler for Comment
func (c Comment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:comment"

	if c.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(c.ID.Val)})
	}

	if c.Author != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"}, Value: c.Author.Val})
	}

	if c.Date != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"}, Value: c.Date.Val})
	}

	if c.Initials != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:initials"}, Value: c.Initials.Val})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Marshal children
	for _, child := range c.Children {
		if child.Paragraph != nil {
			if err := child.Paragraph.MarshalXML(e, xml.StartElement{}); err != nil {
				return err
			}
		} else if child.Table != nil {
			if err := child.Table.MarshalXML(e, xml.StartElement{}); err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for Comment
func (c *Comment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				c.ID = NewDecimalNum(val)
			}
		case "author":
			c.Author = NewCTString(attr.Value)
		case "date":
			c.Date = NewCTString(attr.Value)
		case "initials":
			c.Initials = NewCTString(attr.Value)
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "p":
				para := &Paragraph{}
				if err := d.DecodeElement(para, &elem); err != nil {
					return err
				}
				c.Children = append(c.Children, CommentChild{Paragraph: para})
			case "tbl":
				table := &Table{}
				if err := d.DecodeElement(table, &elem); err != nil {
					return err
				}
				c.Children = append(c.Children, CommentChild{Table: table})
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

// MarshalXML implements xml.Marshaler for CommentRangeStart
func (c CommentRangeStart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:commentRangeStart"

	if c.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(c.ID.Val)})
	}

	if c.DisplacedByCustomXml != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:displacedByCustomXml"}, Value: string(*c.DisplacedByCustomXml)})
	}

	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for CommentRangeStart
func (c *CommentRangeStart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				c.ID = NewDecimalNum(val)
			}
		case "displacedByCustomXml":
			val := stypes.DisplacedByCustomXml(attr.Value)
			c.DisplacedByCustomXml = &val
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

// MarshalXML implements xml.Marshaler for CommentRangeEnd
func (c CommentRangeEnd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:commentRangeEnd"

	if c.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(c.ID.Val)})
	}

	if c.DisplacedByCustomXml != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:displacedByCustomXml"}, Value: string(*c.DisplacedByCustomXml)})
	}

	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for CommentRangeEnd
func (c *CommentRangeEnd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				c.ID = NewDecimalNum(val)
			}
		case "displacedByCustomXml":
			val := stypes.DisplacedByCustomXml(attr.Value)
			c.DisplacedByCustomXml = &val
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

// MarshalXML implements xml.Marshaler for CommentReference
func (c CommentReference) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:commentReference"

	if c.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(c.ID.Val)})
	}

	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for CommentReference
func (c *CommentReference) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			if val, err := strconv.Atoi(attr.Value); err == nil {
				c.ID = NewDecimalNum(val)
			}
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