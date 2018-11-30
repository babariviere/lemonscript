package main

import "math"

var _ Updatable = (*Combined)(nil)

// Combined is a widget that combine 2 widget
type Combined struct {
	prefix, suffix Widget
}

// NewCombined creates a new combined widget
func NewCombined(prefix, suffix Widget) *Combined {
	return &Combined{prefix: prefix, suffix: suffix}
}

// Update updates nested widgets
func (c *Combined) Update() (res bool, err error) {
	if prefix, ok := c.prefix.(Updatable); ok {
		res, err = prefix.Update()
		if err != nil {
			return
		}
	}
	if suffix, ok := c.suffix.(Updatable); ok {
		var newres bool
		newres, err = suffix.Update()
		if err != nil {
			return
		}
		if newres {
			res = newres
		}
	}
	return
}

// Tick refresh rate for combined
func (c Combined) Tick() (rate uint) {
	rate = math.MaxUint32
	if prefix, ok := c.prefix.(Updatable); ok {
		rate = prefix.Tick()
	}
	if suffix, ok := c.prefix.(Updatable); ok {
		srate := suffix.Tick()
		if srate < rate {
			rate = srate
		}
	}
	return
}

// Draw draws to lemonbar
func (c Combined) Draw() string {
	return c.prefix.Draw() + c.suffix.Draw()
}

// DrawWith combine parent results with prefix and suffix
func (c Combined) DrawWith(res string) string {
	return c.prefix.Draw() + res + c.suffix.Draw()
}
