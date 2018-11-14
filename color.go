package main

import (
	"encoding/hex"
	"errors"
)

// Color is a color
type Color struct {
	r, g, b, a byte
	hex        string
}

// NewRGB creates a new color from rgb
func NewRGB(r, g, b byte) Color {
	return Color{r: r, g: g, b: b, a: 255}
}

// NewRGBA creates a new color from rgba
func NewRGBA(r, g, b, a byte) Color {
	return Color{r: r, g: g, b: b, a: a}
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
		c.a = b[0]
		c.r = b[1]
		c.g = b[2]
		c.b = b[3]
	} else if len(b) == 3 {
		c.a = 255
		c.r = b[0]
		c.g = b[1]
		c.b = b[2]
	} else {
		return c, errors.New("expected 3 elements")
	}
	return c, nil
}

// Hex converts color to hex
func (c Color) Hex() string {
	if c.hex != "" {
		return c.hex
	}
	if c.a != 255 {
		c.hex = hex.EncodeToString([]byte{c.a, c.r, c.g, c.b})
	} else {
		c.hex = hex.EncodeToString([]byte{c.r, c.g, c.b})
	}
	return c.hex
}
