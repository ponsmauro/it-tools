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

	tmpl := template.Must(template.New("root").Parse(`{{ define "commander.html" }}OK{{ end }}`))

	h := &Handler{
		toolUC:       uc,
		templates:    tmpl,
		templateHelp: templates.NewTemplateHelper(),
	}

	tests := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{
			name:           "existing tool template returns 200",
			url:            "/tools/commander",
			expectedStatus: 200,
		},
		{
			name:           "nonexistent tool template returns 500",
			url:            "/tools/nonexistent",
			expectedStatus: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.url, nil)
			w := httptest.NewRecorder()

			h.ToolHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestHandler_HomeHandler(t *testing.T) {
	r := repositories.NewToolRepository()
	uc := usecases.NewToolUseCase(r)
	h := &Handler{
		toolUC:       uc,
		templates:    template.Must(template.New("index.html").Parse("Home")),
		templateHelp: templates.NewTemplateHelper(),
	}
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h.HomeHandler(w, req)
	if w.Code != 200 {
		t.Errorf("Home expected 200, got %d", w.Code)
	}
}
