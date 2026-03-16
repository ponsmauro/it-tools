package toolregistry

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister_AndGetHandler(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		expectFound bool
	}{
		{"registered handler is found", "test-tool", true},
		{"unregistered handler returns nil", "unknown-tool", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reset()

			if tt.expectFound {
				Register(tt.id, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				})
			}

			got := GetHandler(tt.id)
			if tt.expectFound && got == nil {
				t.Errorf("expected handler for id=%s, got nil", tt.id)
			}
			if !tt.expectFound && got != nil {
				t.Errorf("expected nil for id=%s, got a handler", tt.id)
			}
		})
	}
}

func TestRegister_HandlerIsCallable(t *testing.T) {
	Reset()

	Register("callable-tool", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})

	h := GetHandler("callable-tool")
	if h == nil {
		t.Fatal("expected handler, got nil")
	}

	req := httptest.NewRequest(http.MethodGet, "/tools/callable-tool", nil)
	w := httptest.NewRecorder()
	h(w, req)

	if w.Code != http.StatusAccepted {
		t.Errorf("expected status %d, got %d", http.StatusAccepted, w.Code)
	}
}

func TestRegister_OverwritesExistingHandler(t *testing.T) {
	Reset()

	Register("overwrite-tool", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	Register("overwrite-tool", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	})

	h := GetHandler("overwrite-tool")
	if h == nil {
		t.Fatal("expected handler, got nil")
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	h(w, req)

	if w.Code != http.StatusTeapot {
		t.Errorf("expected overwritten handler status %d, got %d", http.StatusTeapot, w.Code)
	}
}
