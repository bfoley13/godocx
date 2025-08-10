package godocx

import (
	"encoding/xml"
	"testing"

	"github.com/bfoley13/godocx/docx"
	"github.com/bfoley13/godocx/wml/ctypes"
	"github.com/bfoley13/godocx/wml/stypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestBackgroundMarshalUnmarshalConsistency verifies that Background marshal/unmarshal is consistent
func TestBackgroundMarshalUnmarshalConsistency(t *testing.T) {
	tests := []struct {
		name       string
		background *docx.Background
	}{
		{
			name: "All attributes set",
			background: &docx.Background{
				Color:      stringPtr("FF0000"),
				ThemeColor: themeColorPtr(stypes.ThemeColorAccent1),
				ThemeTint:  stringPtr("80"),
				ThemeShade: stringPtr("50"),
			},
		},
		{
			name: "Only color set",
			background: &docx.Background{
				Color: stringPtr("00FF00"),
			},
		},
		{
			name: "Theme color and tint",
			background: &docx.Background{
				ThemeColor: themeColorPtr(stypes.ThemeColorBackground2),
				ThemeTint:  stringPtr("90"),
			},
		},
		{
			name: "Theme color and shade",
			background: &docx.Background{
				ThemeColor: themeColorPtr(stypes.ThemeColorAccent2),
				ThemeShade: stringPtr("25"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal to XML
			xmlData, err := xml.Marshal(tt.background)
			require.NoError(t, err, "Should marshal without error")
			
			t.Logf("Marshaled XML: %s", string(xmlData))

			// Unmarshal back from XML
			var roundtrip docx.Background
			err = xml.Unmarshal(xmlData, &roundtrip)
			require.NoError(t, err, "Should unmarshal without error")

			// Verify all properties are preserved
			if tt.background.Color != nil {
				require.NotNil(t, roundtrip.Color, "Color should not be nil")
				assert.Equal(t, *tt.background.Color, *roundtrip.Color, "Color should be preserved")
			} else {
				assert.Nil(t, roundtrip.Color, "Color should remain nil")
			}

			if tt.background.ThemeColor != nil {
				require.NotNil(t, roundtrip.ThemeColor, "ThemeColor should not be nil")
				assert.Equal(t, *tt.background.ThemeColor, *roundtrip.ThemeColor, "ThemeColor should be preserved")
			} else {
				assert.Nil(t, roundtrip.ThemeColor, "ThemeColor should remain nil")
			}

			if tt.background.ThemeTint != nil {
				require.NotNil(t, roundtrip.ThemeTint, "ThemeTint should not be nil")
				assert.Equal(t, *tt.background.ThemeTint, *roundtrip.ThemeTint, "ThemeTint should be preserved")
			} else {
				assert.Nil(t, roundtrip.ThemeTint, "ThemeTint should remain nil")
			}

			if tt.background.ThemeShade != nil {
				require.NotNil(t, roundtrip.ThemeShade, "ThemeShade should not be nil")
				assert.Equal(t, *tt.background.ThemeShade, *roundtrip.ThemeShade, "ThemeShade should be preserved")
			} else {
				assert.Nil(t, roundtrip.ThemeShade, "ThemeShade should remain nil")
			}
		})
	}
}

// TestColorMarshalUnmarshalConsistency verifies that Color marshal/unmarshal is consistent
func TestColorMarshalUnmarshalConsistency(t *testing.T) {
	tests := []struct {
		name  string
		color *ctypes.Color
	}{
		{
			name: "All attributes set",
			color: &ctypes.Color{
				Val:        "FF0000",
				ThemeColor: themeColorPtr(stypes.ThemeColorAccent1),
				ThemeTint:  stringPtr("80"),
				ThemeShade: stringPtr("50"),
			},
		},
		{
			name: "Only val set",
			color: &ctypes.Color{
				Val: "00FF00",
			},
		},
		{
			name: "Val and theme color",
			color: &ctypes.Color{
				Val:        "0000FF",
				ThemeColor: themeColorPtr(stypes.ThemeColorBackground1),
			},
		},
		{
			name: "Theme color with tint",
			color: &ctypes.Color{
				Val:        "FFFFFF",
				ThemeColor: themeColorPtr(stypes.ThemeColorAccent3),
				ThemeTint:  stringPtr("60"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal to XML
			xmlData, err := xml.Marshal(tt.color)
			require.NoError(t, err, "Should marshal without error")
			
			t.Logf("Marshaled XML: %s", string(xmlData))

			// Unmarshal back from XML
			var roundtrip ctypes.Color
			err = xml.Unmarshal(xmlData, &roundtrip)
			require.NoError(t, err, "Should unmarshal without error")

			// Verify all properties are preserved
			assert.Equal(t, tt.color.Val, roundtrip.Val, "Val should be preserved")

			if tt.color.ThemeColor != nil {
				require.NotNil(t, roundtrip.ThemeColor, "ThemeColor should not be nil")
				assert.Equal(t, *tt.color.ThemeColor, *roundtrip.ThemeColor, "ThemeColor should be preserved")
			} else {
				assert.Nil(t, roundtrip.ThemeColor, "ThemeColor should remain nil")
			}

			if tt.color.ThemeTint != nil {
				require.NotNil(t, roundtrip.ThemeTint, "ThemeTint should not be nil")
				assert.Equal(t, *tt.color.ThemeTint, *roundtrip.ThemeTint, "ThemeTint should be preserved")
			} else {
				assert.Nil(t, roundtrip.ThemeTint, "ThemeTint should remain nil")
			}

			if tt.color.ThemeShade != nil {
				require.NotNil(t, roundtrip.ThemeShade, "ThemeShade should not be nil")
				assert.Equal(t, *tt.color.ThemeShade, *roundtrip.ThemeShade, "ThemeShade should be preserved")
			} else {
				assert.Nil(t, roundtrip.ThemeShade, "ThemeShade should remain nil")
			}
		})
	}
}

// TestDocumentBackgroundRoundtrip tests that document backgrounds work in full document context
func TestDocumentBackgroundRoundtrip(t *testing.T) {
	// Create a document with a background
	doc, err := NewDocument()
	require.NoError(t, err, "Should create document without error")
	doc.Document.Background = &docx.Background{
		Color:      stringPtr("E0E0E0"),
		ThemeColor: themeColorPtr(stypes.ThemeColorAccent4),
		ThemeTint:  stringPtr("75"),
	}

	// Marshal the entire document
	xmlData, err := xml.Marshal(doc.Document)
	require.NoError(t, err, "Should marshal document without error")
	
	t.Logf("Document XML contains background: %s", string(xmlData))

	// Unmarshal back to a new document
	var roundtripDoc docx.Document
	err = xml.Unmarshal(xmlData, &roundtripDoc)
	require.NoError(t, err, "Should unmarshal document without error")

	// Verify the background was preserved
	require.NotNil(t, roundtripDoc.Background, "Background should be preserved")
	require.NotNil(t, roundtripDoc.Background.Color, "Background color should be preserved")
	assert.Equal(t, "E0E0E0", *roundtripDoc.Background.Color, "Background color value should be preserved")
	
	require.NotNil(t, roundtripDoc.Background.ThemeColor, "Background theme color should be preserved")
	assert.Equal(t, stypes.ThemeColorAccent4, *roundtripDoc.Background.ThemeColor, "Background theme color value should be preserved")
	
	require.NotNil(t, roundtripDoc.Background.ThemeTint, "Background theme tint should be preserved")
	assert.Equal(t, "75", *roundtripDoc.Background.ThemeTint, "Background theme tint value should be preserved")
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func themeColorPtr(tc stypes.ThemeColor) *stypes.ThemeColor {
	return &tc
}