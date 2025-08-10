package dml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShapeBasicMarshal(t *testing.T) {
	shape := NewShape("_x0000_s1025", "#_x0000_t75").
		WithStyle("position:absolute;left:0;text-align:left;margin-left:0;margin-top:0;width:100pt;height:75pt;z-index:1").
		WithFillColor("#ff0000").
		WithStrokeColor("#000000").
		WithStrokeWeight("1pt")

	xmlData, err := xml.Marshal(shape)
	require.NoError(t, err, "Should marshal without error")

	expectedContent := []string{
		`v:shape`,
		`xmlns:v="urn:schemas-microsoft-com:vml"`,
		`xmlns:o="urn:schemas-microsoft-com:office:office"`,
		`id="_x0000_s1025"`,
		`type="#_x0000_t75"`,
		`fillcolor="#ff0000"`,
		`strokecolor="#000000"`,
		`strokeweight="1pt"`,
	}

	xmlString := string(xmlData)
	for _, expected := range expectedContent {
		assert.Contains(t, xmlString, expected, "XML should contain expected content")
	}
}

func TestShapeWithPath(t *testing.T) {
	shape := NewShape("rect1", "#rectangle").
		WithPath("m 0,0 l 100,0 l 100,50 l 0,50 x e")

	xmlData, err := xml.Marshal(shape)
	require.NoError(t, err, "Should marshal without error")

	assert.Contains(t, string(xmlData), `<v:path v="m 0,0 l 100,0 l 100,50 l 0,50 x e"></v:path>`)
}

func TestShapeWithTextBox(t *testing.T) {
	shape := NewShape("text1", "#textbox").
		WithTextBox("<w:p><w:r><w:t>Hello World</w:t></w:r></w:p>")

	xmlData, err := xml.Marshal(shape)
	require.NoError(t, err, "Should marshal without error")

	assert.Contains(t, string(xmlData), `<v:textbox><w:p><w:r><w:t>Hello World</w:t></w:r></w:p></v:textbox>`)
}

func TestShapeWithImageData(t *testing.T) {
	shape := NewShape("img1", "#image").
		WithImageData("rId1")

	xmlData, err := xml.Marshal(shape)
	require.NoError(t, err, "Should marshal without error")

	assert.Contains(t, string(xmlData), `<v:imagedata r:id="rId1"></v:imagedata>`)
}

func TestShapeUnmarshal(t *testing.T) {
	xmlData := `<v:shape xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" id="rect1" type="#rectangle" style="position:absolute;width:100pt;height:50pt" fillcolor="#ff0000" strokecolor="#000000" strokeweight="1pt">
		<v:path v="m 0,0 l 100,0 l 100,50 l 0,50 x e"></v:path>
		<v:textbox><w:p><w:r><w:t>Test Text</w:t></w:r></w:p></v:textbox>
		<v:imagedata r:id="rId1" title="Test Image"></v:imagedata>
		<v:fill type="solid" color="#ff0000" opacity="50%"></v:fill>
		<v:stroke color="#000000" weight="1pt" dashstyle="solid"></v:stroke>
	</v:shape>`

	var shape Shape
	err := xml.Unmarshal([]byte(xmlData), &shape)
	require.NoError(t, err, "Should unmarshal without error")

	// Test basic attributes
	assert.NotNil(t, shape.ID)
	assert.Equal(t, "rect1", *shape.ID)
	assert.NotNil(t, shape.Type)
	assert.Equal(t, "#rectangle", *shape.Type)
	assert.NotNil(t, shape.Style)
	assert.Contains(t, *shape.Style, "position:absolute")
	assert.NotNil(t, shape.FillColor)
	assert.Equal(t, "#ff0000", *shape.FillColor)
	assert.NotNil(t, shape.StrokeColor)
	assert.Equal(t, "#000000", *shape.StrokeColor)
	assert.NotNil(t, shape.StrokeWeight)
	assert.Equal(t, "1pt", *shape.StrokeWeight)

	// Test child elements
	assert.NotNil(t, shape.Path)
	assert.NotNil(t, shape.Path.V)
	assert.Equal(t, "m 0,0 l 100,0 l 100,50 l 0,50 x e", *shape.Path.V)

	assert.NotNil(t, shape.TextBox)
	assert.Contains(t, shape.TextBox.Content, "Test Text")

	assert.NotNil(t, shape.ImageData)
	assert.NotNil(t, shape.ImageData.ID)
	assert.Equal(t, "rId1", *shape.ImageData.ID)
	assert.NotNil(t, shape.ImageData.Title)
	assert.Equal(t, "Test Image", *shape.ImageData.Title)

	assert.NotNil(t, shape.Fill)
	assert.NotNil(t, shape.Fill.Type)
	assert.Equal(t, "solid", *shape.Fill.Type)
	assert.NotNil(t, shape.Fill.Color)
	assert.Equal(t, "#ff0000", *shape.Fill.Color)

	assert.NotNil(t, shape.Stroke)
	assert.NotNil(t, shape.Stroke.Color)
	assert.Equal(t, "#000000", *shape.Stroke.Color)
	assert.NotNil(t, shape.Stroke.Weight)
	assert.Equal(t, "1pt", *shape.Stroke.Weight)
}

