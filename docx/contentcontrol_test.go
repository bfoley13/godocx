package docx

import (
	"testing"
	"time"

	"github.com/bfoley13/godocx/wml/stypes"
	"github.com/stretchr/testify/assert"
)

func TestAddTextContentControl(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Add text content control
	cc := doc.AddTextContentControl("Name Field", "name", "John Doe", false)

	assert.NotNil(t, cc)
	assert.Equal(t, "Name Field", cc.GetAlias())
	assert.Equal(t, "name", cc.GetTag())
	assert.Equal(t, "John Doe", cc.GetText())
}

func TestAddTextContentControlMultiline(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Add multiline text content control
	cc := doc.AddTextContentControl("Comments", "comments", "Line 1\nLine 2", true)

	assert.NotNil(t, cc)
	assert.Equal(t, "Comments", cc.GetAlias())
	assert.Equal(t, "comments", cc.GetTag())
	assert.Equal(t, "Line 1\nLine 2", cc.GetText())
}

func TestAddComboBoxContentControl(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Create list items
	items := []ContentControlListItem{
		{DisplayText: "Option 1", Value: "opt1"},
		{DisplayText: "Option 2", Value: "opt2"},
		{DisplayText: "Option 3", Value: "opt3"},
	}

	cc := doc.AddComboBoxContentControl("Options", "options", items, "opt1")

	assert.NotNil(t, cc)
	assert.Equal(t, "Options", cc.GetAlias())
	assert.Equal(t, "options", cc.GetTag())
	assert.Equal(t, "opt1", cc.GetText())
}

func TestAddDropDownContentControl(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Create list items
	items := []ContentControlListItem{
		{DisplayText: "Choice A", Value: "a"},
		{DisplayText: "Choice B", Value: "b"},
		{DisplayText: "Choice C", Value: "c"},
	}

	cc := doc.AddDropDownContentControl("Choices", "choice", items, "b")

	assert.NotNil(t, cc)
	assert.Equal(t, "Choices", cc.GetAlias())
	assert.Equal(t, "choice", cc.GetTag())
	assert.Equal(t, "b", cc.GetText())
}

func TestAddDateContentControl(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Create a specific date
	testDate := time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)

	cc := doc.AddDateContentControl("Birth Date", "birthdate", &testDate, "yyyy-MM-dd")

	assert.NotNil(t, cc)
	assert.Equal(t, "Birth Date", cc.GetAlias())
	assert.Equal(t, "birthdate", cc.GetTag())
	// The text should contain the formatted date
	assert.Contains(t, cc.GetText(), "2024")
}

func TestAddCheckboxContentControl(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Add checked checkbox
	cc := doc.AddCheckboxContentControl("Agree", "agree", true)

	assert.NotNil(t, cc)
	assert.Equal(t, "Agree", cc.GetAlias())
	assert.Equal(t, "agree", cc.GetTag())
	// Should contain checked symbol
	assert.Contains(t, cc.GetText(), "☑")

	// Add unchecked checkbox
	cc2 := doc.AddCheckboxContentControl("Disagree", "disagree", false)

	assert.NotNil(t, cc2)
	assert.Equal(t, "Disagree", cc2.GetAlias())
	assert.Equal(t, "disagree", cc2.GetTag())
	// Should contain unchecked symbol
	assert.Contains(t, cc2.GetText(), "☐")
}

func TestContentControlSetText(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	cc := doc.AddTextContentControl("Test", "test", "Original", false)
	assert.Equal(t, "Original", cc.GetText())

	// Change the text
	cc.SetText("Modified")
	assert.Equal(t, "Modified", cc.GetText())
}

func TestContentControlSetLock(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	cc := doc.AddTextContentControl("Test", "test", "Text", false)

	// Set lock to content locked
	cc.SetLock(stypes.SdtLockContentLocked)

	// Verify lock is set (would need to check internal structure)
	assert.NotNil(t, cc.sdt.Properties.Lock)
}

func TestContentControlSetTemporary(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	cc := doc.AddTextContentControl("Test", "test", "Text", false)

	// Set as temporary
	cc.SetTemporary(true)

	// Verify temporary is set
	assert.NotNil(t, cc.sdt.Properties.Temporary)
}

func TestReplaceAllInContentControls(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Add regular paragraph
	doc.AddParagraph("Hello World in paragraph")

	// Add text content control with text to replace
	doc.AddTextContentControl("Name", "name", "Hello World in content control", false)

	// Add combo box with text to replace
	items := []ContentControlListItem{
		{DisplayText: "World Option", Value: "world"},
	}
	doc.AddComboBoxContentControl("Options", "options", items, "Hello World in combo")

	// Test replacement
	replacements := doc.ReplaceAll("World", "Universe")

	// Should have replaced 3 occurrences (1 in paragraph + 1 in text control + 1 in combo box)
	assert.Equal(t, 3, replacements)
}

func TestAddBasicContentControl(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Test different content control types
	cc1 := doc.AddContentControl("Rich Text", "richtext", ContentControlTypeRichText)
	assert.NotNil(t, cc1)
	assert.Equal(t, "Rich Text", cc1.GetAlias())

	cc2 := doc.AddContentControl("Picture", "pic", ContentControlTypePicture)
	assert.NotNil(t, cc2)
	assert.Equal(t, "Picture", cc2.GetAlias())

	cc3 := doc.AddContentControl("Group", "group", ContentControlTypeGroup)
	assert.NotNil(t, cc3)
	assert.Equal(t, "Group", cc3.GetAlias())
}

func TestContentControlEmptyText(t *testing.T) {
	doc := NewRootDoc()
	doc.Document = &Document{Root: doc, Body: NewBody(doc)}

	// Add content control without initial text
	cc := doc.AddTextContentControl("Empty", "empty", "", false)

	assert.NotNil(t, cc)
	assert.Equal(t, "Empty", cc.GetAlias())
	assert.Equal(t, "", cc.GetText())

	// Add text later
	cc.SetText("Now has text")
	assert.Equal(t, "Now has text", cc.GetText())
}