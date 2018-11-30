package main

var _ Updatable = (*Bind)(nil)

// Widget is a lemonbar widget that is renderable
type Widget interface {
	// Draw is called to draw text, takes parent widget if any
	Draw() string
}

// Updatable is a widget that needs to update values
type Updatable interface {
	// Update is called before Draw, it is used to update values
	Update() (bool, error)
	// Tick returns tick rate
	Tick() uint
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
func (b *Bind) Update() (res bool, err error) {
	if parent, ok := b.parent.(Updatable); ok {
		res, err = parent.Update()
		if err != nil {
			return
		}
	}
	if nested, ok := b.nested.(Updatable); ok {
		var newres bool
		newres, err = nested.Update()
		if err != nil {
			return
		}
		if newres {
			res = true
		}
	}
	return
}

// Tick return ticks
func (b Bind) Tick() uint {
	if parent, ok := b.parent.(Updatable); ok {
		return parent.Tick()
	}
	if nested, ok := b.nested.(Updatable); ok {
		return nested.Tick()
	}
	return 0
}

// Draw draws to lemonbar
func (b Bind) Draw() string {
	return b.nested.DrawWith(b.parent.Draw())
}
