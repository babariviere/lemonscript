package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var background = NewRGBA(41, 45, 62, 204)
var focused, _ = NewHex("#434758")
var unfocused = background
var urgent, _ = NewHex("#f07178")

var padding = NewPadding(10)

var enableUnderline = NewText("%{+u}")

var mpd, _ = NewMPD("127.0.0.1:6600")

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
	// Too much CPU usage for now
	// TODO: introduce tick system
	NewInterface("enp61s0"),
	padding,
	NewInterface("wlo1"),
	padding,
	mpd,
	padding,
	// TODO: implements trigger system to get text under certain commands
	NewBattery(NewCombined(NewEmpty(), NewText("%"))),
	padding,
}

var tick uint

func drawLoop() {
	tick++
	var redraw bool
	for _, widget := range widgets {
		if wid, ok := widget.(Updatable); ok {
			widtick := wid.Tick()
			if tick%widtick != 0 || widtick == 0 {
				continue
			}
			rd, err := wid.Update()
			if err != nil {
				fmt.Println(err)
				return
			}
			if rd {
				redraw = true
			}
		}
	}
	if redraw {
		var buf string
		//buf += fmt.Sprint("tick: ", tick, " ")
		for _, widget := range widgets {
			buf += widget.Draw()
		}
		fmt.Println(buf)
	}
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
		time.Sleep(100 * time.Millisecond)
	}
}
