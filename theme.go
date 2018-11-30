package main

// Background set background color
type Background Color

// NewBackground creates a new background
func NewBackground(background Color) *Background {
	bg := Background(background)
	return &bg
}

// Draw draws to lemonbar
func (b Background) Draw() string { return "%{B#" + Color(b).Hex() + "}" }

// DrawWith set background for parent
func (b Background) DrawWith(parent string) string {
	return "%{B#" + Color(b).Hex() + "}" + parent + "%{B-}"
}

// Foreground set foreground color
type Foreground Color

// NewForeground creates a new foreground
func NewForeground(foreground Color) *Foreground {
	fg := Foreground(foreground)
	return &fg
}

// Draw draws to lemonbar
func (b Foreground) Draw() string { return "%{F#" + Color(b).Hex() + "}" }

// DrawWith set forground for parent
func (b Foreground) DrawWith(parent string) string {
	return "%{F#" + Color(b).Hex() + "}" + parent + "%{F-}"
}

// Underline set underline color
type Underline Color

// NewUnderline creates a new underline
func NewUnderline(underline Color) *Underline {
	ul := Underline(underline)
	return &ul
}

// Draw draws to lemonbar, use draw with instead (see Bind)
func (b Underline) Draw() string { return "%{+u}%{U#" + Color(b).Hex() + "}" }

// DrawWith underline parent
func (b Underline) DrawWith(parent string) string {
	return "%{+u}%{U#" + Color(b).Hex() + "}" + parent + "%{-u}"
}
