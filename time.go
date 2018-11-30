package main

import "time"

var _ Updatable = (*Time)(nil)

// Time is a widget representing current time
type Time struct {
	time time.Time
}

// NewTime creates a new time widget
func NewTime() Widget {
	return &Time{time: time.Now()}
}

// Update updates widget time
func (t *Time) Update() (bool, error) {
	newtime := time.Now()
	if newtime.Minute() == t.time.Minute() {
		return false, nil
	}
	t.time = newtime
	return true, nil
}

// Tick refresh rate for time
func (t Time) Tick() uint {
	return 60
}

// Draw draws time to widget
func (t Time) Draw() string { return t.time.Format("Mon 2 Jan, 15:04") }
