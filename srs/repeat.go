package srs

import "time"

// Repeat implements Spaced Repetition.
type Repeat struct {
	Object any

	lastDay      Day
	interval     uint
	moveInterval bool // move up next interval?
	history      []Day
}

// NewRepeat initializes a Repeat.
func NewRepeat(object any) *Repeat {
	today := NewDay(time.Now())

	return &Repeat{
		Object:       object,
		lastDay:      today,
		interval:     0,
		moveInterval: true,
		history:      []Day{},
	}
}

// LastReview returns the last day with reviews.
func (r *Repeat) LastReview() Day { return r.lastDay }

// Interval returns the actual interval.
func (r *Repeat) Interval() uint { return r.interval }

// MoveUpNextInterval informs if the interval will be increased in the next review.
func (r *Repeat) MoveUpNextInterval() bool { return r.moveInterval }

// History returns the review history.
func (r *Repeat) History() []Day { return r.history }

// WaitingReview informs if there is a pendient review.
func (r *Repeat) WaitingReview() bool {
	return r.lastDay.Step(int64(r.interval)) == NewDay(time.Now())
}

// ReviewedNow marks as reviewed the current day.
func (r *Repeat) ReviewedNow() {
	if len(r.history) == 0 {
		r.interval, r.moveInterval = r.nextIntervalReview()
	} else if !r.history[len(r.history)-1].Equal(r.lastDay) {
		r.interval, r.moveInterval = r.nextIntervalReview()
	}

	r.history = append(r.history, r.lastDay)
	r.lastDay = NewDay(time.Now())
}

func (r *Repeat) nextIntervalReview() (interval uint, moveInterval bool) {
	if !r.moveInterval {
		return r.interval, true
	}

	return r.interval + 1, false
}
