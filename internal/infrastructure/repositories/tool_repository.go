package repositories

import (
	"sort"
	"time"

	"it-tools/internal/domain/models"
)

// ToolCount is the total number of registered tools. Update when adding or removing tools.
const ToolCount = 24

// ToolRepository implements models.ToolRepository.
// Tools are pre-computed once at construction time for O(1) lookups and zero per-request allocations.
type ToolRepository struct {
	tools []models.Tool
	byID  map[string]models.Tool
}

// NewToolRepository creates a new ToolRepository instance and pre-computes the tool list and index.
func NewToolRepository() *ToolRepository {
	tools := buildTools()
	sort.Slice(tools, func(i, j int) bool {
		return tools[i].Name < tools[j].Name
	})
	byID := make(map[string]models.Tool, len(tools))
	for _, t := range tools {
		byID[t.ID] = t
	}
	return &ToolRepository{tools: tools, byID: byID}
}

// GetAll returns a copy of all available tools (pre-sorted by name).
func (r *ToolRepository) GetAll() []models.Tool {
	result := make([]models.Tool, len(r.tools))
	copy(result, r.tools)
	return result
}

// GetByID returns a tool by its ID in O(1), or models.ErrNotFound if not found.
func (r *ToolRepository) GetByID(id string) (*models.Tool, error) {
	if t, ok := r.byID[id]; ok {
		return &t, nil
	}
	return nil, models.ErrNotFound
}

// buildTools returns the static list of all registered tools.
func buildTools() []models.Tool {
	return []models.Tool{
		{
			ID:          "commander",
			Name:        "File Commander",
			Description: "Dual-pane file manager like Total Commander (frontend)",
			Icon:        "📁",
			Category:    "File Management",
			URL:         "/tools/commander",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "terminal",
			Name:        "Web Terminal",
			Description: "Emulated PC terminal (frontend mock shell)",
			Icon:        "💻",
			Category:    "Development",
			URL:         "/tools/terminal",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "base64",
			Name:        "Base64 Encoder/Decoder",
			Description: "Encode and decode text in Base64 format",
			Icon:        "🔐",
			Category:    "Development",
			URL:         "/tools/base64",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "uuid-generator",
			Name:        "UUID Generator",
			Description: "Generate universally unique identifiers (UUID v4)",
			Icon:        "🎲",
			Category:    "Development",
			URL:         "/tools/uuid-generator",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "hash-generator",
			Name:        "Hash Generator",
			Description: "Generate MD5, SHA-1, SHA-256 and SHA-512 hashes",
			Icon:        "#️⃣",
			Category:    "Security",
			URL:         "/tools/hash-generator",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "url-encoder",
			Name:        "URL Encoder/Decoder",
			Description: "Encode and decode special URL characters",
			Icon:        "🔗",
			Category:    "Development",
			URL:         "/tools/url-encoder",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "cron-parser",
			Name:        "Cron Expression Parser",
			Description: "Validate and explain cron expressions",
			Icon:        "⏰",
			Category:    "Development",
			URL:         "/tools/cron-parser",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "jwt-decoder",
			Name:        "JWT Decoder",
			Description: "Decode and validate JWT tokens",
			Icon:        "🎫",
			Category:    "Security",
			URL:         "/tools/jwt-decoder",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "converters",
			Name:        "Converters",
			Description: "Base64, URL, Color, Timestamp, Case converters",
			Icon:        "🔄",
			Category:    "Development",
			URL:         "/tools/converters",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "formatters",
			Name:        "Formatters",
			Description: "JSON, YAML, XML, SQL, GraphQL formatter & validator",
			Icon:        "✨",
			Category:    "Development",
			URL:         "/tools/formatters",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "regex-tester",
			Name:        "Regex Tester",
			Description: "Test and validate regular expressions",
			Icon:        "🔍",
			Category:    "Development",
			URL:         "/tools/regex-tester",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "timestamp-converter",
			Name:        "Timestamp Converter",
			Description: "Convert between Unix timestamp, ISO date, and other formats",
			Icon:        "🕐",
			Category:    "Development",
			URL:         "/tools/timestamp-converter",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "case-converter",
			Name:        "Case Converter",
			Description: "Convert text between camelCase, snake_case, kebab-case, PascalCase",
			Icon:        "🔤",
			Category:    "Development",
			URL:         "/tools/case-converter",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "password-generator",
			Name:        "Password Generator",
			Description: "Generate secure random passwords with custom options",
			Icon:        "🔑",
			Category:    "Security",
			URL:         "/tools/password-generator",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "text-utils",
			Name:        "Text Utilities",
			Description: "Case converter, HTML escape, env formatter",
			Icon:        "📝",
			Category:    "Development",
			URL:         "/tools/text-utils",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "html-escape",
			Name:        "HTML Escape/Unescape",
			Description: "Encode and decode HTML entities",
			Icon:        "🏷️",
			Category:    "Development",
			URL:         "/tools/html-escape",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "qr-generator",
			Name:        "QR Code Generator",
			Description: "Generate QR codes as SVG or PNG",
			Icon:        "📱",
			Category:    "Utilities",
			URL:         "/tools/qr-generator",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "env-formatter",
			Name:        "Environment Formatter",
			Description: "Format environment variables and .env files",
			Icon:        "⚙️",
			Category:    "Development",
			URL:         "/tools/env-formatter",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "diff-checker",
			Name:        "Diff Checker",
			Description: "Compare two texts and show differences",
			Icon:        "➕➖",
			Category:    "Development",
			URL:         "/tools/diff-checker",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "lorem-generator",
			Name:        "Lorem Ipsum Generator",
			Description: "Generate placeholder text for design and testing",
			Icon:        "📄",
			Category:    "Utilities",
			URL:         "/tools/lorem-generator",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "kill-port",
			Name:        "Kill Port",
			Description: "Find and kill processes occupying a port (lsof/kill)",
			Icon:        "⚡",
			Category:    "DevOps",
			URL:         "/tools/kill-port",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "ip-calculator",
			Name:        "IP Subnet Calculator",
			Description: "Calculate subnets, CIDR, IP ranges",
			Icon:        "🌐",
			Category:    "Network",
			URL:         "/tools/ip-calculator",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "gzip-tool",
			Name:        "Gzip Compressor",
			Description: "Compress/decompress GZIP data",
			Icon:        "📦",
			Category:    "Utilities",
			URL:         "/tools/gzip-tool",
			CreatedAt:   time.Time{},
		},
		{
			ID:          "csv-parser",
			Name:        "CSV Parser & Editor",
			Description: "Parse, sort, filter CSV data",
			Icon:        "📊",
			Category:    "Data",
			URL:         "/tools/csv-parser",
			CreatedAt:   time.Time{},
		},
	}
}
