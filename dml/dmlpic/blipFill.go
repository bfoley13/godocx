package dmlpic

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/bfoley13/godocx/dml/dmlct"
	"github.com/bfoley13/godocx/dml/shapes"
)

type BlipFill struct {
	// 1. Blip
	Blip *Blip `xml:"blip,omitempty"`

	//2.Source Rectangle
	SrcRect *dmlct.RelativeRect `xml:"srcRect,omitempty"`

	// 3. Choice of a:EG_FillModeProperties
	FillModeProps FillModeProps `xml:",any"`

	//Attributes:
	DPI          *uint32 `xml:"dpi,attr,omitempty"`          //DPI Setting
	RotWithShape *bool   `xml:"rotWithShape,attr,omitempty"` //Rotate With Shape
}

// NewBlipFill creates a new BlipFill with the given relationship ID (rID)
// The rID is used to reference the image in the presentation.
func NewBlipFill(rID string) BlipFill {
	return BlipFill{
		Blip: &Blip{
			EmbedID: rID,
		},
	}
}

func (b BlipFill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:blipFill"

	if b.DPI != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dpi"}, Value: fmt.Sprintf("%d", *b.DPI)})
	}

	if b.RotWithShape != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rotWithShape"}, Value: fmt.Sprintf("%t", *b.RotWithShape)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. Blip
	if b.Blip != nil {
		if err := b.Blip.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "a:blip"}}); err != nil {
			return err
		}
	}

	// 2. SrcRect
	if b.SrcRect != nil {
		if err = b.SrcRect.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "a:SrcRect"}}); err != nil {
			return err
		}
	}

	// 3. Choice: FillModProperties
	if err = b.FillModeProps.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type FillModeProps struct {
	Stretch *shapes.Stretch `xml:"stretch,omitempty"`
	Tile    *shapes.Tile    `xml:"tile,omitempty"`
}

func (f FillModeProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if f.Stretch != nil {
		return f.Stretch.MarshalXML(e, xml.StartElement{})
	}

	if f.Tile != nil {
		return f.Tile.MarshalXML(e, xml.StartElement{})
	}

	return nil
}

func (b *BlipFill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "dpi":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dpi := uint32(val)
				b.DPI = &dpi
			}
		case "rotWithShape":
			if val, err := strconv.ParseBool(attr.Value); err == nil {
				b.RotWithShape = &val
			}
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
			case "blip":
				b.Blip = &Blip{}
				if err := d.DecodeElement(b.Blip, &elem); err != nil {
					return err
				}
			case "srcRect":
				b.SrcRect = &dmlct.RelativeRect{}
				if err := d.DecodeElement(b.SrcRect, &elem); err != nil {
					return err
				}
			case "stretch":
				b.FillModeProps.Stretch = &shapes.Stretch{}
				if err := d.DecodeElement(b.FillModeProps.Stretch, &elem); err != nil {
					return err
				}
			case "tile":
				b.FillModeProps.Tile = &shapes.Tile{}
				if err := d.DecodeElement(b.FillModeProps.Tile, &elem); err != nil {
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

func (f *FillModeProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse child elements
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "stretch":
				f.Stretch = &shapes.Stretch{}
				if err := d.DecodeElement(f.Stretch, &elem); err != nil {
					return err
				}
			case "tile":
				f.Tile = &shapes.Tile{}
				if err := d.DecodeElement(f.Tile, &elem); err != nil {
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
