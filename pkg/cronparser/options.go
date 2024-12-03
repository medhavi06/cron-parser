package cronparser

// ParserOptions defines configuration for the cron parser
type ParserOptions struct {
	StrictMode               bool
	AllowedSpecialCharacters []rune
	MaxFields                int
}

// DefaultOptions provides standard configuration
func DefaultOptions() *ParserOptions {
	return &ParserOptions{
		StrictMode:               false,
		AllowedSpecialCharacters: []rune{'*', '/', '-', ','},
		MaxFields:                6, // 5 time fields + command
	}
}

// WithStrictMode creates options with strict validation
func WithStrictMode(enabled bool) *ParserOptions {
	opts := DefaultOptions()
	opts.StrictMode = enabled
	return opts
}
