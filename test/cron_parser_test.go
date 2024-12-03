package unit

import (
	"testing"

	"github.com/medhavi06/cron-parser/pkg/cronparser"
)

func TestCronParser(t *testing.T) {
	parser, err := cronparser.NewCronParser(nil)
	if err != nil {
		t.Fatalf("Failed to create parser: %v", err)
	}

	testCases := []struct {
		name           string
		input          string
		expectError    bool
		expectedResult *cronparser.CronResult
	}{
		{
			name:  "Standard Expression",
			input: "*/15 0 1,15 * 1-5 /usr/bin/find",
			expectedResult: &cronparser.CronResult{
				Minutes:     []int{0, 15, 30, 45},
				Hours:       []int{0},
				DaysOfMonth: []int{1, 15},
				Months:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DaysOfWeek:  []int{1, 2, 3, 4, 5},
				Command:     "/usr/bin/find",
			},
		},
		{
			name:        "Insufficient Fields",
			input:       "* * * *",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parser.Parse(tc.input)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Detailed comparison
			if len(result.Minutes) != len(tc.expectedResult.Minutes) {
				t.Errorf("Minutes mismatch. Got %v, want %v", result.Minutes, tc.expectedResult.Minutes)
			}

			if result.Command != tc.expectedResult.Command {
				t.Errorf("Command mismatch. Got %s, want %s", result.Command, tc.expectedResult.Command)
			}
		})
	}
}
