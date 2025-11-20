// main.go
package main

import "fmt"

// Pizza interface ensures all pizza types implement getPrice.
type Pizza interface {
	getPrice() int
}

// Veggie is a base pizza.
type Veggie struct{}

// getPrice returns the base price of the Veggie pizza.
func (p *Veggie) getPrice() int {
	return 15
}

// CheeseTopping decorates a pizza by adding cheese.
type CheeseTopping struct {
	pizza Pizza
}

// getPrice adds cheese price to the wrapped pizza.
func (c *CheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 10 // Why: add extra cheese price.
}

// TomatoTopping decorates a pizza by adding tomato.
type TomatoTopping struct {
	pizza Pizza
}

// getPrice adds tomato price to the wrapped pizza.
func (t *TomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 5 // Why: add tomato topping price.
}

func main() {
	// Step 1: base pizza
	basePizza := &Veggie{}

	// Step 2: add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: basePizza,
	}

	// Step 3: add tomato topping on top of cheese
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	// Final price
	fmt.Printf(
		"Price of Veggie Pizza with tomato and cheese topping is %d\n",
		pizzaWithCheeseAndTomato.getPrice(),
	)
}
