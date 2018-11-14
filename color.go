package main

import (
	"encoding/hex"
	"errors"
)

// Color is a color
type Color struct {
	R, G, B, A byte
}

// NewRGB creates a new color from rgb
func NewRGB(r, g, b byte) Color {
	return Color{R: r, G: g, B: b, A: 255}
}

// NewRGBA creates a new color from rgba
func NewRGBA(r, g, b, a byte) Color {
	return Color{R: r, G: g, B: b, A: a}
}

// NewHex creates a new color from hex value
func NewHex(h string) (Color, error) {
	if h[0] == '#' {
		h = h[1:]
	}
	var c Color
	b, err := hex.DecodeString(h)
	if err != nil {
		return c, err
	}
	if len(b) >= 4 {
		c.A = b[0]
		c.R = b[1]
		c.G = b[2]
		c.B = b[3]
	} else if len(b) == 3 {
		c.A = 255
		c.R = b[0]
		c.G = b[1]
		c.B = b[2]
	} else {
		return c, errors.New("expected 3 elements")
	}
	return c, nil
}

// Hex converts color to hex
func (c Color) Hex() string {
	if c.A != 255 {
		return hex.EncodeToString([]byte{c.A, c.R, c.G, c.B})
	}
	return hex.EncodeToString([]byte{c.R, c.G, c.B})
}
