package i18n

import (
	"sync"

	"it-tools/internal/config"
)

// Translation keys
const (
	KeyAppTitle      = "app.title"
	KeyAppWelcome    = "app.welcome"
	KeyAppSubtitle   = "app.subtitle"
	KeyNavTools      = "nav.tools"
	KeyNavItTools    = "nav.it-tools"
	KeyNavAbout      = "nav.about"
	KeyToolsAll      = "tools.all"
	KeyAboutTitle    = "about.title"
	KeyAboutMission  = "about.mission"
	KeyAboutFeatures = "about.features"
	KeyAboutTech     = "about.tech"
	KeyFooterCopy    = "footer.copyright"
)

// Translator handles internationalization
type Translator struct {
	translations map[string]map[string]string
	mu           sync.RWMutex
	defaultLang  string
}

// NewTranslator creates a new Translator instance
func NewTranslator() *Translator {
	return &Translator{
		translations: make(map[string]map[string]string),
		defaultLang:  config.DefaultLanguage,
	}
}

// AddTranslation adds a translation for a language
func (t *Translator) AddTranslation(lang string, key string, value string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.translations[lang] == nil {
		t.translations[lang] = make(map[string]string)
	}
	t.translations[lang][key] = value
}

// Get returns the translation for a key
func (t *Translator) Get(lang string, key string) string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if lang == "" {
		lang = t.defaultLang
	}
	if val, ok := t.translations[lang][key]; ok {
		return val
	}
	return key
}

// InitDefaultTranslations initializes default translations
func (t *Translator) InitDefaultTranslations() {
	// English translations
	t.AddTranslation(config.SupportedLangEN, KeyAppTitle, "IT Tools")
	t.AddTranslation(config.SupportedLangEN, KeyAppWelcome, "Welcome to IT Tools")
	t.AddTranslation(config.SupportedLangEN, KeyAppSubtitle, "Collection of useful tools for developers")
	t.AddTranslation(config.SupportedLangEN, KeyNavTools, "Tools")
	t.AddTranslation(config.SupportedLangEN, KeyNavItTools, "IT Tools")
	t.AddTranslation(config.SupportedLangEN, KeyNavAbout, "About")
	t.AddTranslation(config.SupportedLangEN, KeyToolsAll, "All Tools")
	t.AddTranslation(config.SupportedLangEN, KeyAboutTitle, "About IT Tools")
	t.AddTranslation(config.SupportedLangEN, KeyAboutMission, "Our Mission")
	t.AddTranslation(config.SupportedLangEN, KeyAboutFeatures, "Features")
	t.AddTranslation(config.SupportedLangEN, KeyAboutTech, "Technologies")
	t.AddTranslation(config.SupportedLangEN, KeyFooterCopy, "2026 IT Tools - Tools for developers")

	// Spanish translations
	t.AddTranslation(config.SupportedLangES, KeyAppTitle, "IT Tools")
	t.AddTranslation(config.SupportedLangES, KeyAppWelcome, "Bienvenido a IT Tools")
	t.AddTranslation(config.SupportedLangES, KeyAppSubtitle, "Colección de herramientas útiles para desarrolladores")
	t.AddTranslation(config.SupportedLangES, KeyNavTools, "Herramientas")
	t.AddTranslation(config.SupportedLangES, KeyNavItTools, "IT Tools")
	t.AddTranslation(config.SupportedLangES, KeyNavAbout, "Acerca de")
	t.AddTranslation(config.SupportedLangES, KeyToolsAll, "Todas las Herramientas")
	t.AddTranslation(config.SupportedLangES, KeyAboutTitle, "Acerca de IT Tools")
	t.AddTranslation(config.SupportedLangES, KeyAboutMission, "Nuestra Misión")
	t.AddTranslation(config.SupportedLangES, KeyAboutFeatures, "Características")
	t.AddTranslation(config.SupportedLangES, KeyAboutTech, "Tecnologías")
	t.AddTranslation(config.SupportedLangES, KeyFooterCopy, "2026 IT Tools - Herramientas para desarrolladores")
}
