package main

import (
	"fmt"
	"os"

	"github.com/medhavi06/cron-parser/pkg/cronparser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cron-parser \"<cron expression>\"")
		os.Exit(1)
	}

	// Create parser with default options
	options := cronparser.DefaultOptions()
	parser, err := cronparser.NewCronParser(options)
	if err != nil {
		fmt.Println("Error initializing parser:", err)
		os.Exit(1)
	}

	// Parse the cron expression
	result, err := parser.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing cron expression:", err)
		os.Exit(1)
	}

	// Print results
	printCronResult(result)
}

func printCronResult(result *cronparser.CronResult) {
	fieldWidth := 14
	printField := func(name string, values []int) {
		fmt.Printf("%-*s %s\n", fieldWidth, name, formatValues(values))
	}

	printField("minute", result.Minutes)
	printField("hour", result.Hours)
	printField("day of month", result.DaysOfMonth)
	printField("month", result.Months)
	printField("day of week", result.DaysOfWeek)
	fmt.Printf("%-*s %s\n", fieldWidth, "command", result.Command)
}

func formatValues(values []int) string {
	if len(values) == 0 {
		return ""
	}

	result := ""
	for _, val := range values {
		result += fmt.Sprintf("%d ", val)
	}
	return result[:len(result)-1]
}
