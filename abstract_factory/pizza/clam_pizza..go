package pizza

import (
	pizzaingredinentfactory "github.com/projectred/head—first/abstract_factory/pizza_ingredinent_factory"
)

type ClamPizza struct {
	ingredinetFactory pizzaingredinentfactory.PizzaIngredientFactory
	DefaultPizza
}

func NewClamPizza(ingredinetFactory pizzaingredinentfactory.PizzaIngredientFactory) *ClamPizza {
	pizza := &ClamPizza{ingredinetFactory: ingredinetFactory}
	pizza.AbstractPrepare = pizza
	return pizza
}

func (cp *ClamPizza) Prepare() {
	println("Preparing " + cp.name)
	cp.dough = cp.ingredinetFactory.CreateDough()
	cp.sauce = cp.ingredinetFactory.CreateSauce()
	cp.cheese = cp.ingredinetFactory.CreateCheese()
	cp.clam = cp.ingredinetFactory.CreateClam()
}
