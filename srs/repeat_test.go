package srs

import (
	"reflect"
	"testing"
	"time"
)

func TestRepeatGetters(t *testing.T) {
	srr := NewRepeat(nil)
	day := NewDay(time.Now())

	if !srr.LastReview().Equal(day) {
		t.Error("The `lastDay` attribute is not set correctly.")
	}

	if srr.Interval() != 0 {
		t.Error("The `interval` attribute is not set correctly.")
	}

	if !srr.MoveUpNextInterval() {
		t.Error("The `moveInterval` attribute is not set correctly.")
	}

	if !reflect.DeepEqual(srr.History(), []Day{}) {
		t.Error("The history is incorrect.")
	}
}

func TestRepeat_WaitingReview(t *testing.T) {
	srr := NewRepeat(nil)

	if !srr.WaitingReview() {
		t.Error("A review was expected.")
	}

	srr.ReviewedNow()

	if srr.WaitingReview() {
		t.Error("No review expected.")
	}
}
