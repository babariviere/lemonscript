package main

// Clickable is a clickable widget
type Clickable struct {
	action string
}

// NewClickable creates a new clickable widget
func NewClickable(action string) *Clickable {
	return &Clickable{action: action}
}

// Update does nothing
func (c *Clickable) Update() error { return nil }

// Draw does nothing
func (c Clickable) Draw() string { return "" }

// DrawWith draws to lemonbar
func (c Clickable) DrawWith(res string) string {
	return "%{A:" + c.action + ":}" + res + "%{A}"
}
