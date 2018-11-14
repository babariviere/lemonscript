package main

import (
	"log"

	"github.com/mdirkse/i3ipc"
)

// I3 widget
type I3 struct {
	socket     *i3ipc.IPCSocket
	workspaces []i3ipc.Workspace
	normal     NestedWidget
	unfocused  NestedWidget
	urgent     NestedWidget
}

// NewI3 creates a new i3 widget
func NewI3(normal, unfocused, urgent NestedWidget) Widget {
	socket, err := i3ipc.GetIPCSocket()
	if err != nil {
		log.Fatal(err)
	}
	return &I3{socket: socket, normal: normal, unfocused: unfocused, urgent: urgent}
}

// Update updates i3 workspaces
func (i *I3) Update() error {
	if err := i.normal.Update(); err != nil {
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
		if workspace.Urgent {
			res += i.urgent.DrawWith(workspace.Name)
		} else if !workspace.Visible {
			res += i.unfocused.DrawWith(workspace.Name)
		} else {
			res += i.normal.DrawWith(workspace.Name)
		}
	}
	return res
}
