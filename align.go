package main

const (
	// AlignLeft align to left
	AlignLeft Align = 'l'
	// AlignCenter align to center
	AlignCenter Align = 'c'
	// AlignRight align to right
	AlignRight Align = 'r'
)

// Align aligns text to given parameter
type Align byte

// NewAlign creates a new align widget
func NewAlign(a Align) Widget { return &a }

// Draw draws to lemobar
func (a Align) Draw() string { return "%{" + string(a) + "}" }
