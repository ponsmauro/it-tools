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

	// Common Tool Keys
	KeyToolsInput          = "tools.input"
	KeyToolsOutput         = "tools.output"
	KeyToolsCopyOutput     = "tools.copy-output"
	KeyToolsGenerate       = "tools.generate"
	KeyToolsCopy           = "tools.copy"
	KeyToolsCopyAll        = "tools.copy-all"
	KeyToolsClear          = "tools.clear"
	KeyToolsChars          = "tools.chars"
	KeyToolsWords          = "tools.words"
	KeyToolsLines          = "tools.lines"
	KeyToolsEmptyLines     = "tools.empty-lines"
	KeyToolsReplace        = "tools.replace"
	KeyToolsStatsGenerated = "tools.stats.generated"

	// Kill Port Tool
	KeyToolsKillPort            = "tools.kill-port"
	KeyToolsKillPortDesc        = "tools.kill-port.desc"
	KeyToolsKillPortPort        = "tools.kill-port.port"
	KeyToolsKillPortScan        = "tools.kill-port.scan"
	KeyToolsKillPortGuideTitle  = "tools.kill-port.guide.title"
	KeyToolsKillPortOSMac       = "tools.kill-port.os.macos"
	KeyToolsKillPortOSLinux     = "tools.kill-port.os.linux"
	KeyToolsKillPortOSWindows   = "tools.kill-port.os.windows"
	KeyToolsKillPortStepFind    = "tools.kill-port.step.find"
	KeyToolsKillPortStepKill    = "tools.kill-port.step.kill"
	KeyToolsKillPortStepVerify  = "tools.kill-port.step.verify"
	KeyToolsKillPortInvalidPort = "tools.kill-port.invalid-port"
	KeyToolsKillPortScanning    = "tools.kill-port.scanning"
	KeyToolsKillPortProcesses   = "tools.kill-port.processes"
	KeyToolsKillPortKilled      = "tools.kill-port.killed"
	KeyToolsKillPortFreed       = "tools.kill-port.freed"

	// Commander Tool
	KeyToolsCommander            = "tools.commander"
	KeyToolsCommanderTitle       = "tools.commander.title"
	KeyToolsCommanderDesc        = "tools.commander.desc"
	KeyToolsCommanderCopy        = "tools.commander.copy"
	KeyToolsCommanderMove        = "tools.commander.move"
	KeyToolsCommanderDelete      = "tools.commander.delete"
	KeyToolsCommanderCopyAlert   = "tools.commander.copy-alert"
	KeyToolsCommanderMoveAlert   = "tools.commander.move-alert"
	KeyToolsCommanderDeleteAlert = "tools.commander.delete-alert"

	// Terminal Tool
	KeyToolsTerminal                = "tools.terminal"
	KeyToolsTerminalTitle           = "tools.terminal.title"
	KeyToolsTerminalDesc            = "tools.terminal.desc"
	KeyToolsTerminalWelcome         = "tools.terminal.welcome"
	KeyToolsTerminalPlaceholder     = "tools.terminal.placeholder"
	KeyToolsTerminalCleared         = "tools.terminal.cleared"
	KeyToolsTerminalCommandNotFound = "tools.terminal.command-not-found"

	// About Page
	KeyAboutMissionDesc        = "about.mission.desc"
	KeyAboutFeaturesFree       = "about.features.free"
	KeyAboutFeaturesNoReg      = "about.features.no-registration"
	KeyAboutFeaturesPrivacy    = "about.features.privacy"
	KeyAboutFeaturesOpenSource = "about.features.open-source"
	KeyAboutTechDesc           = "about.tech.desc"

	// UUID Generator Tool
	KeyToolsUuidGenerator          = "tools.uuid-generator"
	KeyToolsUuidGeneratorTitle     = "tools.uuid-generator.title"
	KeyToolsUuidGeneratorDesc      = "tools.uuid-generator.desc"
	KeyToolsUuidGeneratorCount     = "tools.uuid-generator.count"
	KeyToolsUuidGeneratorUppercase = "tools.uuid-generator.uppercase"
	KeyToolsUuidGeneratorHyphens   = "tools.uuid-generator.hyphens"
	KeyToolsUuidGeneratorBraces    = "tools.uuid-generator.braces"

	// Text Utils Tool
	KeyToolsTextUtils                   = "tools.text-utils"
	KeyToolsTextUtilsTitle              = "tools.text-utils.title"
	KeyToolsTextUtilsDesc               = "tools.text-utils.desc"
	KeyToolsTextUtilsLineOps            = "tools.text-utils.line-operations"
	KeyToolsTextUtilsCaseOps            = "tools.text-utils.case-operations"
	KeyToolsTextUtilsWhitespace         = "tools.text-utils.whitespace"
	KeyToolsTextUtilsFindReplace        = "tools.text-utils.find-replace"
	KeyToolsTextUtilsFindPlaceholder    = "tools.text-utils.find-placeholder"
	KeyToolsTextUtilsReplacePlaceholder = "tools.text-utils.replace-placeholder"
	KeyToolsTextUtilsCaseSensitive      = "tools.text-utils.case-sensitive"

	// Base64 Tool
	KeyToolsBase64                  = "tools.base64"
	KeyToolsBase64Title             = "tools.base64.title"
	KeyToolsBase64Desc              = "tools.base64.desc"
	KeyToolsBase64Encode            = "tools.base64.encode"
	KeyToolsBase64Decode            = "tools.base64.decode"
	KeyToolsBase64UrlSafe           = "tools.base64.url-safe"
	KeyToolsBase64IncludePadding    = "tools.base64.include-padding"
	KeyToolsBase64InputPlaceholder  = "tools.base64.input-placeholder"
	KeyToolsBase64OutputPlaceholder = "tools.base64.output-placeholder"
	KeyToolsBase64FileUpload        = "tools.base64.file-upload"
	KeyToolsBase64EncodeBtn         = "tools.base64.encode-btn"
	KeyToolsBase64DecodeBtn         = "tools.base64.decode-btn"
	KeyToolsBase64Download          = "tools.base64.download"
	KeyToolsBase64AboutTitle        = "tools.base64.about-title"
	KeyToolsBase64AboutDesc         = "tools.base64.about-desc"
	KeyToolsBase64AboutItem1        = "tools.base64.about-item1"
	KeyToolsBase64AboutItem2        = "tools.base64.about-item2"
	KeyToolsBase64AboutItem3        = "tools.base64.about-item3"
	KeyToolsBase64NoInput           = "tools.base64.no-input"
	KeyToolsBase64EncodeSuccess     = "tools.base64.encode-success"
	KeyToolsBase64EncodeError       = "tools.base64.encode-error"
	KeyToolsBase64DecodeSuccess     = "tools.base64.decode-success"
	KeyToolsBase64DecodeError       = "tools.base64.decode-error"
	KeyToolsBase64Cleared           = "tools.base64.cleared"
	KeyToolsBase64NothingToCopy     = "tools.base64.nothing-to-copy"
	KeyToolsBase64CopySuccess       = "tools.base64.copy-success"
	KeyToolsBase64CopyError         = "tools.base64.copy-error"
	KeyToolsBase64NothingToDownload = "tools.base64.nothing-to-download"
	KeyToolsBase64DownloadStarted   = "tools.base64.download-started"
	KeyToolsBase64DownloadError     = "tools.base64.download-error"
	KeyToolsBase64FileTooLarge      = "tools.base64.file-too-large"
	KeyToolsBase64FileEncoded       = "tools.base64.file-encoded"
	KeyToolsBase64FileError         = "tools.base64.file-error"
	KeyToolsBase64FileReadError     = "tools.base64.file-read-error"

	// Password Generator Tool
	KeyToolsPasswordGenerator                       = "tools.password-generator"
	KeyToolsPasswordGeneratorTitle                  = "tools.password-generator.title"
	KeyToolsPasswordGeneratorDesc                   = "tools.password-generator.desc"
	KeyToolsPasswordGeneratorLength                 = "tools.password-generator.length"
	KeyToolsPasswordGeneratorUppercase              = "tools.password-generator.uppercase"
	KeyToolsPasswordGeneratorLowercase              = "tools.password-generator.lowercase"
	KeyToolsPasswordGeneratorNumbers                = "tools.password-generator.numbers"
	KeyToolsPasswordGeneratorSymbols                = "tools.password-generator.symbols"
	KeyToolsPasswordGeneratorExcludeSimilar         = "tools.password-generator.exclude-similar"
	KeyToolsPasswordGeneratorExcludeAmbiguous       = "tools.password-generator.exclude-ambiguous"
	KeyToolsPasswordGeneratorNoRepeats              = "tools.password-generator.no-repeats"
	KeyToolsPasswordGeneratorCustomChars            = "tools.password-generator.custom-chars"
	KeyToolsPasswordGeneratorCustomCharsPlaceholder = "tools.password-generator.custom-chars-placeholder"
	KeyToolsPasswordGeneratorCustomCharsHint        = "tools.password-generator.custom-chars-hint"
	KeyToolsPasswordGeneratorGenerateMultiple       = "tools.password-generator.generate-multiple"
	KeyToolsPasswordGeneratorAboutTitle             = "tools.password-generator.about.title"
	KeyToolsPasswordGeneratorAboutItem1             = "tools.password-generator.about.item1"
	KeyToolsPasswordGeneratorAboutItem2             = "tools.password-generator.about.item2"
	KeyToolsPasswordGeneratorAboutItem3             = "tools.password-generator.about.item3"
	KeyToolsPasswordGeneratorAboutItem4             = "tools.password-generator.about.item4"

	// Text Utils Additional Keys
	KeyToolsReverseLines   = "tools.reverse-lines"
	KeyToolsSortLines      = "tools.sort-lines"
	KeyToolsUniqueLines    = "tools.unique-lines"
	KeyToolsAddLineNumbers = "tools.add-line-numbers"
	KeyToolsUppercase      = "tools.uppercase"
	KeyToolsLowercase      = "tools.lowercase"
	KeyToolsTitleCase      = "tools.title-case"
	KeyToolsSentenceCase   = "tools.sentence-case"
	KeyToolsTrimStart      = "tools.trim-start"
	KeyToolsTrimEnd        = "tools.trim-end"
	KeyToolsTrimBoth       = "tools.trim-both"
	KeyToolsClearAll       = "tools.clear-all"
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
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortInvalidPort, "Invalid port number. Please enter a number between 1 and 65535.")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortScanning, "Scanning port")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortProcesses, "Processes:")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortKilled, "Killed PIDs")
	t.AddTranslation(config.SupportedLangEN, KeyToolsKillPortFreed, "Freed!")

	// Commander Tool
	t.AddTranslation(config.SupportedLangEN, KeyToolsCommander, "File Commander")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCommanderTitle, "🗂️ File Commander")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCommanderDesc, "Dual-pane file manager (frontend demo with virtual FS)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCommanderCopy, "Copy")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCommanderMove, "Move")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCommanderDelete, "Delete")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCommanderCopyAlert, "Copy left → right (mock)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCommanderMoveAlert, "Move left → right (mock)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCommanderDeleteAlert, "Delete selected (mock)")

	// Terminal Tool
	t.AddTranslation(config.SupportedLangEN, KeyToolsTerminal, "Web Terminal")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTerminalTitle, "💻 Web Terminal")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTerminalDesc, "Mock PC terminal emulator (frontend JS shell)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTerminalWelcome, "Welcome to mock shell v1.0")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTerminalPlaceholder, "Type commands: ls, cd dir1, pwd, cat file1.txt, clear...")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTerminalCleared, "Terminal cleared")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTerminalCommandNotFound, "command not found")

	// About Page
	t.AddTranslation(config.SupportedLangEN, KeyAboutMissionDesc, "IT Tools is a collection of online tools designed to facilitate the daily work of developers. From JSON formatting to hash generation, you will find useful tools for your workflow.")
	t.AddTranslation(config.SupportedLangEN, KeyAboutFeaturesFree, "Completely free")
	t.AddTranslation(config.SupportedLangEN, KeyAboutFeaturesNoReg, "No registration required")
	t.AddTranslation(config.SupportedLangEN, KeyAboutFeaturesPrivacy, "Privacy guaranteed")
	t.AddTranslation(config.SupportedLangEN, KeyAboutFeaturesOpenSource, "Open source")
	t.AddTranslation(config.SupportedLangEN, KeyAboutTechDesc, "This project is built with Go (Golang) and uses HTML templates for the frontend.")

	// Common Tool Keys
	t.AddTranslation(config.SupportedLangEN, KeyToolsInput, "Input")
	t.AddTranslation(config.SupportedLangEN, KeyToolsOutput, "Output")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCopyOutput, "Copy Output")
	t.AddTranslation(config.SupportedLangEN, KeyToolsGenerate, "Generate")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCopy, "Copy")
	t.AddTranslation(config.SupportedLangEN, KeyToolsCopyAll, "Copy All")
	t.AddTranslation(config.SupportedLangEN, KeyToolsClear, "Clear")
	t.AddTranslation(config.SupportedLangEN, KeyToolsChars, "Characters")
	t.AddTranslation(config.SupportedLangEN, KeyToolsWords, "Words")
	t.AddTranslation(config.SupportedLangEN, KeyToolsLines, "Lines")
	t.AddTranslation(config.SupportedLangEN, KeyToolsEmptyLines, "Empty Lines")
	t.AddTranslation(config.SupportedLangEN, KeyToolsReplace, "Replace")
	t.AddTranslation(config.SupportedLangEN, KeyToolsStatsGenerated, "Generated")

	// UUID Generator Tool
	t.AddTranslation(config.SupportedLangEN, KeyToolsUuidGenerator, "UUID Generator")
	t.AddTranslation(config.SupportedLangEN, KeyToolsUuidGeneratorTitle, "UUID Generator")
	t.AddTranslation(config.SupportedLangEN, KeyToolsUuidGeneratorDesc, "Generate RFC 4122 compliant UUID v4 tokens using crypto.randomUUID()")
	t.AddTranslation(config.SupportedLangEN, KeyToolsUuidGeneratorCount, "Number of UUIDs")
	t.AddTranslation(config.SupportedLangEN, KeyToolsUuidGeneratorUppercase, "Uppercase")
	t.AddTranslation(config.SupportedLangEN, KeyToolsUuidGeneratorHyphens, "Include hyphens")
	t.AddTranslation(config.SupportedLangEN, KeyToolsUuidGeneratorBraces, "Include braces")

	// Text Utils Tool
	t.AddTranslation(config.SupportedLangEN, KeyToolsTextUtils, "Text Utils")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTextUtilsTitle, "Text Utils")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTextUtilsDesc, "Transform and manipulate text with various operations")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTextUtilsLineOps, "Line Operations")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTextUtilsCaseOps, "Case Operations")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTextUtilsWhitespace, "Whitespace")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTextUtilsFindReplace, "Find & Replace")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTextUtilsFindPlaceholder, "Find text...")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTextUtilsReplacePlaceholder, "Replace with...")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTextUtilsCaseSensitive, "Case sensitive")

	// Text Utils Additional Keys
	t.AddTranslation(config.SupportedLangEN, KeyToolsReverseLines, "Reverse Lines")
	t.AddTranslation(config.SupportedLangEN, KeyToolsSortLines, "Sort Lines")
	t.AddTranslation(config.SupportedLangEN, KeyToolsUniqueLines, "Unique Lines")
	t.AddTranslation(config.SupportedLangEN, KeyToolsAddLineNumbers, "Add Line Numbers")
	t.AddTranslation(config.SupportedLangEN, KeyToolsUppercase, "UPPERCASE")
	t.AddTranslation(config.SupportedLangEN, KeyToolsLowercase, "lowercase")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTitleCase, "Title Case")
	t.AddTranslation(config.SupportedLangEN, KeyToolsSentenceCase, "Sentence case")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTrimStart, "Trim Start")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTrimEnd, "Trim End")
	t.AddTranslation(config.SupportedLangEN, KeyToolsTrimBoth, "Trim Both")
	t.AddTranslation(config.SupportedLangEN, KeyToolsClearAll, "Clear All")

	// Base64 Tool
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64, "Base64")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64Title, "Base64 Encoder/Decoder")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64Desc, "Encode and decode Base64 strings")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64Encode, "Encode")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64Decode, "Decode")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64UrlSafe, "URL safe")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64IncludePadding, "Include padding")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64InputPlaceholder, "Enter text to encode...")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64OutputPlaceholder, "Encoded/decoded result will appear here...")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64FileUpload, "Or upload a file")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64EncodeBtn, "Encode")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64DecodeBtn, "Decode")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64Download, "Download")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64AboutTitle, "About Base64")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64AboutDesc, "Base64 is a group of binary-to-text encoding schemes that represent binary data in ASCII string format.")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64AboutItem1, "Commonly used for encoding data in URLs, emails, and storing complex data in text fields")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64AboutItem2, "Each Base64 digit represents exactly 6 bits of data")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64AboutItem3, "The resulting Base64 string is approximately 33% larger than the original binary data")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64NoInput, "Please enter some text to process")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64EncodeSuccess, "Encoded successfully!")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64EncodeError, "Error encoding: invalid input")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64DecodeSuccess, "Decoded successfully!")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64DecodeError, "Error decoding: invalid Base64 string")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64Cleared, "Cleared")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64NothingToCopy, "Nothing to copy")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64CopySuccess, "Copied to clipboard!")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64CopyError, "Failed to copy to clipboard")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64NothingToDownload, "Nothing to download")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64DownloadStarted, "Download started!")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64DownloadError, "Error downloading file")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64FileTooLarge, "File too large. Maximum size is 5MB.")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64FileEncoded, "File encoded successfully!")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64FileError, "Error processing file")
	t.AddTranslation(config.SupportedLangEN, KeyToolsBase64FileReadError, "Error reading file")

	// Password Generator Tool
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGenerator, "Password Generator")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorTitle, "🔐 Password Generator")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorDesc, "Generate secure random passwords")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorLength, "Length")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorUppercase, "Uppercase (A-Z)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorLowercase, "Lowercase (a-z)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorNumbers, "Numbers (0-9)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorSymbols, "Symbols (!@#$...)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorExcludeSimilar, "Exclude similar (iIlL1oO0)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorExcludeAmbiguous, "Exclude ambiguous ({ } [ ] ( ) / \\ ' \" ` ~ , ; : . < >)")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorNoRepeats, "No repeated characters")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorCustomChars, "Custom characters")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorCustomCharsPlaceholder, "Enter custom character set...")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorCustomCharsHint, "Leave empty to use character sets above")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorGenerateMultiple, "Generate Multiple")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorAboutTitle, "About Password Security")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorAboutItem1, "Use at least 12 characters for adequate security")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorAboutItem2, "Include a mix of uppercase, lowercase, numbers, and symbols")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorAboutItem3, "Avoid using personal information or common words")
	t.AddTranslation(config.SupportedLangEN, KeyToolsPasswordGeneratorAboutItem4, "Use a different password for each account")

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
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortInvalidPort, "Número de puerto inválido. Por favor ingrese un número entre 1 y 65535.")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortScanning, "Escaneando puerto")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortProcesses, "Procesos:")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortKilled, "PIDs eliminados")
	t.AddTranslation(config.SupportedLangES, KeyToolsKillPortFreed, "¡Liberado!")

	// Commander Tool
	t.AddTranslation(config.SupportedLangES, KeyToolsCommander, "Explorador de Archivos")
	t.AddTranslation(config.SupportedLangES, KeyToolsCommanderTitle, "🗂️ Explorador de Archivos")
	t.AddTranslation(config.SupportedLangES, KeyToolsCommanderDesc, "Gestor de archivos de panel dual (demo frontend con FS virtual)")
	t.AddTranslation(config.SupportedLangES, KeyToolsCommanderCopy, "Copiar")
	t.AddTranslation(config.SupportedLangES, KeyToolsCommanderMove, "Mover")
	t.AddTranslation(config.SupportedLangES, KeyToolsCommanderDelete, "Eliminar")
	t.AddTranslation(config.SupportedLangES, KeyToolsCommanderCopyAlert, "Copiar izquierda → derecha (simulación)")
	t.AddTranslation(config.SupportedLangES, KeyToolsCommanderMoveAlert, "Mover izquierda → derecha (simulación)")
	t.AddTranslation(config.SupportedLangES, KeyToolsCommanderDeleteAlert, "Eliminar seleccionado (simulación)")

	// Terminal Tool
	t.AddTranslation(config.SupportedLangES, KeyToolsTerminal, "Terminal Web")
	t.AddTranslation(config.SupportedLangES, KeyToolsTerminalTitle, "💻 Terminal Web")
	t.AddTranslation(config.SupportedLangES, KeyToolsTerminalDesc, "Emulador de terminal PC (shell JS frontend)")
	t.AddTranslation(config.SupportedLangES, KeyToolsTerminalWelcome, "Bienvenido a shell simulado v1.0")
	t.AddTranslation(config.SupportedLangES, KeyToolsTerminalPlaceholder, "Escribe comandos: ls, cd dir1, pwd, cat file1.txt, clear...")
	t.AddTranslation(config.SupportedLangES, KeyToolsTerminalCleared, "Terminal limpiado")
	t.AddTranslation(config.SupportedLangES, KeyToolsTerminalCommandNotFound, "comando no encontrado")

	// About Page
	t.AddTranslation(config.SupportedLangES, KeyAboutMissionDesc, "IT Tools es una colección de herramientas en línea diseñadas para facilitar el trabajo diario de los desarrolladores. Desde el formateo de JSON hasta la generación de hash, encontrarás herramientas útiles para tu flujo de trabajo.")
	t.AddTranslation(config.SupportedLangES, KeyAboutFeaturesFree, "Completamente gratis")
	t.AddTranslation(config.SupportedLangES, KeyAboutFeaturesNoReg, "No requiere registro")
	t.AddTranslation(config.SupportedLangES, KeyAboutFeaturesPrivacy, "Privacidad garantizada")
	t.AddTranslation(config.SupportedLangES, KeyAboutFeaturesOpenSource, "Código abierto")
	t.AddTranslation(config.SupportedLangES, KeyAboutTechDesc, "Este proyecto está construido con Go (Golang) y utiliza plantillas HTML para el frontend.")

	// Common Tool Keys
	t.AddTranslation(config.SupportedLangES, KeyToolsInput, "Entrada")
	t.AddTranslation(config.SupportedLangES, KeyToolsOutput, "Salida")
	t.AddTranslation(config.SupportedLangES, KeyToolsCopyOutput, "Copiar Salida")
	t.AddTranslation(config.SupportedLangES, KeyToolsGenerate, "Generar")
	t.AddTranslation(config.SupportedLangES, KeyToolsCopy, "Copiar")
	t.AddTranslation(config.SupportedLangES, KeyToolsCopyAll, "Copiar Todo")
	t.AddTranslation(config.SupportedLangES, KeyToolsClear, "Limpiar")
	t.AddTranslation(config.SupportedLangES, KeyToolsChars, "Caracteres")
	t.AddTranslation(config.SupportedLangES, KeyToolsWords, "Palabras")
	t.AddTranslation(config.SupportedLangES, KeyToolsLines, "Líneas")
	t.AddTranslation(config.SupportedLangES, KeyToolsEmptyLines, "Líneas Vacías")
	t.AddTranslation(config.SupportedLangES, KeyToolsReplace, "Reemplazar")
	t.AddTranslation(config.SupportedLangES, KeyToolsStatsGenerated, "Generados")

	// UUID Generator Tool
	t.AddTranslation(config.SupportedLangES, KeyToolsUuidGenerator, "Generador de UUID")
	t.AddTranslation(config.SupportedLangES, KeyToolsUuidGeneratorTitle, "Generador de UUID")
	t.AddTranslation(config.SupportedLangES, KeyToolsUuidGeneratorDesc, "Genera tokens UUID v4 compatibles con RFC 4122 usando crypto.randomUUID()")
	t.AddTranslation(config.SupportedLangES, KeyToolsUuidGeneratorCount, "Número de UUIDs")
	t.AddTranslation(config.SupportedLangES, KeyToolsUuidGeneratorUppercase, "Mayúsculas")
	t.AddTranslation(config.SupportedLangES, KeyToolsUuidGeneratorHyphens, "Incluir guiones")
	t.AddTranslation(config.SupportedLangES, KeyToolsUuidGeneratorBraces, "Incluir llaves")

	// Text Utils Tool
	t.AddTranslation(config.SupportedLangES, KeyToolsTextUtils, "Utilidades de Texto")
	t.AddTranslation(config.SupportedLangES, KeyToolsTextUtilsTitle, "Utilidades de Texto")
	t.AddTranslation(config.SupportedLangES, KeyToolsTextUtilsDesc, "Transforma y manipula texto con varias operaciones")
	t.AddTranslation(config.SupportedLangES, KeyToolsTextUtilsLineOps, "Operaciones de Línea")
	t.AddTranslation(config.SupportedLangES, KeyToolsTextUtilsCaseOps, "Operaciones de Mayúsculas")
	t.AddTranslation(config.SupportedLangES, KeyToolsTextUtilsWhitespace, "Espacios en Blanco")
	t.AddTranslation(config.SupportedLangES, KeyToolsTextUtilsFindReplace, "Buscar y Reemplazar")
	t.AddTranslation(config.SupportedLangES, KeyToolsTextUtilsFindPlaceholder, "Buscar texto...")
	t.AddTranslation(config.SupportedLangES, KeyToolsTextUtilsReplacePlaceholder, "Reemplazar con...")
	t.AddTranslation(config.SupportedLangES, KeyToolsTextUtilsCaseSensitive, "Distinguir mayúsculas")

	// Text Utils Additional Keys
	t.AddTranslation(config.SupportedLangES, KeyToolsReverseLines, "Invertir Líneas")
	t.AddTranslation(config.SupportedLangES, KeyToolsSortLines, "Ordenar Líneas")
	t.AddTranslation(config.SupportedLangES, KeyToolsUniqueLines, "Líneas Únicas")
	t.AddTranslation(config.SupportedLangES, KeyToolsAddLineNumbers, "Añadir Números de Línea")
	t.AddTranslation(config.SupportedLangES, KeyToolsUppercase, "MAYÚSCULAS")
	t.AddTranslation(config.SupportedLangES, KeyToolsLowercase, "minúsculas")
	t.AddTranslation(config.SupportedLangES, KeyToolsTitleCase, "Título")
	t.AddTranslation(config.SupportedLangES, KeyToolsSentenceCase, "Oración")
	t.AddTranslation(config.SupportedLangES, KeyToolsTrimStart, "Recortar Inicio")
	t.AddTranslation(config.SupportedLangES, KeyToolsTrimEnd, "Recortar Fin")
	t.AddTranslation(config.SupportedLangES, KeyToolsTrimBoth, "Recortar Ambos")
	t.AddTranslation(config.SupportedLangES, KeyToolsClearAll, "Limpiar Todo")

	// Base64 Tool
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64, "Base64")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64Title, "Codificador/Decodificador Base64")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64Desc, "Codifica y decodifica cadenas Base64")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64Encode, "Codificar")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64Decode, "Decodificar")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64UrlSafe, "Seguro para URL")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64IncludePadding, "Incluir relleno")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64InputPlaceholder, "Ingrese texto para codificar...")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64OutputPlaceholder, "El resultado codificado/decodificado aparecerá aquí...")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64FileUpload, "O subir un archivo")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64EncodeBtn, "Codificar")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64DecodeBtn, "Decodificar")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64Download, "Descargar")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64AboutTitle, "Acerca de Base64")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64AboutDesc, "Base64 es un grupo de esquemas de codificación de binario a texto que representan datos binarios en formato de cadena ASCII.")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64AboutItem1, "Comúnmente usado para codificar datos en URLs, correos electrónicos y almacenar datos complejos en campos de texto")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64AboutItem2, "Cada dígito Base64 representa exactamente 6 bits de datos")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64AboutItem3, "La cadena Base64 resultante es aproximadamente 33% más grande que los datos binarios originales")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64NoInput, "Por favor ingrese texto para procesar")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64EncodeSuccess, "¡Codificado exitosamente!")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64EncodeError, "Error al codificar: entrada inválida")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64DecodeSuccess, "¡Decodificado exitosamente!")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64DecodeError, "Error al decodificar: cadena Base64 inválida")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64Cleared, "Limpiado")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64NothingToCopy, "Nada que copiar")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64CopySuccess, "¡Copiado al portapapeles!")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64CopyError, "Error al copiar al portapapeles")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64NothingToDownload, "Nada que descargar")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64DownloadStarted, "¡Descarga iniciada!")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64DownloadError, "Error al descargar archivo")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64FileTooLarge, "Archivo muy grande. El tamaño máximo es 5MB.")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64FileEncoded, "¡Archivo codificado exitosamente!")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64FileError, "Error al procesar archivo")
	t.AddTranslation(config.SupportedLangES, KeyToolsBase64FileReadError, "Error al leer archivo")

	// Password Generator Tool
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGenerator, "Generador de Contraseñas")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorTitle, "🔐 Generador de Contraseñas")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorDesc, "Genera contraseñas seguras aleatorias")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorLength, "Longitud")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorUppercase, "Mayúsculas (A-Z)")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorLowercase, "Minúsculas (a-z)")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorNumbers, "Números (0-9)")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorSymbols, "Símbolos (!@#$...)")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorExcludeSimilar, "Excluir similares (iIlL1oO0)")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorExcludeAmbiguous, "Excluir ambiguos ({ } [ ] ( ) / \\ ' \" ` ~ , ; : . < >)")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorNoRepeats, "Sin caracteres repetidos")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorCustomChars, "Caracteres personalizados")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorCustomCharsPlaceholder, "Ingrese conjunto de caracteres personalizado...")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorCustomCharsHint, "Dejar vacío para usar los conjuntos de caracteres anteriores")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorGenerateMultiple, "Generar Múltiples")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorAboutTitle, "Acerca de la Seguridad de Contraseñas")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorAboutItem1, "Use al menos 12 caracteres para una seguridad adecuada")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorAboutItem2, "Incluya una combinación de mayúsculas, minúsculas, números y símbolos")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorAboutItem3, "Evite usar información personal o palabras comunes")
	t.AddTranslation(config.SupportedLangES, KeyToolsPasswordGeneratorAboutItem4, "Use una contraseña diferente para cada cuenta")
}
