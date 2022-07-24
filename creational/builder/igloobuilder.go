package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (ib *IglooBuilder) setWindowType() {
	ib.windowType = "Snow window"
}

func (ib *IglooBuilder) setDoorType() {
	ib.doorType = "Snow door"
}

func (ib *IglooBuilder) setNumFloor() {
	ib.floor = 1
}

func (ib *IglooBuilder) getHouse() House {
	return House{
		windowType: ib.windowType,
		doorType:   ib.doorType,
		floor:      ib.floor,
	}
}
