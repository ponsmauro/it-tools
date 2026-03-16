package repositories

import (
	"errors"
	"testing"

	"it-tools/internal/domain/models"
)

func TestToolRepository_GetAll(t *testing.T) {
	r := NewToolRepository()
	got := r.GetAll()
	if len(got) == 0 {
		t.Fatal("no tools returned")
	}
	if len(got) != ToolCount {
		t.Errorf("expected %d tools, got %d", ToolCount, len(got))
	}
}

func TestToolRepository_GetByID(t *testing.T) {
	r := NewToolRepository()
	tests := []struct {
		name        string
		id          string
		expectFound bool
	}{
		{"existing tool commander", "commander", true},
		{"existing tool base64", "base64", true},
		{"nonexistent tool", "nonexistent", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetByID(tt.id)
			if tt.expectFound {
				if err != nil {
					t.Errorf("unexpected error for id=%s: %v", tt.id, err)
				}
				if got == nil || got.ID != tt.id {
					t.Errorf("expected tool with id=%s", tt.id)
				}
			} else {
				if !errors.Is(err, models.ErrNotFound) {
					t.Errorf("expected ErrNotFound for id=%s, got %v", tt.id, err)
				}
				if got != nil {
					t.Errorf("expected nil tool for id=%s", tt.id)
				}
			}
		})
	}
}
