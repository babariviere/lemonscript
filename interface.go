package main

import (
	"net"
)

var _ Updatable = (*NetInterface)(nil)

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
func (i *NetInterface) Update() (bool, error) {
	ni, err := net.InterfaceByName(i.name)
	if err != nil {
		return false, err
	}
	var addr string
	up := ni.Flags&net.FlagUp == net.FlagUp
	if up {
		addrs, err := ni.Addrs()
		if err != nil {
			return false, err
		}
		if len(addrs) == 0 {
			addr = "not connected"
		} else {
			addr = addrs[0].String()
		}
	}
	if addr != i.addr || up != i.up {
		i.addr = addr
		i.up = up
		return true, nil
	}
	return false, nil
}

// Tick refresh rate for interface
func (i NetInterface) Tick() uint {
	return 60
}

// Draw draws to lemonbar
func (i NetInterface) Draw() string {
	if !i.up {
		return i.name + ": down"
	}
	return i.name + ": " + i.addr
}
