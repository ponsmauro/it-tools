package handlers

import (
	"net/http/httptest"
	"testing"

	"html/template"

	"it-tools/internal/application/usecases"
	"it-tools/internal/infrastructure/repositories"
	"it-tools/internal/infrastructure/templates"
)

func TestHandler_ToolHandler(t *testing.T) {
	r := repositories.NewToolRepository()
	uc := usecases.NewToolUseCase(r)
	h := &Handler{
		toolUC:       uc,
		templates:    template.Must(template.New("test").Parse("OK")),
		templateHelp: &templates.TemplateHelper{},
	}

	// Existing tool
	req := httptest.NewRequest("GET", "/tools/commander", nil)
	w := httptest.NewRecorder()
	h.ToolHandler(w, req)
	if w.Code != 200 {
		t.Errorf("Home expected 200, got %d", w.Code)
	}

	// Not found
	req2 := httptest.NewRequest("GET", "/tools/nonexistent", nil)
	w2 := httptest.NewRecorder()
	h.ToolHandler(w2, req2)
	if w2.Code != 404 {
		t.Errorf("Not found expected 404, got %d", w2.Code)
	}
}

func TestHandler_HomeHandler(t *testing.T) {
	r := repositories.NewToolRepository()
	uc := usecases.NewToolUseCase(r)
	h := &Handler{
		toolUC:       uc,
		templates:    template.Must(template.New("index.html").Parse("Home")),
		templateHelp: &templates.TemplateHelper{},
	}
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h.HomeHandler(w, req)
	if w.Code != 200 {
		t.Errorf("Home expected 200, got %d", w.Code)
	}
}