func TestShapeRoundtrip(t *testing.T) {
	original := NewShape("test1", "#rect").
		WithStyle("position:absolute;width:100pt;height:50pt").
		WithFillColor("#ff0000").
		WithStrokeColor("#000000").
		WithStrokeWeight("1pt").
		WithPath("m 0,0 l 100,0 l 100,50 l 0,50 x e").
		WithTextBox("<w:p><w:r><w:t>Test</w:t></w:r></w:p>")

	// Marshal to XML
	xmlData, err := xml.Marshal(original)
	require.NoError(t, err, "Should marshal without error")

	// Unmarshal back from XML
	var unmarshaled Shape
	err = xml.Unmarshal(xmlData, &unmarshaled)
	require.NoError(t, err, "Should unmarshal without error")

	// Verify roundtrip preserved data
	assert.Equal(t, *original.ID, *unmarshaled.ID)
	assert.Equal(t, *original.Type, *unmarshaled.Type)
	assert.Equal(t, *original.FillColor, *unmarshaled.FillColor)
	assert.Equal(t, *original.StrokeColor, *unmarshaled.StrokeColor)
	assert.Equal(t, *original.StrokeWeight, *unmarshaled.StrokeWeight)
	assert.Equal(t, *original.Path.V, *unmarshaled.Path.V)
	assert.Equal(t, original.TextBox.Content, unmarshaled.TextBox.Content)
}

func TestShapePathMarshal(t *testing.T) {
	path := &ShapePath{
		V:           stringPtr("m 0,0 l 100,0 l 100,50 x e"),
		ConnectType: stringPtr("rect"),
		ConnectLocs: stringPtr("0,0;50,25;100,50"),
	}

	xmlData, err := xml.Marshal(path)
	require.NoError(t, err, "Should marshal without error")

	expectedContent := []string{
		`v:path`,
		`v="m 0,0 l 100,0 l 100,50 x e"`,
		`connecttype="rect"`,
		`connectlocs="0,0;50,25;100,50"`,
	}

	xmlString := string(xmlData)
	for _, expected := range expectedContent {
		assert.Contains(t, xmlString, expected, "XML should contain expected content")
	}
}

func TestShapePathUnmarshal(t *testing.T) {
	xmlData := `<v:path v="m 0,0 l 100,0 l 100,50 x e" connecttype="rect" connectlocs="0,0;50,25;100,50"></v:path>`

	var path ShapePath
	err := xml.Unmarshal([]byte(xmlData), &path)
	require.NoError(t, err, "Should unmarshal without error")

	assert.NotNil(t, path.V)
	assert.Equal(t, "m 0,0 l 100,0 l 100,50 x e", *path.V)
	assert.NotNil(t, path.ConnectType)
	assert.Equal(t, "rect", *path.ConnectType)
	assert.NotNil(t, path.ConnectLocs)
	assert.Equal(t, "0,0;50,25;100,50", *path.ConnectLocs)
}

func TestShapeTextBoxMarshal(t *testing.T) {
	textBox := &ShapeTextBox{
		Style:   stringPtr("mso-fit-shape-to-text:t"),
		Content: "<w:p><w:r><w:t>Hello World</w:t></w:r></w:p>",
	}

	xmlData, err := xml.Marshal(textBox)
	require.NoError(t, err, "Should marshal without error")

	xmlString := string(xmlData)
	assert.Contains(t, xmlString, `v:textbox`)
	assert.Contains(t, xmlString, `style="mso-fit-shape-to-text:t"`)
	assert.Contains(t, xmlString, `<w:p><w:r><w:t>Hello World</w:t></w:r></w:p>`)
}

func TestShapeTextBoxUnmarshal(t *testing.T) {
	xmlData := `<v:textbox style="mso-fit-shape-to-text:t"><w:p><w:r><w:t>Hello World</w:t></w:r></w:p></v:textbox>`

	var textBox ShapeTextBox
	err := xml.Unmarshal([]byte(xmlData), &textBox)
	require.NoError(t, err, "Should unmarshal without error")

	assert.NotNil(t, textBox.Style)
	assert.Equal(t, "mso-fit-shape-to-text:t", *textBox.Style)
	assert.Contains(t, textBox.Content, "Hello World")
}

func TestShapeImageDataMarshal(t *testing.T) {
	imageData := &ShapeImageData{
		ID:    stringPtr("rId1"),
		Title: stringPtr("Test Image"),
	}

	xmlData, err := xml.Marshal(imageData)
	require.NoError(t, err, "Should marshal without error")

	xmlString := string(xmlData)
	assert.Contains(t, xmlString, `v:imagedata`)
	assert.Contains(t, xmlString, `r:id="rId1"`)
	assert.Contains(t, xmlString, `title="Test Image"`)
}

