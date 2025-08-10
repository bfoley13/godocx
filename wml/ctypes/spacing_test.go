package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/bfoley13/godocx/internal"
	"github.com/bfoley13/godocx/wml/stypes"
)

func TestSpacing_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Spacing
		expected string
	}{
		{
			name: "All fields set",
			input: Spacing{
				Before:            internal.ToPtr(uint64(120)),
				After:             internal.ToPtr(uint64(240)),
				BeforeLines:       internal.ToPtr(10),
				BeforeAutospacing: internal.ToPtr(stypes.OnOffOn),
				AfterAutospacing:  internal.ToPtr(stypes.OnOffOff),
				Line:              internal.ToPtr(360),
				LineRule:          internal.ToPtr(stypes.LineSpacingRuleExact),
			},
			expected: `<w:spacing w:before="120" w:after="240" w:beforeLines="10" w:beforeAutospacing="on" w:afterAutospacing="off" w:line="360" w:lineRule="exact"></w:spacing>`,
		},
		{
			name: "Some fields set",
			input: Spacing{
				Before:   internal.ToPtr(uint64(120)),
				After:    internal.ToPtr(uint64(240)),
				Line:     internal.ToPtr(360),
				LineRule: internal.ToPtr(stypes.LineSpacingRuleAuto),
			},
			expected: `<w:spacing w:before="120" w:after="240" w:line="360" w:lineRule="auto"></w:spacing>`,
		},
		{
			name:     "No fields set",
			input:    Spacing{},
			expected: `<w:spacing></w:spacing>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: "w:spacing"}}
			if err := tt.input.MarshalXML(encoder, start); err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			// Finalize encoding
			encoder.Flush()

			got := strings.TrimSpace(result.String())
			if got != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, got)
			}
		})
	}
}

func TestSpacing_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xmlInput string
		expected Spacing
	}{
		{
			name:     "All fields set",
			xmlInput: `<w:spacing w:before="120" w:after="240" w:beforeLines="10" w:beforeAutospacing="on" w:afterAutospacing="off" w:line="360" w:lineRule="exact"></w:spacing>`,
			expected: Spacing{
				Before:            internal.ToPtr(uint64(120)),
				After:             internal.ToPtr(uint64(240)),
				BeforeLines:       internal.ToPtr(10),
				BeforeAutospacing: internal.ToPtr(stypes.OnOffOn),
				AfterAutospacing:  internal.ToPtr(stypes.OnOffOff),
				Line:              internal.ToPtr(360),
				LineRule:          internal.ToPtr(stypes.LineSpacingRuleExact),
			},
		},
		{
			name:     "Some fields set",
			xmlInput: `<w:spacing w:before="120" w:after="240" w:line="360" w:lineRule="auto"></w:spacing>`,
			expected: Spacing{
				Before:   internal.ToPtr(uint64(120)),
				After:    internal.ToPtr(uint64(240)),
				Line:     internal.ToPtr(360),
				LineRule: internal.ToPtr(stypes.LineSpacingRuleAuto),
			},
		},
		{
			name:     "No fields set",
			xmlInput: `<w:spacing></w:spacing>`,
			expected: Spacing{},
		},
		{
			name:     "Only before and after spacing",
			xmlInput: `<w:spacing w:before="240" w:after="120"></w:spacing>`,
			expected: Spacing{
				Before: internal.ToPtr(uint64(240)),
				After:  internal.ToPtr(uint64(120)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var spacing Spacing

			err := xml.Unmarshal([]byte(tt.xmlInput), &spacing)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare all fields
			if (spacing.Before == nil) != (tt.expected.Before == nil) ||
				(spacing.Before != nil && tt.expected.Before != nil && *spacing.Before != *tt.expected.Before) {
				t.Errorf("Before mismatch. Expected: %v, Got: %v", 
					tt.expected.Before, spacing.Before)
			}

			if (spacing.After == nil) != (tt.expected.After == nil) ||
				(spacing.After != nil && tt.expected.After != nil && *spacing.After != *tt.expected.After) {
				t.Errorf("After mismatch. Expected: %v, Got: %v", 
					tt.expected.After, spacing.After)
			}

			if (spacing.BeforeLines == nil) != (tt.expected.BeforeLines == nil) ||
				(spacing.BeforeLines != nil && tt.expected.BeforeLines != nil && *spacing.BeforeLines != *tt.expected.BeforeLines) {
				t.Errorf("BeforeLines mismatch. Expected: %v, Got: %v", 
					tt.expected.BeforeLines, spacing.BeforeLines)
			}

			if (spacing.BeforeAutospacing == nil) != (tt.expected.BeforeAutospacing == nil) ||
				(spacing.BeforeAutospacing != nil && tt.expected.BeforeAutospacing != nil && *spacing.BeforeAutospacing != *tt.expected.BeforeAutospacing) {
				t.Errorf("BeforeAutospacing mismatch. Expected: %v, Got: %v", 
					tt.expected.BeforeAutospacing, spacing.BeforeAutospacing)
			}

			if (spacing.AfterAutospacing == nil) != (tt.expected.AfterAutospacing == nil) ||
				(spacing.AfterAutospacing != nil && tt.expected.AfterAutospacing != nil && *spacing.AfterAutospacing != *tt.expected.AfterAutospacing) {
				t.Errorf("AfterAutospacing mismatch. Expected: %v, Got: %v", 
					tt.expected.AfterAutospacing, spacing.AfterAutospacing)
			}

			if (spacing.Line == nil) != (tt.expected.Line == nil) ||
				(spacing.Line != nil && tt.expected.Line != nil && *spacing.Line != *tt.expected.Line) {
				t.Errorf("Line mismatch. Expected: %v, Got: %v", 
					tt.expected.Line, spacing.Line)
			}

			if (spacing.LineRule == nil) != (tt.expected.LineRule == nil) ||
				(spacing.LineRule != nil && tt.expected.LineRule != nil && *spacing.LineRule != *tt.expected.LineRule) {
				t.Errorf("LineRule mismatch. Expected: %v, Got: %v", 
					tt.expected.LineRule, spacing.LineRule)
			}
		})
	}
}

func TestSpacing_RoundTrip(t *testing.T) {
	original := Spacing{
		Before:            internal.ToPtr(uint64(120)),
		After:             internal.ToPtr(uint64(240)),
		BeforeLines:       internal.ToPtr(10),
		BeforeAutospacing: internal.ToPtr(stypes.OnOffOn),
		AfterAutospacing:  internal.ToPtr(stypes.OnOffOff),
		Line:              internal.ToPtr(360),
		LineRule:          internal.ToPtr(stypes.LineSpacingRuleExact),
	}

	// Marshal to XML
	var buf strings.Builder
	encoder := xml.NewEncoder(&buf)
	start := xml.StartElement{Name: xml.Name{Local: "w:spacing"}}
	err := original.MarshalXML(encoder, start)
	if err != nil {
		t.Fatalf("Error marshaling: %v", err)
	}
	encoder.Flush()

	// Unmarshal back
	var unmarshaled Spacing
	err = xml.Unmarshal([]byte(buf.String()), &unmarshaled)
	if err != nil {
		t.Fatalf("Error unmarshaling: %v", err)
	}

	// Compare
	if *unmarshaled.Before != *original.Before {
		t.Errorf("Before round-trip mismatch. Expected: %d, Got: %d", *original.Before, *unmarshaled.Before)
	}
	if *unmarshaled.After != *original.After {
		t.Errorf("After round-trip mismatch. Expected: %d, Got: %d", *original.After, *unmarshaled.After)
	}
	if *unmarshaled.Line != *original.Line {
		t.Errorf("Line round-trip mismatch. Expected: %d, Got: %d", *original.Line, *unmarshaled.Line)
	}
	if *unmarshaled.LineRule != *original.LineRule {
		t.Errorf("LineRule round-trip mismatch. Expected: %s, Got: %s", *original.LineRule, *unmarshaled.LineRule)
	}
}
