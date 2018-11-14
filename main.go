package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var background = NewRGBA(41, 45, 62, 204)
var focused = background
var unfocused, _ = NewHex("#434758")
var urgent, _ = NewHex("#f07178")

var padding = NewPadding(10)

var enableUnderline = NewText("%{+u}")

var widgets = []Widget{
	NewBackground(background),
	NewI3(
		NewCombined(
			NewCombined(NewBackground(focused), padding),
			padding,
		),
		NewCombined(
			NewCombined(NewBackground(unfocused), padding),
			padding,
		),
		NewUnderline(
			urgent,
		),
	),
	NewBackground(background),
	NewAlign(AlignCenter),
	padding,
	NewTime(),
	padding,
	NewAlign(AlignRight),
	NewBattery(NewCombined(NewEmpty(), NewText("%"))),
	padding,
}

func drawLoop() {
	var buf string
	for _, widget := range widgets {
		if err := widget.Update(); err != nil {
			fmt.Println(err)
			return
		}
		buf += widget.Draw()
	}
	fmt.Println(buf)
}

func cpuProfile() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	for i := 0; i < 10000; i++ {
		drawLoop()
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "cpu.prof" {
		cpuProfile()
		return
	}
	for {
		drawLoop()
		time.Sleep(500 * time.Millisecond)
	}
}
