package templates

import (
	"html/template"
	"sync"

	"it-tools/internal/infrastructure/i18n"
)

// TemplateHelper provides template functions
type TemplateHelper struct {
	translator *i18n.Translator
	funcMap    template.FuncMap
	mu         sync.RWMutex
}

// NewTemplateHelper creates a new TemplateHelper
func NewTemplateHelper() *TemplateHelper {
	th := &TemplateHelper{
		translator: i18n.NewTranslator(),
	}
	th.translator.InitDefaultTranslations()
	th.funcMap = template.FuncMap{
		"tr": th.Translate,
	}
	return th
}

// Translate returns the translation for a key
func (th *TemplateHelper) Translate(key string) string {
	return th.translator.Get("en", key)
}

// TranslateTo returns the translation for a key in a specific language
func (th *TemplateHelper) TranslateTo(lang string, key string) string {
	return th.translator.Get(lang, key)
}

// GetFuncMap returns the template function map
func (th *TemplateHelper) GetFuncMap() template.FuncMap {
	return th.funcMap
}

// ParseGlob parses templates with the helper functions
func (th *TemplateHelper) ParseGlob(pattern string) (*template.Template, error) {
	tmpl := template.New("").Funcs(th.funcMap)
	return tmpl.ParseGlob(pattern)
}
