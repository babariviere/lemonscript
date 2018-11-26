package main

import (
	"bytes"
	"os/exec"
	"strings"
)

// Command is a widget that executes system command
type Command struct {
	name   string
	args   []string
	result bytes.Buffer
}

// NewCommand creates a new command widget
func NewCommand(name string, args ...string) *Command {
	c := &Command{name: name, args: args}
	return c
}

// Update run commands
func (c *Command) Update() error {
	c.result.Reset()
	cmd := exec.Command(c.name, c.args...)
	cmd.Stdout = &c.result
	return cmd.Run()
}

// Draw draws to lemonbar
func (c Command) Draw() string { return strings.TrimSpace(c.result.String()) }
