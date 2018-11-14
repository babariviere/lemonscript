package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
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
