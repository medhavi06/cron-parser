package cronparser

import (
	"fmt"
	"sort"
	"strings"

	fieldparser "github.com/medhavi06/cron-parser/pkg/parser"
)

// CronParser handles parsing of cron expressions
type CronParser struct {
	fieldParser Parser
}

// CronResult represents the parsed result of a cron expression
type CronResult struct {
	Minutes     []int
	Hours       []int
	DaysOfMonth []int
	Months      []int
	DaysOfWeek  []int
	Command     string
}

// NewCronParser creates a new CronParser with given options
func NewCronParser() (*CronParser, error) {
	parser := &CronParser{
		fieldParser: &fieldparser.StandardFieldParser{},
	}

	return parser, nil
}

// Validate ensures the result meets basic requirements
func (cr *CronResult) Validate() error {
	if len(cr.Minutes) == 0 || len(cr.Hours) == 0 {
		return fmt.Errorf("minutes and hours cannot be empty")
	}
	return nil
}

// Helper function to sort and remove duplicates
func uniqueSortedInts(input []int) []int {
	unique := make(map[int]bool)
	var result []int
	for _, val := range input {
		if !unique[val] {
			unique[val] = true
			result = append(result, val)
		}
	}
	sort.Ints(result)
	return result
}

// Normalize sorts and removes duplicates from the result
func (cr *CronResult) Normalize() {
	cr.Minutes = uniqueSortedInts(cr.Minutes)
	cr.Hours = uniqueSortedInts(cr.Hours)
	cr.DaysOfMonth = uniqueSortedInts(cr.DaysOfMonth)
	cr.Months = uniqueSortedInts(cr.Months)
	cr.DaysOfWeek = uniqueSortedInts(cr.DaysOfWeek)
}

// Parse processes a full cron expression
func (cp *CronParser) Parse(expression string) (*CronResult, error) {
	// Split the expression into parts
	parts := strings.Split(expression, " ")
	// parts := strings.Fields(expression)
	if len(parts) < 6 {
		return nil, fmt.Errorf("invalid cron expression: not enough fields")
	}

	result := &CronResult{
		Command: strings.Join(parts[5:], " "),
	}

	// Define field specifications
	fields := []struct {
		name     string
		value    *[]int
		min, max int
	}{
		{"minute", &result.Minutes, 0, 59},
		{"hour", &result.Hours, 0, 23},
		{"day of month", &result.DaysOfMonth, 1, 31},
		{"month", &result.Months, 1, 12},
		{"day of week", &result.DaysOfWeek, 0, 6},
	}

	// Parse each field
	for i, field := range fields {
		values, err := cp.fieldParser.Parse(parts[i], field.min, field.max)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", field.name, err)
		}
		*field.value = values
	}

	// Normalize the result
	result.Normalize()

	// Validate the result
	if err := result.Validate(); err != nil {
		return nil, err
	}

	return result, nil
}
