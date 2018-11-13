package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Battery is a widget that fetch battery percentage
type Battery struct {
	path           string
	perc           byte
	prefix, suffix string
}

// NewBattery creates a new battery widget
func NewBattery(prefix, suffix string) Widget {
	b := &Battery{prefix: prefix, suffix: suffix}
	files, err := ioutil.ReadDir("/sys/class/power_supply")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.Index(file.Name(), "BAT") == 0 {
			b.path = "/sys/class/power_supply/" + file.Name() + "/capacity"
			return b
		}
	}
	log.Fatal("no battery found")
	return nil
}

// Update fetch battery percentage
func (b *Battery) Update() error {
	content, err := ioutil.ReadFile(b.path)
	if err != nil {
		return err
	}
	fmt.Sscan(string(content), &b.perc)
	return nil
}

// Draw to lemonbar
func (b Battery) Draw() string {
	return b.prefix + fmt.Sprint(b.perc) + b.suffix
}
