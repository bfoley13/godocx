package dml

import (
	"encoding/xml"
	"strconv"

	"github.com/bfoley13/godocx/dml/dmlct"
	"github.com/bfoley13/godocx/dml/dmlst"
)

func (w *WrapNone) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// WrapNone is a simple element with no content
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

func (ws *WrapSquare) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "wrapText":
			ws.WrapText = dmlst.WrapText(attr.Value)
		case "distT":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dist := uint(val)
				ws.DistT = &dist
			}
		case "distB":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dist := uint(val)
				ws.DistB = &dist
			}
		case "distL":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dist := uint(val)
				ws.DistL = &dist
			}
		case "distR":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dist := uint(val)
				ws.DistR = &dist
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
			case "effectExtent":
				ws.EffectExtent = &EffectExtent{}
				if err := d.DecodeElement(ws.EffectExtent, &elem); err != nil {
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

func (wp *WrapPolygon) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "edited" {
			if val, err := strconv.ParseBool(attr.Value); err == nil {
				wp.Edited = &val
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
			case "start":
				if err := d.DecodeElement(&wp.Start, &elem); err != nil {
					return err
				}
			case "lineTo":
				var point dmlct.Point2D
				if err := d.DecodeElement(&point, &elem); err != nil {
					return err
				}
				wp.LineTo = append(wp.LineTo, point)
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

func (w *WrapTight) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "wrapText":
			w.WrapText = dmlst.WrapText(attr.Value)
		case "distL":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dist := uint(val)
				w.DistL = &dist
			}
		case "distR":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dist := uint(val)
				w.DistR = &dist
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
			case "wrapPolygon":
				if err := d.DecodeElement(&w.WrapPolygon, &elem); err != nil {
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

func (w *WrapThrough) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "wrapText":
			w.WrapText = dmlst.WrapText(attr.Value)
		case "distL":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dist := uint(val)
				w.DistL = &dist
			}
		case "distR":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dist := uint(val)
				w.DistR = &dist
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
			case "wrapPolygon":
				if err := d.DecodeElement(&w.WrapPolygon, &elem); err != nil {
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

func (w *WrapTopBtm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "distT":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dist := uint(val)
				w.DistT = &dist
			}
		case "distB":
			if val, err := strconv.ParseUint(attr.Value, 10, 32); err == nil {
				dist := uint(val)
				w.DistB = &dist
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
			case "effectExtent":
				w.EffectExtent = &EffectExtent{}
				if err := d.DecodeElement(w.EffectExtent, &elem); err != nil {
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