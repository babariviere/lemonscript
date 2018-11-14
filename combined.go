package main

// Combined is a widget that combine 2 widget
type Combined struct {
	prefix, suffix Widget
}

// NewCombined creates a new combined widget
func NewCombined(prefix, suffix Widget) *Combined {
	return &Combined{prefix: prefix, suffix: suffix}
}

// Update updates nested widgets
func (c *Combined) Update() error {
	if err := c.prefix.Update(); err != nil {
		return err
	}
	if err := c.suffix.Update(); err != nil {
		return err
	}
	return nil
}

// Draw draws to lemonbar
func (c Combined) Draw() string {
	return c.prefix.Draw() + c.suffix.Draw()
}

// DrawWith combine parent results with prefix and suffix
func (c Combined) DrawWith(res string) string {
	return c.prefix.Draw() + res + c.suffix.Draw()
}
