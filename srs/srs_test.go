package srs

import (
	"testing"
	"time"
)

func TestDayComparation(t *testing.T) {
	dayA := NewDay(time.Now())
	dayB := NewDay(time.Now())

	if !dayA.Equal(dayB) || !dayB.Equal(dayA) {
		t.Log("Day A:", dayA)
		t.Log("Day B:", dayB)
		t.Error("THe days must be the same.")
	}

	dayB.Next()

	if dayA.Equal(dayB) || dayB.Equal(dayA) {
		t.Log("Day A:", dayA)
		t.Log("Day B:", dayB)
		t.Error("The days should not be the same.")
	}
}

func TestDayStep(t *testing.T) {
	tt, err := time.Parse(time.Layout, time.Layout)
	if err != nil {
		t.Fatal(err)
	}

	day := NewDay(tt)
	t.Log("First Day:", day)

	d := day.time.Day()
	day.Step(4)
	t.Log("Stepping 4:", day)

	if day.time.Day() != d+4 {
		t.Error("They did not skip the expected days.")
	}
}

func TestDayRound(t *testing.T) {
	day := NewDay(time.Now())

	hour, min, sec := day.time.Clock()

	if hour+min+sec != 0 {
		t.Error("The day clock is different from zero.")
	}
}
