package usecases

import (
	"it-tools/internal/domain/models"
)

// ToolUseCase handles tool-related business logic
type ToolUseCase struct {
	repository models.ToolRepository
}

// NewToolUseCase creates a new ToolUseCase instance
func NewToolUseCase(repository models.ToolRepository) *ToolUseCase {
	return &ToolUseCase{
		repository: repository,
	}
}

// GetAllTools retrieves all available tools
func (uc *ToolUseCase) GetAllTools() []models.Tool {
	return uc.repository.GetAll()
}

// GetToolByID retrieves a tool by its ID
func (uc *ToolUseCase) GetToolByID(id string) (*models.Tool, error) {
	return uc.repository.GetByID(id)
}
