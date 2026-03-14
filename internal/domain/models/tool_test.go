package models

import "testing"

func TestTool_ID(t *testing.T) {
	tests := []struct {
		name string
		id   string
	}{
		{"valid", "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tool := Tool{ID: tt.id}
			if tool.ID != tt.id {
				t.Errorf("expected ID %s, got %s", tt.id, tool.ID)
			}
		})
	}
}
