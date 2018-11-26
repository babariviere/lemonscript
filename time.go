package main

import "time"

// Time is a widget representing current time
type Time struct {
	time time.Time
}

// NewTime creates a new time widget
func NewTime() Widget {
	return &Time{time: time.Now()}
}

// Update updates widget time
func (t *Time) Update() error {
	t.time = time.Now()
	return nil
}

// Draw draws time to widget
func (t Time) Draw() string { return t.time.Format("Mon 2 Jan 15:04:05") }
