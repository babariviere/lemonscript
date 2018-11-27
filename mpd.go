package main

import (
	gompd "github.com/fhs/gompd/mpd"
)

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
func (m *MPD) Update() (err error) {
	m.status, err = m.client.Status()
	if err != nil {
		return
	}
	m.song, err = m.client.CurrentSong()
	return
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
