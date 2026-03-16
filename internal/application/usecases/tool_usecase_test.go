package usecases

import (
	"errors"
	"testing"

	"it-tools/internal/domain/models"
	"it-tools/internal/infrastructure/repositories"
)

func TestToolUseCase_GetAllTools(t *testing.T) {
	tests := []struct {
		name          string
		expectedCount int
	}{
		{"returns all tools", repositories.ToolCount},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repositories.NewToolRepository()
			uc := NewToolUseCase(r)
			got := uc.GetAllTools()
			if len(got) == 0 {
				t.Fatal("expected tools, got none")
			}
			if len(got) != tt.expectedCount {
				t.Errorf("expected %d tools, got %d", tt.expectedCount, len(got))
			}
		})
	}
}

func TestToolUseCase_GetToolByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		expectNil bool
	}{
		{"existing tool base64", "base64", false},
		{"existing tool uuid-generator", "uuid-generator", false},
		{"non-existing tool", "does-not-exist", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repositories.NewToolRepository()
			uc := NewToolUseCase(r)
			got, err := uc.GetToolByID(tt.id)
			if tt.expectNil {
				if !errors.Is(err, models.ErrNotFound) {
					t.Errorf("expected ErrNotFound for id=%s, got %v", tt.id, err)
				}
				if got != nil {
					t.Errorf("expected nil tool for id=%s", tt.id)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for id=%s: %v", tt.id, err)
				}
				if got == nil || got.ID != tt.id {
					t.Errorf("expected tool with id=%s, got %v", tt.id, got)
				}
			}
		})
	}
}
