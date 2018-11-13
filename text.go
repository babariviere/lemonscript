package main

// Text is a text widget
type Text string

// NewText creates a new text
func NewText(text string) Widget {
	t := Text(text)
	return &t
}

// Update does nothing
func (t *Text) Update() error { return nil }

// Draw draws text to lemonbar
func (t Text) Draw() string { return string(t) }
