package usecases

import (
	"testing"

	"it-tools/internal/infrastructure/repositories"
)

func TestToolUseCase_GetAllTools(t *testing.T) {
	r := repositories.NewToolRepository()
	uc := NewToolUseCase(r)
	got := uc.GetAllTools()
	if len(got) == 0 {
		t.Fatal("no tools")
	}
	if len(got) != 24 {
		t.Errorf("expected 24 tools, got %d", len(got))
	}
}

func TestToolUseCase_GetToolByID(t *testing.T) {
	r := repositories.NewToolRepository()
	uc := NewToolUseCase(r)
	got, err := uc.GetToolByID("base64")
	if err != nil {
		t.Error(err)
	}
	if got == nil || got.ID != "base64" {
		t.Error("expected base64 tool")
	}
}
