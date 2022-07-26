package main

type OnCommand struct {
	device Device
}

func (oc *OnCommand) execute() {
	oc.device.on()
}