package main

import (
	"fmt"
	"log"
	"net/http"

	"it-tools/internal/application/usecases"
	"it-tools/internal/config"
	"it-tools/internal/infrastructure/handlers"
	"it-tools/internal/infrastructure/repositories"
	"it-tools/internal/infrastructure/templates"
	killport "it-tools/internal/tools/kill-port"
)

func main() {
	// Initialize repository
	toolRepo := repositories.NewToolRepository()

	// Initialize use case
	toolUC := usecases.NewToolUseCase(toolRepo)

	// Initialize template helper
	templateHelp := templates.NewTemplateHelper()

	// Load templates
	templatePath := "./static/templates/*.html"
	tmpl, err := templateHelp.ParseGlob(templatePath)
	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}
	log.Println("Templates loaded successfully")

	// Initialize handler
	handler := handlers.NewHandler(tmpl, toolUC, templateHelp)

	// Static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Init kill-port registry
	killport.Init()

	// Routes
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/tools/", handler.ToolHandler)

	// Initialize kill-port tool
	killport.Init()

	// Start server
	port := config.DefaultPort
	log.Printf("Server started at http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
