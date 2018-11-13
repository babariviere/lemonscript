package main

import (
	"fmt"
	"time"
)

var background = NewRGBA(41, 45, 62, 204)

var widgets = []Widget{
	NewBackground(background),
	NewAlign(AlignCenter),
	NewText("Hello World"),
	NewPadding(2),
	NewAlign(AlignRight),
	NewBattery("", "%"),
	NewPadding(2),
	NewTime(),
}

func main() {
	for {
		var buf string
		for _, widget := range widgets {
			if err := widget.Update(); err != nil {
				fmt.Println(err)
				goto EndLoop
			}
			buf += widget.Draw()
		}
		fmt.Println(buf)
	EndLoop:
		time.Sleep(1000)
	}
}
