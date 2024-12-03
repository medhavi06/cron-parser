package cronparser

// FieldParser defines the interface for parsing individual cron fields
type FieldParser interface {
	Parse(field string, min, max int) ([]int, error)
}
