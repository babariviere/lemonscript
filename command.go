package main

import (
	"bytes"
	"os/exec"
	"strings"
)

var _ Updatable = (*Command)(nil)

// Command is a widget that executes system command
type Command struct {
	name   string
	args   []string
	result bytes.Buffer
	tick   uint
}

// NewCommand creates a new command widget
func NewCommand(tick uint, name string, args ...string) *Command {
	c := &Command{tick: tick, name: name, args: args}
	return c
}

// Update run commands
func (c *Command) Update() (bool, error) {
	var buf bytes.Buffer
	cmd := exec.Command(c.name, c.args...)
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		return false, err
	}
	if buf.String() != c.result.String() {
		c.result = buf
		return true, nil
	}
	return false, nil
}

// Tick refresh rate for command
func (c Command) Tick() uint {
	return c.tick
}

// Draw draws to lemonbar
func (c Command) Draw() string { return strings.TrimSpace(c.result.String()) }
