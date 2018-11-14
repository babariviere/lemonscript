package main

// Widget is a lemonbar widget that is renderable
type Widget interface {
	// Update is called before Draw, it is used to update values
	Update() error
	// Draw is called to draw text, takes parent widget if any
	Draw() string
}

// NestedWidget is a widget that can have optional result
type NestedWidget interface {
	Widget
	// DrawWith draws with result from parent as parameter
	DrawWith(string) string
}

// Empty is an empty widget
type Empty struct{}

// NewEmpty creates a new empty Widget
func NewEmpty() Widget { return &Empty{} }

// Update does nothing
func (e Empty) Update() error { return nil }

// Draw does nothing
func (e Empty) Draw() string { return "" }

// DrawWith return given value
func (e Empty) DrawWith(res string) string { return res }

// Bind allow to bind output from parent to nested
type Bind struct {
	parent Widget
	nested NestedWidget
}

// NewBind create a new bind widget
func NewBind(parent Widget, nested NestedWidget) *Bind {
	return &Bind{parent: parent, nested: nested}
}

// Update updates all widgets
func (b *Bind) Update() error {
	if err := b.parent.Update(); err != nil {
		return err
	}
	if err := b.nested.Update(); err != nil {
		return err
	}
	return nil
}

// Draw draws to lemonbar
func (b Bind) Draw() string {
	return b.nested.DrawWith(b.parent.Draw())
}
