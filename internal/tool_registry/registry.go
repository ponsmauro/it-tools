package toolregistry

import (
	"net/http"
	"sync"
)

// ToolHandlerFunc is the handler type for tools
type ToolHandlerFunc func(http.ResponseWriter, *http.Request)

var (
	mu       sync.RWMutex
	handlers = make(map[string]ToolHandlerFunc)
)

// Register adds a tool handler by ID
func Register(id string, handler ToolHandlerFunc) {
	mu.Lock()
	defer mu.Unlock()
	handlers[id] = handler
}

// GetHandler returns the handler for the given tool ID, or nil if not found
func GetHandler(id string) ToolHandlerFunc {
	mu.RLock()
	defer mu.RUnlock()
	return handlers[id]
}

// Reset clears all registered handlers. Intended for use in tests only.
func Reset() {
	mu.Lock()
	defer mu.Unlock()
	handlers = make(map[string]ToolHandlerFunc)
}
