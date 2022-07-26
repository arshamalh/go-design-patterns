package main

import "fmt"

// The MakePizzaCommand is a struct which contains
// the number of pizzas to make, as well as the
// restaurant as its attributes
type MakePizzaCommand struct {
	count          int
	restaurant *Restaurant
}

func (c *MakePizzaCommand) execute() {
	// Reduce the total clean dishes of the restaurant
	// and print a message once done
	c.restaurant.CleanedDishes -= c.count
	fmt.Println("made", c.count, "pizzas")
}