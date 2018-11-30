package main

import "fmt"

// Padding adds space between widgets
type Padding string

// NewPadding creates a new padding
func NewPadding(i int) Widget {
	p := Padding(fmt.Sprint(i))
	return &p
}

// Draw draws to lemonbar
func (p Padding) Draw() string { return "%{O" + string(p) + "}" }
