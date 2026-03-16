package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"it-tools/internal/config"
	"it-tools/internal/domain/models"
	"it-tools/internal/infrastructure/templates"
)

const maxRequestBodyBytes = 1024

// ToolUseCasePort defines the application operations required by the handler layer.
// Using an interface decouples the handler from the concrete use case implementation.
type ToolUseCasePort interface {
	GetAllTools() []models.Tool
	GetToolByID(id string) (*models.Tool, error)
}

// Handler contains all HTTP handlers
type Handler struct {
	templates    *template.Template
	toolUC       ToolUseCasePort
	templateHelp *templates.TemplateHelper
}

// NewHandler creates a new Handler instance
func NewHandler(templates *template.Template, toolUC ToolUseCasePort, templateHelp *templates.TemplateHelper) *Handler {
	return &Handler{
		templates:    templates,
		toolUC:       toolUC,
		templateHelp: templateHelp,
	}
}

// HomeHandler serves the home page with all tools
func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	h.servePage(w, r, config.IndexTemplate, "IT Tools", config.RouteHome)
}

// AboutHandler serves the about page
func (h *Handler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	h.servePage(w, r, config.AboutTemplate, "About - IT Tools", config.RouteAbout)
}

// ToolHandler serves individual tools via GET (full page) or POST (AJAX tab content)
func (h *Handler) ToolHandler(w http.ResponseWriter, r *http.Request) {
	toolID := r.URL.Path[len("/tools/"):]
	if toolID == "" {
		http.NotFound(w, r)
		return
	}

	// Validate toolID against the known set of tools
	if _, err := h.toolUC.GetToolByID(toolID); errors.Is(err, models.ErrNotFound) {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		templateName := toolID + ".html"
		h.servePage(w, r, templateName, toolID+" - IT Tools", "/tools/"+toolID)
	case http.MethodPost:
		h.handleToolPost(w, r, toolID)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// handleToolPost handles POST /tools/{id} dynamic content
func (h *Handler) handleToolPost(w http.ResponseWriter, r *http.Request, toolID string) {
	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)

	var req struct {
		Action string `json:"action"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("JSON decode error: tool=%s error=%v", toolID, err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Action != "open" {
		http.Error(w, "Action must be 'open'", http.StatusBadRequest)
		return
	}

	templateName := toolID + ".html"
	lang := h.getLanguage(r)
	data := h.buildPageData(toolID+" - IT Tools", lang, "/tools/"+toolID)

	var buf bytes.Buffer
	err := h.templates.ExecuteTemplate(&buf, templateName, data)
	if err != nil {
		log.Printf("Error executing tool template: tool=%s, template=%s, error=%v", toolID, templateName, err)
		http.Error(w, fmt.Sprintf("Template not found for tool: %s", toolID), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Printf("Error writing tool response: tool=%s, error=%v", toolID, err)
	}
}

// servePage is a reusable method to serve any page
func (h *Handler) servePage(w http.ResponseWriter, r *http.Request, page string, title string, currentRoute string) {
	lang := h.getLanguage(r)
	data := h.buildPageData(title, lang, currentRoute)
	h.executeTemplate(w, page, data)
}

// buildPageData is a reusable method to build common page data
func (h *Handler) buildPageData(title string, lang string, currentRoute string) map[string]interface{} {
	tr := func(key string) string { return h.templateHelp.TranslateTo(lang, key) }
	return map[string]interface{}{
		"Title":        title,
		"Tools":        h.toolUC.GetAllTools(),
		"Lang":         lang,
		"CurrentRoute": currentRoute,
		"tr":           tr,
	}
}

// getLanguage extracts and validates the language from the request
func (h *Handler) getLanguage(r *http.Request) string {
	lang := r.URL.Query().Get("lang")
	if lang == config.SupportedLangEN || lang == config.SupportedLangES {
		return lang
	}
	return config.DefaultLanguage
}

// executeTemplate executes a template with error handling
func (h *Handler) executeTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	var buf bytes.Buffer
	err := h.templates.ExecuteTemplate(&buf, tmpl, data)
	if err != nil {
		log.Printf("Error executing template: page=%s, error=%v", tmpl, err)
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Printf("Error writing response: page=%s, error=%v", tmpl, err)
	}
}
