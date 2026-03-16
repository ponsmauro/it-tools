package templates

import (
	"testing"
)

func TestTemplateHelper_Translate(t *testing.T) {
	th := NewTemplateHelper()

	tests := []struct {
		name     string
		key      string
		expected string
	}{
		{"translate en", "app.title", "IT Tools"},
		{"translate missing", "missing.key", "missing.key"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := th.Translate(tt.key)
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestTemplateHelper_TranslateTo(t *testing.T) {
	th := NewTemplateHelper()
	th.translator.InitDefaultTranslations()

	tests := []struct {
		name     string
		lang     string
		key      string
		expected string
	}{
		{"translate to en", "en", "app.title", "IT Tools"},
		{"translate to es", "es", "app.title", "IT Tools"},
		{"translate to missing lang", "fr", "app.title", "app.title"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := th.TranslateTo(tt.lang, tt.key)
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestTemplateHelper_GetFuncMap(t *testing.T) {
	th := NewTemplateHelper()
	fm := th.GetFuncMap()

	if fm == nil {
		t.Fatal("expected FuncMap, got nil")
	}

	if _, ok := fm["tr"]; !ok {
		t.Error("expected 'tr' function in FuncMap")
	}
}

func TestTemplateHelper_ParseGlob(t *testing.T) {
	th := NewTemplateHelper()
	
	// Test with a non-existent pattern
	_, err := th.ParseGlob("nonexistent/*.html")
	if err == nil {
		t.Error("expected error parsing nonexistent glob, got nil")
	}
}
