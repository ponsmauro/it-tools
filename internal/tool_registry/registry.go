package toolregistry // note: camelcase for module

import "net/http"

// ToolHandlerFunc is the handler type for tools
type ToolHandlerFunc func(http.ResponseWriter, *http.Request)

// handlers map
var handlers = make(map[string]ToolHandlerFunc)

// Register adds tool handler
func Register(id string, handler ToolHandlerFunc, _ string) { // ignore base64Tool
	handlers[id] = handler
}

// GetHandler returns handler or nil
func GetHandler(id string) ToolHandlerFunc {
	if h, ok := handlers[id]; ok {
		return h
	}
	return nil
}
