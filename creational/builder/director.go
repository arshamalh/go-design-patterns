package main

type Director struct {
	builder IBuilder
}

func newDirector(builder IBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) setBuilder(builder IBuilder) {
	d.builder = builder
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setNumFloor()
	d.builder.setWindowType()
	return d.builder.getHouse()
}