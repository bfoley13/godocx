package dml

import (
	"encoding/xml"
	"fmt"
)

// Shape represents a wp:shape element for drawing shapes in Word documents
type Shape struct {
	// Required attributes
	ID   *string `xml:"id,attr,omitempty"`
	Type *string `xml:"type,attr,omitempty"`

	// Optional attributes
	Style      *string `xml:"style,attr,omitempty"`
	FillColor  *string `xml:"fillcolor,attr,omitempty"`
	StrokeColor *string `xml:"strokecolor,attr,omitempty"`
	StrokeWeight *string `xml:"strokeweight,attr,omitempty"`

	// Child elements
	Path       *ShapePath        `xml:"path,omitempty"`
	TextBox    *ShapeTextBox     `xml:"textbox,omitempty"`
	ImageData  *ShapeImageData   `xml:"imagedata,omitempty"`
	Fill       *ShapeFill        `xml:"fill,omitempty"`
	Stroke     *ShapeStroke      `xml:"stroke,omitempty"`
}

// ShapePath represents the path element for vector shapes
type ShapePath struct {
	V            *string `xml:"v,attr,omitempty"`
	ConnectType  *string `xml:"connecttype,attr,omitempty"`
	ConnectLocs  *string `xml:"connectlocs,attr,omitempty"`
}

// ShapeTextBox represents the textbox element for shapes containing text
type ShapeTextBox struct {
	Style *string `xml:"style,attr,omitempty"`
	// Text content would be WordprocessingML paragraphs
	Content string `xml:",innerxml"`
}

// MarshalXML for ShapeTextBox to set the correct element name
func (stb ShapeTextBox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Override the element name to v:textbox
	start.Name.Local = "v:textbox"
	
	// Let the default marshaling handle attributes and innerxml content
	type Alias ShapeTextBox
	return e.EncodeElement(Alias(stb), start)
}

// ShapeImageData represents imagedata element for shapes containing images
type ShapeImageData struct {
	ID    *string `xml:"r:id,attr,omitempty"`
	Title *string `xml:"title,attr,omitempty"`
}

// ShapeFill represents fill properties for shapes
type ShapeFill struct {
	Type   *string `xml:"type,attr,omitempty"`
	Color  *string `xml:"color,attr,omitempty"`
	Opacity *string `xml:"opacity,attr,omitempty"`
}

// ShapeStroke represents stroke properties for shapes
type ShapeStroke struct {
	Color    *string `xml:"color,attr,omitempty"`
	Weight   *string `xml:"weight,attr,omitempty"`
	DashStyle *string `xml:"dashstyle,attr,omitempty"`
}

// NewShape creates a new Shape with the specified ID and type
func NewShape(id, shapeType string) *Shape {
	return &Shape{
		ID:   &id,
		Type: &shapeType,
	}
}

// WithStyle sets the style attribute for the shape
func (s *Shape) WithStyle(style string) *Shape {
	s.Style = &style
	return s
}

// WithFillColor sets the fill color for the shape
func (s *Shape) WithFillColor(color string) *Shape {
	s.FillColor = &color
	return s
}

// WithStrokeColor sets the stroke color for the shape
func (s *Shape) WithStrokeColor(color string) *Shape {
	s.StrokeColor = &color
	return s
}

// WithStrokeWeight sets the stroke weight for the shape
func (s *Shape) WithStrokeWeight(weight string) *Shape {
	s.StrokeWeight = &weight
	return s
}

// WithPath adds a path element to the shape
func (s *Shape) WithPath(v string) *Shape {
	s.Path = &ShapePath{V: &v}
	return s
}

// WithTextBox adds a textbox element to the shape
func (s *Shape) WithTextBox(content string) *Shape {
	s.TextBox = &ShapeTextBox{Content: content}
	return s
}

// WithImageData adds image data to the shape
func (s *Shape) WithImageData(relationshipID string) *Shape {
	s.ImageData = &ShapeImageData{ID: &relationshipID}
	return s
}

