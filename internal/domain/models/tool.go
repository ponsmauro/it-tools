package models

import (
	"errors"
	"time"
)

// ErrNotFound is returned when a tool is not found by ID
var ErrNotFound = errors.New("tool not found")

// Tool represents a developer tool in the application
type Tool struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Category    string    `json:"category"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
}

// ToolRepository defines the interface for tool data access
type ToolRepository interface {
	GetAll() []Tool
	GetByID(id string) (*Tool, error)
}
