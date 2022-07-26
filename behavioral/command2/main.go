package main

import "fmt"

func main() {
	// initialize a new resaurant
	r := NewResteraunt(10)

	// create the list of tasks to be executed
	tasks := []Command{
		r.MakePizza(2),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	// create the cooks that will execute the tasks
	cooks := []*Cook{
		{},
		{},
	}

	// Assign tasks to cooks alternating between the existing
	// cooks.
	for i, task := range tasks {
		// Using the modulus of the current task index, we can
		// alternate between different cooks
		cook := cooks[i%len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	// Now that all the cooks have their commands, we can call
	// the `executeCommands` method that will have each cook
	// execute their respective commands
	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
	}
}