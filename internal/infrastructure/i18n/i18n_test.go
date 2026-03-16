package i18n

import (
	"testing"
)

func TestTranslator_Get(t *testing.T) {
	tr := NewTranslator()
	tr.AddTranslation("en", "hello", "Hello")
	tr.AddTranslation("es", "hello", "Hola")

	tests := []struct {
		name     string
		lang     string
		key      string
		expected string
	}{
		{"existing key en", "en", "hello", "Hello"},
		{"existing key es", "es", "hello", "Hola"},
		{"missing key returns key", "en", "missing", "missing"},
		{"missing lang returns key", "fr", "hello", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tr.Get(tt.lang, tt.key)
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestInitDefaultTranslations(t *testing.T) {
	tr := NewTranslator()
	tr.InitDefaultTranslations()

	tests := []struct {
		name     string
		lang     string
		key      string
		expected string
	}{
		{"app title en", "en", "app.title", "IT Tools"},
		{"app title es", "es", "app.title", "IT Tools"},
		{"tools all en", "en", "tools.all", "All Tools"},
		{"tools all es", "es", "tools.all", "Todas las Herramientas"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tr.Get(tt.lang, tt.key)
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
