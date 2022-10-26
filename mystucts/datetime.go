package mystructs

import (
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

// Convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format(time.RFC3339), nil
}

// You could also use the standard Stringer interface
func (date *DateTime) String() string {
	return date.Time.Format(time.RFC3339) // Redundant, just for example
}

// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	csv = strings.ReplaceAll(csv, " ", "T") + "Z"
	date.Time, err = time.Parse(time.RFC3339, csv)
	return err
}