func TestShapeImageDataUnmarshal(t *testing.T) {
	xmlData := `<v:imagedata r:id="rId1" title="Test Image"></v:imagedata>`

	var imageData ShapeImageData
	err := xml.Unmarshal([]byte(xmlData), &imageData)
	require.NoError(t, err, "Should unmarshal without error")

	assert.NotNil(t, imageData.ID)
	assert.Equal(t, "rId1", *imageData.ID)
	assert.NotNil(t, imageData.Title)
	assert.Equal(t, "Test Image", *imageData.Title)
}

func TestShapeFillMarshal(t *testing.T) {
	fill := &ShapeFill{
		Type:    stringPtr("solid"),
		Color:   stringPtr("#ff0000"),
		Opacity: stringPtr("50%"),
	}

	xmlData, err := xml.Marshal(fill)
	require.NoError(t, err, "Should marshal without error")

	xmlString := string(xmlData)
	assert.Contains(t, xmlString, `v:fill`)
	assert.Contains(t, xmlString, `type="solid"`)
	assert.Contains(t, xmlString, `color="#ff0000"`)
	assert.Contains(t, xmlString, `opacity="50%"`)
}

func TestShapeFillUnmarshal(t *testing.T) {
	xmlData := `<v:fill type="solid" color="#ff0000" opacity="50%"></v:fill>`

	var fill ShapeFill
	err := xml.Unmarshal([]byte(xmlData), &fill)
	require.NoError(t, err, "Should unmarshal without error")

	assert.NotNil(t, fill.Type)
	assert.Equal(t, "solid", *fill.Type)
	assert.NotNil(t, fill.Color)
	assert.Equal(t, "#ff0000", *fill.Color)
	assert.NotNil(t, fill.Opacity)
	assert.Equal(t, "50%", *fill.Opacity)
}

func TestShapeStrokeMarshal(t *testing.T) {
	stroke := &ShapeStroke{
		Color:     stringPtr("#000000"),
		Weight:    stringPtr("2pt"),
		DashStyle: stringPtr("dash"),
	}

	xmlData, err := xml.Marshal(stroke)
	require.NoError(t, err, "Should marshal without error")

	xmlString := string(xmlData)
	assert.Contains(t, xmlString, `v:stroke`)
	assert.Contains(t, xmlString, `color="#000000"`)
	assert.Contains(t, xmlString, `weight="2pt"`)
	assert.Contains(t, xmlString, `dashstyle="dash"`)
}

func TestShapeStrokeUnmarshal(t *testing.T) {
	xmlData := `<v:stroke color="#000000" weight="2pt" dashstyle="dash"></v:stroke>`

	var stroke ShapeStroke
	err := xml.Unmarshal([]byte(xmlData), &stroke)
	require.NoError(t, err, "Should unmarshal without error")

	assert.NotNil(t, stroke.Color)
	assert.Equal(t, "#000000", *stroke.Color)
	assert.NotNil(t, stroke.Weight)
	assert.Equal(t, "2pt", *stroke.Weight)
	assert.NotNil(t, stroke.DashStyle)
	assert.Equal(t, "dash", *stroke.DashStyle)
}

func TestDrawingWithShape(t *testing.T) {
	drawing := &Drawing{
		Shape: []*Shape{
			NewShape("rect1", "#rectangle").
				WithStyle("position:absolute;width:100pt;height:50pt").
				WithFillColor("#ff0000"),
		},
	}

	xmlData, err := xml.Marshal(drawing)
	require.NoError(t, err, "Should marshal without error")

	xmlString := string(xmlData)
	assert.Contains(t, xmlString, `w:drawing`)
	assert.Contains(t, xmlString, `v:shape`)
	assert.Contains(t, xmlString, `id="rect1"`)
	assert.Contains(t, xmlString, `fillcolor="#ff0000"`)
}

func TestDrawingUnmarshalWithShape(t *testing.T) {
	xmlData := `<w:drawing xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
		<v:shape xmlns:v="urn:schemas-microsoft-com:vml" id="rect1" type="#rectangle" fillcolor="#ff0000">
		</v:shape>
	</w:drawing>`

	var drawing Drawing
	err := xml.Unmarshal([]byte(xmlData), &drawing)
	require.NoError(t, err, "Should unmarshal without error")

	assert.Len(t, drawing.Shape, 1)
	assert.NotNil(t, drawing.Shape[0].ID)
	assert.Equal(t, "rect1", *drawing.Shape[0].ID)
	assert.NotNil(t, drawing.Shape[0].Type)
	assert.Equal(t, "#rectangle", *drawing.Shape[0].Type)
	assert.NotNil(t, drawing.Shape[0].FillColor)
	assert.Equal(t, "#ff0000", *drawing.Shape[0].FillColor)
}

// Helper function for tests
func stringPtr(s string) *string {
	return &s
}