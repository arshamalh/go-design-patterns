package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	iglooBuilder := getBuilder("igloo")
	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	
	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}