package main

type Widget interface {
	// Update is called before Draw, it is used to update values
	Update() error
	// Draw is called to draw text
	Draw() string
}
