package repositories

import (
	"sort"
	"time"

	"it-tools/internal/domain/models"
)

// ToolRepository implements models.ToolRepository
type ToolRepository struct{}

// NewToolRepository creates a new ToolRepository instance
func NewToolRepository() *ToolRepository {
	return &ToolRepository{}
}

// GetAll returns all available tools
func (r *ToolRepository) GetAll() []models.Tool {
	tools := []models.Tool{
		{
			ID:          "commander",
			Name:        "File Commander",
			Description: "Dual-pane file manager like Total Commander (frontend)",
			Icon:        "📁",
			Category:    "File Management",
			URL:         "/tools/commander",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "terminal",
			Name:        "Web Terminal",
			Description: "Emulated PC terminal (frontend mock shell)",
			Icon:        "💻",
			Category:    "Development",
			URL:         "/tools/terminal",
			CreatedAt:   time.Now(),
		},

		{
			ID:          "base64",
			Name:        "Base64 Encoder/Decoder",
			Description: "Encode and decode text in Base64 format",
			Icon:        "🔐",
			Category:    "Development",
			URL:         "/tools/base64",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "uuid-generator",
			Name:        "UUID Generator",
			Description: "Generate universally unique identifiers (UUID v4)",
			Icon:        "🎲",
			Category:    "Development",
			URL:         "/tools/uuid-generator",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "hash-generator",
			Name:        "Hash Generator",
			Description: "Generate MD5, SHA-1, SHA-256 and SHA-512 hashes",
			Icon:        "#️⃣",
			Category:    "Security",
			URL:         "/tools/hash-generator",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "url-encoder",
			Name:        "URL Encoder/Decoder",
			Description: "Encode and decode special URL characters",
			Icon:        "🔗",
			Category:    "Development",
			URL:         "/tools/url-encoder",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "cron-parser",
			Name:        "Cron Expression Parser",
			Description: "Validate and explain cron expressions",
			Icon:        "⏰",
			Category:    "Development",
			URL:         "/tools/cron-parser",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "jwt-decoder",
			Name:        "JWT Decoder",
			Description: "Decode and validate JWT tokens",
			Icon:        "🎫",
			Category:    "Security",
			URL:         "/tools/jwt-decoder",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "converters",
			Name:        "Converters",
			Description: "Base64, URL, Color, Timestamp, Case converters",
			Icon:        "🔄",
			Category:    "Development",
			URL:         "/tools/converters",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "formatters",
			Name:        "Formatters",
			Description: "JSON, YAML, XML, SQL, GraphQL formatter & validator",
			Icon:        "✨",
			Category:    "Development",
			URL:         "/tools/formatters",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "regex-tester",
			Name:        "Regex Tester",
			Description: "Test and validate regular expressions",
			Icon:        "🔍",
			Category:    "Development",
			URL:         "/tools/regex-tester",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "timestamp-converter",
			Name:        "Timestamp Converter",
			Description: "Convert between Unix timestamp, ISO date, and other formats",
			Icon:        "🕐",
			Category:    "Development",
			URL:         "/tools/timestamp-converter",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "case-converter",
			Name:        "Case Converter",
			Description: "Convert text between camelCase, snake_case, kebab-case, PascalCase",
			Icon:        "🔤",
			Category:    "Development",
			URL:         "/tools/case-converter",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "password-generator",
			Name:        "Password Generator",
			Description: "Generate secure random passwords with custom options",
			Icon:        "🔑",
			Category:    "Security",
			URL:         "/tools/password-generator",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "text-utils",
			Name:        "Text Utilities",
			Description: "Case converter, HTML escape, env formatter",
			Icon:        "📝",
			Category:    "Development",
			URL:         "/tools/text-utils",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "html-escape",
			Name:        "HTML Escape/Unescape",
			Description: "Encode and decode HTML entities",
			Icon:        "🏷️",
			Category:    "Development",
			URL:         "/tools/html-escape",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "qr-generator",
			Name:        "QR Code Generator",
			Description: "Generate QR codes as SVG or PNG",
			Icon:        "📱",
			Category:    "Utilities",
			URL:         "/tools/qr-generator",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "env-formatter",
			Name:        "Environment Formatter",
			Description: "Format environment variables and .env files",
			Icon:        "⚙️",
			Category:    "Development",
			URL:         "/tools/env-formatter",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "diff-checker",
			Name:        "Diff Checker",
			Description: "Compare two texts and show differences",
			Icon:        "➕➖",
			Category:    "Development",
			URL:         "/tools/diff-checker",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "lorem-generator",
			Name:        "Lorem Ipsum Generator",
			Description: "Generate placeholder text for design and testing",
			Icon:        "📄",
			Category:    "Utilities",
			URL:         "/tools/lorem-generator",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "kill-port",
			Name:        "Kill Port",
			Description: "Mata procesos ocupando puerto (lsof/kill)",
			Icon:        "⚡",
			Category:    "DevOps",
			URL:         "/tools/kill-port",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "ip-calculator",
			Name:        "IP Subnet Calculator",
			Description: "Calculate subnets, CIDR, IP ranges",
			Icon:        "🌐",
			Category:    "Network",
			URL:         "/tools/ip-calculator",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "gzip-tool",
			Name:        "Gzip Compressor",
			Description: "Compress/decompress GZIP data",
			Icon:        "📦",
			Category:    "Utilities",
			URL:         "/tools/gzip-tool",
			CreatedAt:   time.Now(),
		},
		{
			ID:          "csv-parser",
			Name:        "CSV Parser & Editor",
			Description: "Parse, sort, filter CSV data",
			Icon:        "📊",
			Category:    "Data",
			URL:         "/tools/csv-parser",
			CreatedAt:   time.Now(),
		},
	}

	sort.Slice(tools, func(i, j int) bool {
		return tools[i].Name < tools[j].Name
	})

	return tools
}

// GetByID returns a tool by its ID
func (r *ToolRepository) GetByID(id string) (*models.Tool, error) {
	tools := r.GetAll()
	for _, tool := range tools {
		if tool.ID == id {
			return &tool, nil
		}
	}
	return nil, nil
}
