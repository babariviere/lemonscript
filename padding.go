package main

import "fmt"

// Padding adds space between widgets
type Padding int

// NewPadding creates a new padding
func NewPadding(i int) Widget {
	p := Padding(i)
	return &p
}

// Update does nothing
func (p *Padding) Update() error { return nil }

// Draw draws to lemonbar
func (p Padding) Draw() string { return "%{O" + fmt.Sprint(int(p)) + "}" }
