package main

import (
	"fmt"
	"log"
	"sync"

	"go.i3wm.org/i3"
)

// I3 widget
type I3 struct {
	mutex      *sync.Mutex
	workspaces []i3.Workspace
	focused    NestedWidget
	unfocused  NestedWidget
	urgent     NestedWidget
}

// updates structure from workspace event
func registerI3(widget *I3) {
	eventReceiver := i3.Subscribe(i3.WorkspaceEventType)
	for eventReceiver.Next() {
		widget.mutex.Lock()
		_ = eventReceiver.Event().(*i3.WorkspaceEvent)
		widget.workspaces, _ = i3.GetWorkspaces()
		widget.mutex.Unlock()
	}
}

// NewI3 creates a new i3 widget
func NewI3(focused, unfocused, urgent NestedWidget) *I3 {
	workspaces, err := i3.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
	}
	widget := &I3{mutex: &sync.Mutex{}, workspaces: workspaces, focused: focused, unfocused: unfocused, urgent: urgent}
	go registerI3(widget)
	return widget
}

// Update updates i3 workspaces
func (i *I3) Update() error {
	return nil
}

// Draw draws to lemonbar
func (i I3) Draw() string {
	var res string
	i.mutex.Lock()
	defer i.mutex.Unlock()
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
