package main

import (
	"fmt"
	"log"

	"github.com/mdirkse/i3ipc"
)

// I3 widget
type I3 struct {
	socket     *i3ipc.IPCSocket
	workspaces []i3ipc.Workspace
	focused    NestedWidget
	unfocused  NestedWidget
	urgent     NestedWidget
}

// NewI3 creates a new i3 widget
func NewI3(focused, unfocused, urgent NestedWidget) Widget {
	socket, err := i3ipc.GetIPCSocket()
	if err != nil {
		log.Fatal(err)
	}
	return &I3{socket: socket, focused: focused, unfocused: unfocused, urgent: urgent}
}

// Update updates i3 workspaces
func (i *I3) Update() error {
	if err := i.focused.Update(); err != nil {
		return err
	}
	if err := i.unfocused.Update(); err != nil {
		return err
	}
	if err := i.urgent.Update(); err != nil {
		return err
	}
	var err error
	i.workspaces, err = i.socket.GetWorkspaces()
	return err
}

// Draw draws to lemonbar
func (i I3) Draw() string {
	var res string
	for _, workspace := range i.workspaces {
		var block string
		clickable := NewClickable(fmt.Sprintf("i3-msg workspace %d", workspace.Num))
		if workspace.Visible {
			block = i.focused.DrawWith(workspace.Name)
		} else {
			block = i.unfocused.DrawWith(workspace.Name)
		}
		if workspace.Urgent {
			block = i.urgent.DrawWith(block)
		}
		res += clickable.DrawWith(block)
	}
	return res
}
