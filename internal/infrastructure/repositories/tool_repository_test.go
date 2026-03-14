package repositories

import (
	"testing"
)

func TestToolRepository_GetAll(t *testing.T) {
	r := NewToolRepository()
	got := r.GetAll()
	if len(got) == 0 {
		t.Fatal("no tools returned")
	}
	if len(got) != 24 {
		t.Errorf("expected 24 tools, got %d", len(got))
	}
}

func TestToolRepository_GetByID(t *testing.T) {
	r := NewToolRepository()
	tests := []struct {
		name  string
		id    string
		found bool
	}{
		{"commander", "commander", true},
		{"base64", "base64", true},
		{"nonexistent", "nonexistent", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetByID(tt.id)
			if err != nil {
				t.Error(err)
			}
			if tt.found {
				if got == nil || got.ID != tt.id {
					t.Errorf("expected tool %s", tt.id)
				}
			} else {
				if got != nil {
					t.Error("expected nil for nonexistent")
				}
			}
		})
	}
}
