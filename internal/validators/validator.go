package validators

import (
	"fmt"
	"unicode"
)

// ValidateInput checks the input for potential issues
func ValidateInput(input string) error {
	if len(input) == 0 {
		return fmt.Errorf("input cannot be empty")
	}
	for _, char := range input {
		if !isValidChar(char) {
			return fmt.Errorf("invalid character found: %c", char)
		}
	}

	return nil
}

// isValidChar checks if a character is allowed in cron expressions
func isValidChar(char rune) bool {
	return unicode.IsDigit(char) ||
		char == '*' ||
		char == '/' ||
		char == '-' ||
		char == ',' ||
		unicode.IsSpace(char)
}
