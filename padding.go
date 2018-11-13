package main

import "strings"

// Padding adds space between widgets
type Padding string

// NewPadding creates a new padding
func NewPadding(i int) Widget {
	p := Padding(strings.Repeat(" ", i))
	return &p
}

// Update does nothing
func (p *Padding) Update() error { return nil }

// Draw draws to lemonbar
func (p Padding) Draw() string { return string(p) }
