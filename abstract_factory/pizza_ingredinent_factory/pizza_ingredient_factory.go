package pizzaingredinentfactory

import (
	"github.com/projectred/head—first/abstract_factory/ingredinent"
)

type PizzaIngredientFactory interface {
	CreateDough() ingredinent.Dough
	CreateSauce() ingredinent.Sauce
	CreateCheese() ingredinent.Cheese
	CreateVeggies() []ingredinent.Veggie
	CreatePepperoni() ingredinent.Pepperoni
	CreateClam() ingredinent.Clams
}
