package docx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceAll_BasicReplacement(t *testing.T) {
	// Create a minimal document structure for testing
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Add some text
	doc.AddParagraph("Hello World! This World is beautiful.")
	doc.AddParagraph("Another World paragraph.")

	// Test replacement
	replacements := doc.ReplaceAll("World", "Universe")

	// Should have replaced 3 occurrences
	assert.Equal(t, 3, replacements)
}

func TestReplaceAll_EmptyString(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	doc.AddParagraph("Hello World!")

	// Test with empty oldText - should return 0
	replacements := doc.ReplaceAll("", "Universe")
	assert.Equal(t, 0, replacements)
}

func TestReplaceAll_NoMatches(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	doc.AddParagraph("Hello World!")

	// Test with non-existent text
	replacements := doc.ReplaceAll("Mars", "Universe")
	assert.Equal(t, 0, replacements)
}

func TestReplaceAll_FormattedText(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Add formatted text
	para := doc.AddParagraph("Hello ")
	para.AddText("World").Bold(true)
	para.AddText("! Welcome to this World.")

	// Test replacement - should work across runs
	replacements := doc.ReplaceAll("World", "Universe")

	// Should have replaced 2 occurrences
	assert.Equal(t, 2, replacements)
}

func TestReplaceAll_WithTable(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Add paragraph text
	doc.AddParagraph("Hello World!")

	// Add table with text
	table := doc.AddTable()
	row := table.AddRow()
	cell := row.AddCell()
	cell.AddParagraph("World in table")

	// Another row
	row2 := table.AddRow()
	cell2 := row2.AddCell()
	cell2.AddParagraph("Another World cell")

	// Test replacement
	replacements := doc.ReplaceAll("World", "Universe")

	// Should have replaced 3 occurrences (1 in paragraph + 2 in table)
	assert.Equal(t, 3, replacements)
}

func TestReplaceAll_WithHyperlinks(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Add paragraph with hyperlink
	para := doc.AddParagraph("Visit ")
	para.AddLink("World Wide Web", "https://www.example.com")
	para.AddText(" for more World information.")

	// Test replacement
	replacements := doc.ReplaceAll("World", "Universe")

	// Should have replaced 2 occurrences
	assert.Equal(t, 2, replacements)
}

func TestReplaceAll_CaseSensitive(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	doc.AddParagraph("Hello World! Hello world!")

	// Test case-sensitive replacement
	replacements := doc.ReplaceAll("World", "Universe")

	// Should only replace the capitalized "World"
	assert.Equal(t, 1, replacements)
}

func TestReplaceAll_OverlappingReplacements(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	doc.AddParagraph("aaabaaab")

	// Test overlapping pattern - should replace non-overlapping occurrences
	replacements := doc.ReplaceAll("aaa", "bb")

	// Should replace twice (aaabaaab -> bbabbbab, first "aaa" at start, second "aaa" at end)
	assert.Equal(t, 2, replacements)
}