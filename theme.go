package main

// Background set background color
type Background Color

// NewBackground creates a new background
func NewBackground(background Color) Widget {
	bg := Background(background)
	return &bg
}

// Update does nothing
func (b *Background) Update() error { return nil }

// Draw draws to lemonbar
func (b Background) Draw() string { return "%{B#" + Color(b).Hex() + "}" }

// Foreground set foreground color
type Foreground Color

// NewForeground creates a new foreground
func NewForeground(foreground Color) Widget {
	fg := Foreground(foreground)
	return &fg
}

// Update does nothing
func (b *Foreground) Update() error { return nil }

// Draw draws to lemonbar
func (b Foreground) Draw() string { return "%{F#" + Color(b).Hex() + "}" }

// Underline set underline color
type Underline Color

// NewUnderline creates a new underline
func NewUnderline(underline Color) Widget {
	ul := Underline(underline)
	return &ul
}

// Update does nothing
func (b *Underline) Update() error { return nil }

// Draw draws to lemonbar
func (b Underline) Draw() string { return "%{U#" + Color(b).Hex() + "}" }
