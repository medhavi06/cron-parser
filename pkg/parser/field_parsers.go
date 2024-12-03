package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// StandardFieldParser implements the FieldParser interface
type StandardFieldParser struct{}

func (fp *StandardFieldParser) Parse(field string, min, max int) ([]int, error) {
	// Handle wildcard
	if field == "*" {
		return generateRange(min, max), nil
	}

	// Handle step values
	if strings.Contains(field, "/") {
		return fp.parseStepValue(field, min, max)
	}

	// Handle ranges and lists
	if strings.Contains(field, ",") || strings.Contains(field, "-") {
		return fp.parseListOrRange(field, min, max)
	}

	// Single value
	val, err := strconv.Atoi(field)
	if err != nil {
		return nil, fmt.Errorf("invalid value: %s", field)
	}
	if val < min || val > max {
		return nil, fmt.Errorf("value out of range: %d (min: %d, max: %d)", val, min, max)
	}
	return []int{val}, nil
}

func (fp *StandardFieldParser) parseStepValue(field string, min, max int) ([]int, error) {
	parts := strings.Split(field, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid step value: %s", field)
	}

	step, err := strconv.Atoi(parts[1])
	if err != nil || step <= 0 {
		return nil, fmt.Errorf("invalid step: %s", parts[1])
	}

	var startVal, endVal int
	if parts[0] == "*" {
		startVal = min
		endVal = max
	} else {
		startVal, err = strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid start value: %s", parts[0])
		}
		endVal = max
	}

	var result []int
	for i := startVal; i <= endVal; i += step {
		result = append(result, i)
	}
	return result, nil
}

func (fp *StandardFieldParser) parseListOrRange(field string, min, max int) ([]int, error) {
	var result []int
	parts := strings.Split(field, ",")

	for _, part := range parts {
		// Check if it's a range
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid start of range: %s", rangeParts[0])
			}

			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid end of range: %s", rangeParts[1])
			}

			if start < min || end > max || start > end {
				return nil, fmt.Errorf("range out of bounds: %s", part)
			}

			result = append(result, generateRange(start, end)...)
		} else {
			// Single value in list
			val, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid value: %s", part)
			}
			if val < min || val > max {
				return nil, fmt.Errorf("value out of range: %d", val)
			}
			result = append(result, val)
		}
	}

	return result, nil
}

// Helper function to generate a range of values
func generateRange(start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}
