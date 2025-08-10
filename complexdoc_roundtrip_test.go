package godocx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestComplexDocRoundtrip tests loading the complexdoc.docx file and saving it
// without any modifications to verify round-trip functionality
func TestComplexDocRoundtrip(t *testing.T) {
	// Load the complex document
	doc, err := OpenDocument("testdata/complexdoc.docx")
	require.NoError(t, err, "Failed to open complexdoc.docx")
	require.NotNil(t, doc, "Document should not be nil")

	// Save the document without any modifications
	outputFile := "testdata/complexdoc_roundtrip_test.docx"
	err = doc.SaveTo(outputFile)
	require.NoError(t, err, "Failed to save test document")

	// Verify that we can re-open the saved document
	mutatedDoc, err := OpenDocument(outputFile)
	require.NoError(t, err, "Failed to re-open test document")
	require.NotNil(t, mutatedDoc, "Re-opened document should not be nil")

	// Basic structure validation - ensure the document has content
	assert.NotNil(t, mutatedDoc.Document, "Document should have content")
	if mutatedDoc.Document != nil && mutatedDoc.Document.Body != nil {
		assert.NotEmpty(t, mutatedDoc.Document.Body.Children, "Document body should have children")
	}

	t.Log("Successfully completed round-trip test for complexdoc.docx")
}