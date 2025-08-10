package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/bfoley13/godocx/wml/stypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun_FieldChar_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		run      Run
		expected string
	}{
		{
			name: "Run with field begin character",
			run: Run{
				Children: []RunChild{
					{
						FldChar: &FieldChar{
							FldCharType: &GenSingleStrVal[stypes.FldCharType]{
								Val: stypes.FldCharTypeBegin,
							},
						},
					},
				},
			},
			expected: `<w:r><w:fldChar w:fldCharType="begin"></w:fldChar></w:r>`,
		},
		{
			name: "Run with field end character",
			run: Run{
				Children: []RunChild{
					{
						FldChar: &FieldChar{
							FldCharType: &GenSingleStrVal[stypes.FldCharType]{
								Val: stypes.FldCharTypeEnd,
							},
						},
					},
				},
			},
			expected: `<w:r><w:fldChar w:fldCharType="end"></w:fldChar></w:r>`,
		},
		{
			name: "Run with field separate character",
			run: Run{
				Children: []RunChild{
					{
						FldChar: &FieldChar{
							FldCharType: &GenSingleStrVal[stypes.FldCharType]{
								Val: stypes.FldCharTypeSeparate,
							},
						},
					},
				},
			},
			expected: `<w:r><w:fldChar w:fldCharType="separate"></w:fldChar></w:r>`,
		},
		{
			name: "Run with field character and text",
			run: Run{
				Children: []RunChild{
					{
						FldChar: &FieldChar{
							FldCharType: &GenSingleStrVal[stypes.FldCharType]{
								Val: stypes.FldCharTypeBegin,
							},
						},
					},
					{
						Text: &Text{Text: "PAGE"},
					},
				},
			},
			expected: `<w:r><w:fldChar w:fldCharType="begin"></w:fldChar><w:t>PAGE</w:t></w:r>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			enc := xml.NewEncoder(&buf)

			err := tt.run.MarshalXML(enc, xml.StartElement{})
			require.NoError(t, err)
			enc.Flush()

			assert.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestRun_FieldChar_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xmlInput string
		expected Run
	}{
		{
			name:     "Run with field begin character",
			xmlInput: `<w:r><w:fldChar w:fldCharType="begin"></w:fldChar></w:r>`,
			expected: Run{
				Children: []RunChild{
					{
						FldChar: &FieldChar{
							FldCharType: &GenSingleStrVal[stypes.FldCharType]{
								Val: stypes.FldCharTypeBegin,
							},
						},
					},
				},
			},
		},
		{
			name:     "Run with field end character",
			xmlInput: `<w:r><w:fldChar w:fldCharType="end"></w:fldChar></w:r>`,
			expected: Run{
				Children: []RunChild{
					{
						FldChar: &FieldChar{
							FldCharType: &GenSingleStrVal[stypes.FldCharType]{
								Val: stypes.FldCharTypeEnd,
							},
						},
					},
				},
			},
		},
		{
			name:     "Run with field separate character",
			xmlInput: `<w:r><w:fldChar w:fldCharType="separate"></w:fldChar></w:r>`,
			expected: Run{
				Children: []RunChild{
					{
						FldChar: &FieldChar{
							FldCharType: &GenSingleStrVal[stypes.FldCharType]{
								Val: stypes.FldCharTypeSeparate,
							},
						},
					},
				},
			},
		},
		{
			name:     "Run with just instruction text",
			xmlInput: `<w:r><w:instrText>PAGE</w:instrText></w:r>`,
			expected: Run{
				Children: []RunChild{
					{
						InstrText: &Text{Text: "PAGE"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var run Run

			err := xml.Unmarshal([]byte(tt.xmlInput), &run)
			require.NoError(t, err)

			assert.Equal(t, len(tt.expected.Children), len(run.Children))
			
			for i, expectedChild := range tt.expected.Children {
				if expectedChild.FldChar != nil {
					require.NotNil(t, run.Children[i].FldChar)
					assert.Equal(t, expectedChild.FldChar.FldCharType.Val, run.Children[i].FldChar.FldCharType.Val)
				}
				if expectedChild.InstrText != nil {
					require.NotNil(t, run.Children[i].InstrText)
					assert.Equal(t, expectedChild.InstrText.Text, run.Children[i].InstrText.Text)
				}
			}
		})
	}
}

func TestRun_FieldRoundTrip(t *testing.T) {
	// Test complete field structure: begin -> instrText -> separate -> text -> end
	originalXML := `<w:r><w:fldChar w:fldCharType="begin"></w:fldChar></w:r><w:r><w:instrText>PAGE</w:instrText></w:r><w:r><w:fldChar w:fldCharType="separate"></w:fldChar></w:r><w:r><w:t>1</w:t></w:r><w:r><w:fldChar w:fldCharType="end"></w:fldChar></w:r>`
	
	// Parse multiple runs (simulating field structure)
	runs := make([]Run, 5)
	xmlParts := []string{
		`<w:r><w:fldChar w:fldCharType="begin"></w:fldChar></w:r>`,
		`<w:r><w:instrText>PAGE</w:instrText></w:r>`,
		`<w:r><w:fldChar w:fldCharType="separate"></w:fldChar></w:r>`,
		`<w:r><w:t>1</w:t></w:r>`,
		`<w:r><w:fldChar w:fldCharType="end"></w:fldChar></w:r>`,
	}
	
	for i, xmlPart := range xmlParts {
		err := xml.Unmarshal([]byte(xmlPart), &runs[i])
		require.NoError(t, err)
	}
	
	// Verify field structure
	assert.NotNil(t, runs[0].Children[0].FldChar)
	assert.Equal(t, stypes.FldCharTypeBegin, runs[0].Children[0].FldChar.FldCharType.Val)
	
	assert.NotNil(t, runs[1].Children[0].InstrText)
	assert.Equal(t, "PAGE", runs[1].Children[0].InstrText.Text)
	
	assert.NotNil(t, runs[2].Children[0].FldChar)
	assert.Equal(t, stypes.FldCharTypeSeparate, runs[2].Children[0].FldChar.FldCharType.Val)
	
	assert.NotNil(t, runs[3].Children[0].Text)
	assert.Equal(t, "1", runs[3].Children[0].Text.Text)
	
	assert.NotNil(t, runs[4].Children[0].FldChar)
	assert.Equal(t, stypes.FldCharTypeEnd, runs[4].Children[0].FldChar.FldCharType.Val)
	
	// Marshal back to XML and verify structure is preserved
	var buf strings.Builder
	enc := xml.NewEncoder(&buf)
	
	for _, run := range runs {
		err := run.MarshalXML(enc, xml.StartElement{})
		require.NoError(t, err)
	}
	enc.Flush()
	
	assert.Equal(t, originalXML, buf.String())
}