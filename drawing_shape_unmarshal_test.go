package godocx

import (
	"encoding/xml"
	"testing"

	"github.com/bfoley13/godocx/dml"
	"github.com/bfoley13/godocx/dml/dmlpic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGraphicUnmarshalXML tests that Graphic can unmarshal from XML
func TestGraphicUnmarshalXML(t *testing.T) {
	xmlData := `<a:graphic xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
		<a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/picture">
			<pic:pic xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture"/>
		</a:graphicData>
	</a:graphic>`

	var graphic dml.Graphic
	err := xml.Unmarshal([]byte(xmlData), &graphic)
	require.NoError(t, err, "Should unmarshal without error")
	
	assert.NotNil(t, graphic.Data, "GraphicData should be populated")
	assert.Equal(t, "http://schemas.openxmlformats.org/drawingml/2006/picture", graphic.Data.URI)
	assert.NotNil(t, graphic.Data.Pic, "Pic should be populated")
}

// TestBlipUnmarshalXML tests that Blip can unmarshal from XML
func TestBlipUnmarshalXML(t *testing.T) {
	xmlData := `<a:blip r:embed="rId1" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"/>`

	var blip dmlpic.Blip
	err := xml.Unmarshal([]byte(xmlData), &blip)
	require.NoError(t, err, "Should unmarshal without error")
	
	assert.Equal(t, "rId1", blip.EmbedID, "EmbedID should be populated")
}

// TestWrapSquareUnmarshalXML tests that WrapSquare can unmarshal from XML
func TestWrapSquareUnmarshalXML(t *testing.T) {
	xmlData := `<wp:wrapSquare wrapText="bothSides" distT="0" distB="0" distL="114300" distR="114300" xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing">
		<wp:effectExtent l="0" t="0" r="0" b="0"/>
	</wp:wrapSquare>`

	var wrapSquare dml.WrapSquare
	err := xml.Unmarshal([]byte(xmlData), &wrapSquare)
	require.NoError(t, err, "Should unmarshal without error")
	
	assert.Equal(t, "bothSides", string(wrapSquare.WrapText), "WrapText should be populated")
	assert.NotNil(t, wrapSquare.DistL, "DistL should be populated")
	assert.Equal(t, uint(114300), *wrapSquare.DistL, "DistL should have correct value")
	assert.NotNil(t, wrapSquare.EffectExtent, "EffectExtent should be populated")
}

// TestDrawingShapeRoundtrip tests marshal/unmarshal roundtrip for drawing and shape elements
func TestDrawingShapeRoundtrip(t *testing.T) {
	// Create a simple drawing with graphics
	graphic := &dml.Graphic{
		Data: &dml.GraphicData{
			URI: "http://schemas.openxmlformats.org/drawingml/2006/picture",
			Pic: &dmlpic.Pic{
				BlipFill: dmlpic.BlipFill{
					Blip: &dmlpic.Blip{
						EmbedID: "rId1",
					},
				},
			},
		},
	}

	// Marshal to XML
	xmlData, err := xml.Marshal(graphic)
	require.NoError(t, err, "Should marshal without error")

	// Unmarshal back from XML
	var unmarshaledGraphic dml.Graphic
	err = xml.Unmarshal(xmlData, &unmarshaledGraphic)
	require.NoError(t, err, "Should unmarshal without error")

	// Verify the roundtrip preserved data
	assert.NotNil(t, unmarshaledGraphic.Data, "GraphicData should be preserved")
	assert.Equal(t, graphic.Data.URI, unmarshaledGraphic.Data.URI, "URI should be preserved")
	assert.NotNil(t, unmarshaledGraphic.Data.Pic, "Pic should be preserved")
	assert.NotNil(t, unmarshaledGraphic.Data.Pic.BlipFill.Blip, "Blip should be preserved")
	assert.Equal(t, "rId1", unmarshaledGraphic.Data.Pic.BlipFill.Blip.EmbedID, "EmbedID should be preserved")
}

// TestBlipFillRoundtrip tests marshal/unmarshal roundtrip for BlipFill
func TestBlipFillRoundtrip(t *testing.T) {
	// Create a BlipFill with attributes
	dpi := uint32(96)
	rotWithShape := true
	
	blipFill := &dmlpic.BlipFill{
		DPI:          &dpi,
		RotWithShape: &rotWithShape,
		Blip: &dmlpic.Blip{
			EmbedID: "rId2",
		},
	}

	// Marshal to XML
	xmlData, err := xml.Marshal(blipFill)
	require.NoError(t, err, "Should marshal without error")
	
	t.Logf("Marshaled XML: %s", string(xmlData))

	// Unmarshal back from XML
	var unmarshaledBlipFill dmlpic.BlipFill
	err = xml.Unmarshal(xmlData, &unmarshaledBlipFill)
	require.NoError(t, err, "Should unmarshal without error")

	// Verify the roundtrip preserved data
	assert.NotNil(t, unmarshaledBlipFill.DPI, "DPI should be preserved")
	assert.Equal(t, uint32(96), *unmarshaledBlipFill.DPI, "DPI value should be preserved")
	assert.NotNil(t, unmarshaledBlipFill.RotWithShape, "RotWithShape should be preserved")
	assert.Equal(t, true, *unmarshaledBlipFill.RotWithShape, "RotWithShape value should be preserved")
	assert.NotNil(t, unmarshaledBlipFill.Blip, "Blip should be preserved")
	assert.Equal(t, "rId2", unmarshaledBlipFill.Blip.EmbedID, "EmbedID should be preserved")
}

// TestNonVisualGraphicFramePropRoundtrip tests marshal/unmarshal roundtrip for NonVisualGraphicFrameProp
func TestNonVisualGraphicFramePropRoundtrip(t *testing.T) {
	// Create NonVisualGraphicFrameProp
	nvGraphicFrameProp := &dml.NonVisualGraphicFrameProp{
		GraphicFrameLocks: &dml.GraphicFrameLocks{},
	}

	// Marshal to XML
	xmlData, err := xml.Marshal(nvGraphicFrameProp)
	require.NoError(t, err, "Should marshal without error")

	// Unmarshal back from XML
	var unmarshaledProp dml.NonVisualGraphicFrameProp
	err = xml.Unmarshal(xmlData, &unmarshaledProp)
	require.NoError(t, err, "Should unmarshal without error")

	// Verify the roundtrip preserved data
	assert.NotNil(t, unmarshaledProp.GraphicFrameLocks, "GraphicFrameLocks should be preserved")
}