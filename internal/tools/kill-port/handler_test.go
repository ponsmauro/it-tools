package killport

import (
	"net/http"
	"net/http/httptest"
	"testing"

	toolregistry "it-tools/internal/tool_registry"
)

func TestServeHTTP_ReturnsNotFound(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{"GET returns 404", http.MethodGet, http.StatusNotFound},
		{"POST returns 404", http.MethodPost, http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/tools/kill-port", nil)
			w := httptest.NewRecorder()
			ServeHTTP(w, req)
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestInit_RegistersHandler(t *testing.T) {
	// Reset registry for isolation
	toolregistry.Reset()

	Init()

	h := toolregistry.GetHandler("kill-port")
	if h == nil {
		t.Error("expected kill-port handler to be registered after Init()")
	}
}
