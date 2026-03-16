package killport

import (
	"net/http"

	toolregistry "it-tools/internal/tool_registry"
)

// ServeHTTP is a placeholder handler — kill-port content is served via the template engine
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// Init registers the kill-port tool in the tool registry
func Init() {
	toolregistry.Register("kill-port", ServeHTTP)
}