// MarshalXML marshals the Shape to XML
func (s Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "v:shape"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"},
		{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"},
	}

	// Add attributes
	if s.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"}, Value: *s.ID})
	}
	if s.Type != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: *s.Type})
	}
	if s.Style != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"}, Value: *s.Style})
	}
	if s.FillColor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillcolor"}, Value: *s.FillColor})
	}
	if s.StrokeColor != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strokecolor"}, Value: *s.StrokeColor})
	}
	if s.StrokeWeight != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strokeweight"}, Value: *s.StrokeWeight})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// Encode child elements
	if s.Path != nil {
		if err := s.Path.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "v:path"}}); err != nil {
			return fmt.Errorf("marshalling Path: %w", err)
		}
	}
	if s.TextBox != nil {
		if err := e.EncodeElement(s.TextBox, xml.StartElement{Name: xml.Name{Local: "v:textbox"}}); err != nil {
			return fmt.Errorf("marshalling TextBox: %w", err)
		}
	}
	if s.ImageData != nil {
		if err := s.ImageData.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "v:imagedata"}}); err != nil {
			return fmt.Errorf("marshalling ImageData: %w", err)
		}
	}
	if s.Fill != nil {
		if err := s.Fill.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "v:fill"}}); err != nil {
			return fmt.Errorf("marshalling Fill: %w", err)
		}
	}
	if s.Stroke != nil {
		if err := s.Stroke.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "v:stroke"}}); err != nil {
			return fmt.Errorf("marshalling Stroke: %w", err)
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML unmarshals the Shape from XML
func (s *Shape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		attrValue := attr.Value // Create a copy to avoid pointer aliasing
		switch attr.Name.Local {
		case "id":
			s.ID = &attrValue
		case "type":
			s.Type = &attrValue
		case "style":
			s.Style = &attrValue
		case "fillcolor":
			s.FillColor = &attrValue
		case "strokecolor":
			s.StrokeColor = &attrValue
		case "strokeweight":
			s.StrokeWeight = &attrValue
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
			case "path":
				s.Path = &ShapePath{}
				if err := d.DecodeElement(s.Path, &elem); err != nil {
					return err
				}
			case "textbox":
				s.TextBox = &ShapeTextBox{}
				if err := d.DecodeElement(s.TextBox, &elem); err != nil {
					return err
				}
			case "imagedata":
				s.ImageData = &ShapeImageData{}
				if err := d.DecodeElement(s.ImageData, &elem); err != nil {
					return err
				}
			case "fill":
				s.Fill = &ShapeFill{}
				if err := d.DecodeElement(s.Fill, &elem); err != nil {
					return err
				}
			case "stroke":
				s.Stroke = &ShapeStroke{}
				if err := d.DecodeElement(s.Stroke, &elem); err != nil {
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

// MarshalXML for ShapePath
func (sp ShapePath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "v:path"
	start.Attr = []xml.Attr{}

	if sp.V != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "v"}, Value: *sp.V})
	}
	if sp.ConnectType != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connecttype"}, Value: *sp.ConnectType})
	}
	if sp.ConnectLocs != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connectlocs"}, Value: *sp.ConnectLocs})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML for ShapePath
func (sp *ShapePath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		attrValue := attr.Value // Create a copy to avoid pointer aliasing
		switch attr.Name.Local {
		case "v":
			sp.V = &attrValue
		case "connecttype":
			sp.ConnectType = &attrValue
		case "connectlocs":
			sp.ConnectLocs = &attrValue
		}
	}

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


// MarshalXML for ShapeImageData
func (sid ShapeImageData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "v:imagedata"
	start.Attr = []xml.Attr{}

	if sid.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"}, Value: *sid.ID})
	}
	if sid.Title != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "title"}, Value: *sid.Title})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML for ShapeImageData
func (sid *ShapeImageData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		attrValue := attr.Value // Create a copy to avoid pointer aliasing
		switch attr.Name.Local {
		case "id":
			sid.ID = &attrValue
		case "title":
			sid.Title = &attrValue
		}
	}

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

// MarshalXML for ShapeFill
func (sf ShapeFill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "v:fill"
	start.Attr = []xml.Attr{}

	if sf.Type != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: *sf.Type})
	}
	if sf.Color != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "color"}, Value: *sf.Color})
	}
	if sf.Opacity != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "opacity"}, Value: *sf.Opacity})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML for ShapeFill
func (sf *ShapeFill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		attrValue := attr.Value // Create a copy to avoid pointer aliasing
		switch attr.Name.Local {
		case "type":
			sf.Type = &attrValue
		case "color":
			sf.Color = &attrValue
		case "opacity":
			sf.Opacity = &attrValue
		}
	}

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

// MarshalXML for ShapeStroke
func (ss ShapeStroke) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "v:stroke"
	start.Attr = []xml.Attr{}

	if ss.Color != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "color"}, Value: *ss.Color})
	}
	if ss.Weight != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "weight"}, Value: *ss.Weight})
	}
	if ss.DashStyle != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dashstyle"}, Value: *ss.DashStyle})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML for ShapeStroke
func (ss *ShapeStroke) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		attrValue := attr.Value // Create a copy to avoid pointer aliasing
		switch attr.Name.Local {
		case "color":
			ss.Color = &attrValue
		case "weight":
			ss.Weight = &attrValue
		case "dashstyle":
			ss.DashStyle = &attrValue
		}
	}

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