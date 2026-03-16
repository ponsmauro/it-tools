package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"it-tools/internal/application/usecases"
	"it-tools/internal/config"
	"it-tools/internal/infrastructure/handlers"
	"it-tools/internal/infrastructure/repositories"
	"it-tools/internal/infrastructure/templates"
	killport "it-tools/internal/tools/kill-port"
)

// securityHeaders adds mandatory HTTP security headers to every response
func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Get working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}

	// Initialize repository
	toolRepo := repositories.NewToolRepository()

	// Initialize use case
	toolUC := usecases.NewToolUseCase(toolRepo)

	// Initialize template helper
	templateHelp := templates.NewTemplateHelper()

	// Load templates
	templatePath := filepath.Join(wd, config.StaticDir, "templates", config.TemplatePattern)
	tmpl, err := templateHelp.ParseGlob(templatePath)
	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}

	// Initialize handler
	handler := handlers.NewHandler(tmpl, toolUC, templateHelp)

	// Static files
	staticPath := filepath.Join(wd, config.StaticDir)
	fs := http.FileServer(http.Dir(staticPath))
	http.Handle(config.RouteStatic, http.StripPrefix("/static/", fs))

	// Register tools
	killport.Init()

	// Routes
	http.HandleFunc(config.RouteHome, handler.HomeHandler)
	http.HandleFunc(config.RouteAbout, handler.AboutHandler)
	http.HandleFunc("/tools/", handler.ToolHandler)

	// Start server with security headers middleware
	port := config.DefaultPort
	log.Printf("Server started at http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), securityHeaders(http.DefaultServeMux)))

}
