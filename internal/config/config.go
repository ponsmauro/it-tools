package config

// Server configuration constants
const (
	DefaultPort = 8080
)

// Language constants
const (
	DefaultLanguage = "en"
	SupportedLangEN = "en"
	SupportedLangES = "es"
)

// Path constants
const (
	TemplatesDir = "templates"
	StaticDir    = "static"
	StaticCSSDir = "static/css"
)

// Template constants
const (
	LayoutTemplate  = "layout.html"
	IndexTemplate   = "index.html"
	AboutTemplate   = "about.html"
	TemplatePattern = "*.html"
)

// Route constants
const (
	RouteHome   = "/"
	RouteAbout  = "/about"
	RouteStatic = "/static/"
)
