package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"html/template"

	"it-tools/internal/application/usecases"
	"it-tools/internal/infrastructure/repositories"
	"it-tools/internal/infrastructure/templates"
)

func newTestHandler(tmpl *template.Template) *Handler {
	r := repositories.NewToolRepository()
	uc := usecases.NewToolUseCase(r)
	return &Handler{
		toolUC:       uc,
		templates:    tmpl,
		templateHelp: templates.NewTemplateHelper(),
	}
}

func TestHandler_ToolHandler(t *testing.T) {
	tmpl := template.Must(template.New("root").Parse(`{{ define "commander.html" }}OK{{ end }}`))
	h := newTestHandler(tmpl)

	tests := []struct {
		name           string
		method         string
		url            string
		body           string
		expectedStatus int
	}{
		{
			name:           "GET existing tool returns 200",
			method:         http.MethodGet,
			url:            "/tools/commander",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "GET nonexistent tool returns 404",
			method:         http.MethodGet,
			url:            "/tools/nonexistent",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "GET empty tool ID returns 404",
			method:         http.MethodGet,
			url:            "/tools/",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "PUT returns 405",
			method:         http.MethodPut,
			url:            "/tools/commander",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "DELETE returns 405",
			method:         http.MethodDelete,
			url:            "/tools/commander",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "POST with invalid JSON returns 400",
			method:         http.MethodPost,
			url:            "/tools/commander",
			body:           "not-json",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "POST with wrong action returns 400",
			method:         http.MethodPost,
			url:            "/tools/commander",
			body:           `{"action":"close"}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "POST with action=open and existing template returns 200",
			method:         http.MethodPost,
			url:            "/tools/commander",
			body:           `{"action":"open"}`,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.body != "" {
				req = httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			} else {
				req = httptest.NewRequest(tt.method, tt.url, nil)
			}
			w := httptest.NewRecorder()
			h.ToolHandler(w, req)
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestHandler_HomeHandler(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{"GET home returns 200", http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newTestHandler(template.Must(template.New("index.html").Parse("Home")))
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			h.HomeHandler(w, req)
			if w.Code != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestHandler_AboutHandler(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{"GET about returns 200", http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newTestHandler(template.Must(template.New("about.html").Parse("About")))
			req := httptest.NewRequest(http.MethodGet, "/about", nil)
			w := httptest.NewRecorder()
			h.AboutHandler(w, req)
			if w.Code != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestHandler_GetLanguage(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{"no lang param returns default", "", "en"},
		{"valid lang en", "?lang=en", "en"},
		{"valid lang es", "?lang=es", "es"},
		{"invalid lang falls back to default", "?lang=fr", "en"},
		{"injection attempt falls back to default", "?lang=<script>", "en"},
	}
	h := newTestHandler(template.Must(template.New("t").Parse("")))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/"+tt.query, nil)
			got := h.getLanguage(req)
			if got != tt.expected {
				t.Errorf("expected lang %q, got %q", tt.expected, got)
			}
		})
	}
}
