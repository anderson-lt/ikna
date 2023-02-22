package srs

import "time"

// Day is a helper for working with absolute days in `time.Time`.
type Day struct {
	time time.Time
}

// NewDay creates and initializes a new `Day`.
func NewDay(t time.Time) Day {
	return Day{t.UTC().Round(time.Hour * 24)}
}

// Equal compares if two `Day`s refer to the same day.
func (d Day) Equal(day Day) bool {
	return d.time.Equal(day.time)
}

// Next is an abbreviation of `Day.Step(1)`.
func (d Day) Next() Day {
	return d.Step(1)
}

// Step add N amount of days.
func (d Day) Step(n int64) Day {
	return Day{d.time.Add(time.Hour * 24 * time.Duration(n))}
}

// String returns a formatted representation.
func (d Day) String() string {
	return d.time.Format("<Day 2 (Jan 2006)>")
}
