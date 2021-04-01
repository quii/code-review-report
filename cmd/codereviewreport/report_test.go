package main

import (
	"testing"
	"time"
)

func TestMonday(t *testing.T) {
	date := LastMonday()

	if date.Weekday() != time.Monday {
		t.Errorf("Last Monday is not a Monday")
	}

	if date.After(time.Now()) {
		t.Errorf("Last Monday is in the future")
	}
}
