package godocx

import "github.com/bfoley13/godocx/wml/ctypes"

// countGraphicsInNestedTable counts graphics in a nested ctypes.Table
func countGraphicsInNestedTable(table *ctypes.Table) int {
	count := 0
	
	for _, rowContent := range table.RowContents {
		if rowContent.Row != nil {
			for _, cellContent := range rowContent.Row.Contents {
				if cellContent.Cell != nil {
					for _, content := range cellContent.Cell.Contents {
						if content.Paragraph != nil {
							count += countGraphicsInCTypeParagraph(content.Paragraph)
						} else if content.Table != nil {
							count += countGraphicsInNestedTable(content.Table)
						}
					}
				}
			}
		}
	}
	
	return count
}

// extractGraphicsFromNestedTable extracts graphics from a nested ctypes.Table
func extractGraphicsFromNestedTable(table *ctypes.Table) []GraphicsInfo {
	var graphics []GraphicsInfo
	
	for _, rowContent := range table.RowContents {
		if rowContent.Row != nil {
			for _, cellContent := range rowContent.Row.Contents {
				if cellContent.Cell != nil {
					for _, content := range cellContent.Cell.Contents {
						if content.Paragraph != nil {
							graphics = append(graphics, extractGraphicsFromCTypeParagraph(content.Paragraph)...)
						} else if content.Table != nil {
							graphics = append(graphics, extractGraphicsFromNestedTable(content.Table)...)
						}
					}
				}
			}
		}
	}
	
	return graphics
}