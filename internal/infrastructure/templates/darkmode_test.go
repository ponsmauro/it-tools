package templates_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDarkModeColors(t *testing.T) {
	// Path to templates directory relative to this test file
	templatesDir := filepath.Join("..", "..", "..", "static", "templates")

	// Forbidden light mode colors that should have been replaced with CSS variables
	forbiddenColors := []string{
		"#ffffff", // White background
		"#f8fafc", // Light gray background
		"#f5f5f5", // Light gray background (alternate)
		"#f0f0f0", // Light gray background (alternate)
		"#e2e8f0", // Light border
		"#d1d5db", // Light gray border
		"#9ca3af", // Light gray text
		"#374151", // Dark text on light background
		"#111827", // Near-black text (light mode)
		// Note: #1e293b is used as card background in dark mode — not forbidden.
	}

	// Read all files in the templates directory
	files, err := os.ReadDir(templatesDir)
	if err != nil {
		t.Fatalf("Failed to read templates directory: %v", err)
	}

	// Table-driven test structure
	tests := []struct {
		name     string
		filename string
	}{}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".html") {
			tests = append(tests, struct {
				name     string
				filename string
			}{
				name:     "Check " + file.Name(),
				filename: file.Name(),
			})
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := filepath.Join(templatesDir, tt.filename)
			contentBytes, err := os.ReadFile(filePath)
			if err != nil {
				t.Fatalf("Failed to read file %s: %v", tt.filename, err)
			}

			content := strings.ToLower(string(contentBytes))

			for _, color := range forbiddenColors {
				if strings.Contains(content, color) {
					t.Errorf("File %s contains forbidden light mode color %s", tt.filename, color)
				}
			}
		})
	}
}
