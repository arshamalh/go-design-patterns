package main

// MakePizzaCommand,
// MakeSaladCommand,
// CleanDishesCommand
// all implement the Command interface with their execute method.

// The restaurant contains the total dishes and the total cleaned dishes.
type Restaurant struct {
	TotalDishes int
	CleanedDishes int
}

// `NewRestaurant` constructs a new restaurant instance with 10 dishes,
// all of them being clean
func NewResteraunt(number_of_dishes int) *Restaurant {
	return &Restaurant{
		TotalDishes:   number_of_dishes,
		CleanedDishes: number_of_dishes,
	}
}


func (r *Restaurant) MakePizza(count int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		count:          count,
	}
}

func (r *Restaurant) MakeSalad(count int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		count:          count,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}