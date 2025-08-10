package ctypes

import (
	"encoding/xml"
	"strconv"
)

// Indent represents the Paragraph Indentation structure.
type Indent struct {
	Left           *int    `xml:"left,attr,omitempty"`           // Left Indentation
	LeftChars      *int    `xml:"leftChars,attr,omitempty"`      // Left Indentation in Character Units
	Right          *int    `xml:"right,attr,omitempty"`          // Right Indentation
	RightChars     *int    `xml:"rightChars,attr,omitempty"`     // Right Indentation in Character Units
	Hanging        *uint64 `xml:"hanging,attr,omitempty"`        // Indentation Removed from First Line
	HangingChars   *int    `xml:"hangingChars,attr,omitempty"`   // Indentation Removed From First Line in Character Units
	FirstLine      *uint64 `xml:"firstLine,attr,omitempty"`      // Additional First Line Indentation
	FirstLineChars *int    `xml:"firstLineChars,attr,omitempty"` // Additional First Line Indentation in Character Units
}

// MarshalXML implements the xml.Marshaler interface for Indent.
func (i Indent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:ind"
	start.Attr = []xml.Attr{}

	if i.Left != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:left"}, Value: strconv.Itoa(*i.Left)})
	}

	if i.LeftChars != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:leftChars"}, Value: strconv.Itoa(*i.LeftChars)})
	}

	if i.Right != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:right"}, Value: strconv.Itoa(*i.Right)})
	}

	if i.RightChars != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rightChars"}, Value: strconv.Itoa(*i.RightChars)})
	}

	if i.Hanging != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hanging"}, Value: strconv.FormatUint(*i.Hanging, 10)})
	}

	if i.HangingChars != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hangingChars"}, Value: strconv.Itoa(*i.HangingChars)})
	}

	if i.FirstLine != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:firstLine"}, Value: strconv.FormatUint(*i.FirstLine, 10)})
	}

	if i.FirstLineChars != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:firstLineChars"}, Value: strconv.Itoa(*i.FirstLineChars)})
	}

	return e.EncodeElement("", start)
}

func (i *Indent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "left":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				i.Left = &val
			}
		case "leftChars":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				i.LeftChars = &val
			}
		case "right":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				i.Right = &val
			}
		case "rightChars":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				i.RightChars = &val
			}
		case "hanging":
			if val, err := strconv.ParseUint(attr.Value, 10, 64); err == nil {
				i.Hanging = &val
			}
		case "hangingChars":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				i.HangingChars = &val
			}
		case "firstLine":
			if val, err := strconv.ParseUint(attr.Value, 10, 64); err == nil {
				i.FirstLine = &val
			}
		case "firstLineChars":
			if val, err := strconv.Atoi(attr.Value); err == nil {
				i.FirstLineChars = &val
			}
		}
	}

	// Skip to end element since indent is self-closing
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
