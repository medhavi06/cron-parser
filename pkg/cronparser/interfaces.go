package cronparser

// Parser defines the interface for parsing individual cron fields
type Parser interface {
	Parse(field string, min, max int) ([]int, error)
}
