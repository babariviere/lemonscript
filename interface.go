package main

import (
	"net"
)

// NetInterface is a widget displaying net interface settings
type NetInterface struct {
	name string
	up   bool
	addr string
}

// NewInterface creates a new interface widget
func NewInterface(name string) Widget {
	return &NetInterface{name: name}
}

// Update updates interface info
func (i *NetInterface) Update() error {
	ni, err := net.InterfaceByName(i.name)
	if err != nil {
		return err
	}
	i.up = ni.Flags&net.FlagUp == net.FlagUp
	if i.up {
		addrs, err := ni.Addrs()
		if err != nil {
			return err
		}
		if len(addrs) == 0 {
			i.addr = "not connected"
		} else {
			i.addr = addrs[0].String()
		}
	}
	return nil
}

// Draw draws to lemonbar
func (i NetInterface) Draw() string {
	if !i.up {
		return i.name + ": down"
	}
	return i.name + ": " + i.addr
}
