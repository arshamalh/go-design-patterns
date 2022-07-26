package main

import "fmt"

// The MakeSaladCommand is similar to the MakePizza command
type MakeSaladCommand struct {
	count          int
	restaurant *Restaurant
}

func (c *MakeSaladCommand) execute() {
	c.restaurant.CleanedDishes -= c.count
	fmt.Println("made", c.count, "salads")
}
