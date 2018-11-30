package main

import (
	gompd "github.com/fhs/gompd/mpd"
)

var _ Updatable = (*MPD)(nil)

// MPD services
type MPD struct {
	client *gompd.Client
	status gompd.Attrs
	song   gompd.Attrs
}

// NewMPD creates a new mpd client
func NewMPD(addr string) (Widget, error) {
	client, err := gompd.Dial("tcp", addr)
	return &MPD{client: client}, err
}

// Update fetch music status
func (m *MPD) Update() (redraw bool, err error) {
	var status, song gompd.Attrs
	status, err = m.client.Status()
	if err != nil {
		return
	}
	song, err = m.client.CurrentSong()
	if song["Artist"] != m.song["Artist"] ||
		song["Title"] != m.song["Title"] ||
		status["state"] != m.status["state"] {
		redraw = true
		m.song = song
		m.status = status
	}
	return
}

// Tick refresh rate for mpd
func (m MPD) Tick() uint {
	return 10
}

// Draw draws to lemonbar
func (m MPD) Draw() string {
	if m.status["state"] == "stop" {
		return "stopped"
	}
	var logo string
	if m.status["state"] == "play" {
		logo = "\uf04b"
	} else {
		logo = "\uf04c"
	}
	return logo + "  " + m.song["Artist"] + " - " + m.song["Title"]
}
