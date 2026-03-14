package i18n

import (
	"sync"

	"it-tools/internal/config"
)

// Translation keys
const (
	KeyAppTitle                = "app.title"
	KeyAppWelcome              = "app.welcome"
	KeyAppSubtitle             = "app.subtitle"
	KeyNavTools                = "nav.tools"
	KeyNavItTools              = "nav.it-tools"
	KeyNavAbout                = "nav.about"
	KeyToolsAll                = "tools.all"
	KeyAboutTitle              = "about.title"
	KeyAboutMission            = "about.mission"
	KeyAboutFeatures           = "about.features"
	KeyAboutTech               = "about.tech"
	KeyFooterCopy              = "footer.copyright"
	KeyToolsKillPort           = "tools.kill-port"
	KeyToolsKillPortDesc       = "tools.kill-port.desc"
	KeyToolsKillPortPort       = "tools.kill-port.port"
	KeyToolsKillPortScan       = "tools.kill-port.scan"
	KeyToolsKillPortGuideTitle = "tools.kill-port.guide.title"
	KeyToolsKillPortOSMac      = "tools.kill-port.os.macos"
	KeyToolsKillPortOSLinux    = "tools.kill-port.os.linux"
	KeyToolsKillPortOSWindows  = "tools.kill-port.os.windows"
	KeyToolsKillPortStepFind   = "tools.kill-port.step.find"
	KeyToolsKillPortStepKill   = "tools.kill-port.step.kill"
	KeyToolsKillPortStepVerify = "tools.kill-port.step.verify"
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
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPort, "Kill Port")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortDesc, "Find and kill processes occupying specific ports (macOS lsof + kill mock)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortPort, "Port")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortScan, "Scan & Kill")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortGuideTitle, "How to kill a port from terminal")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortOSMac, "macOS")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortOSLinux, "Linux")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortOSWindows, "Windows")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortStepFind, "Find the process using the port")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortStepKill, "Kill the process with PID")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortStepVerify, "Verify the port is free")

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
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPort, "Matar Puerto")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortDesc, "Encuentra y mata procesos ocupando puertos específicos (mock macOS lsof + kill)")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortPort, "Puerto")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortScan, "Escanear y Matar")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortGuideTitle, "Cómo matar un puerto desde terminal")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortOSMac, "macOS")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortOSLinux, "Linux")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortOSWindows, "Windows")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortStepFind, "Encontrar el proceso que usa el puerto")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortStepKill, "Matar el proceso por PID")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortStepVerify, "Verificar que el puerto quedó libre")
}
