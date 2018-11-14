package main

import (
	"fmt"
	"time"
)

var background = NewRGBA(41, 45, 62, 204)

var widgets = []Widget{
	NewBackground(background),
	NewI3(
		NewCombined(
			NewCombined(NewBackground(NewRGB(10, 10, 10)), NewPadding(2)),
			NewPadding(2),
		),
		NewCombined(
			NewCombined(NewBackground(NewRGB(100, 100, 100)), NewPadding(2)),
			NewPadding(2),
		),
		NewCombined(
			NewCombined(NewBackground(NewRGB(255, 0, 0)), NewPadding(2)),
			NewPadding(2),
		),
	),
	NewBackground(background),
	NewAlign(AlignCenter),
	NewText("Hello World"),
	NewPadding(2),
	NewAlign(AlignRight),
	NewBattery(NewCombined(NewEmpty(), NewText("%"))),
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
		time.Sleep(1 * time.Second)
	}
}
