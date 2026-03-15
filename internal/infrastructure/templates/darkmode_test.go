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

	// Forbidden light mode colors that should have been replaced
	forbiddenColors := []string{
		"#ffffff", // White background
		"#f8fafc", // Light gray background
		"#e2e8f0", // Light border
		// Note: #1e293b was dark text in light mode, but it's used as card background in dark mode, so we can't forbid it.
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
