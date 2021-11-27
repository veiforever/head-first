package pizza

import (
	pizzaingredinentfactory "github.com/projectred/head—first/abstract_factory/pizza_ingredinent_factory"
)

type CheesePizza struct {
	ingredinetFactory pizzaingredinentfactory.PizzaIngredientFactory
	DefaultPizza
}

func NewCheesePizza(ingredinetFactory pizzaingredinentfactory.PizzaIngredientFactory) *CheesePizza {
	pizza := &CheesePizza{ingredinetFactory: ingredinetFactory}
	pizza.AbstractPrepare = pizza
	return pizza
}

func (cp *CheesePizza) Prepare() {
	println("Preparing " + cp.name)
	cp.dough = cp.ingredinetFactory.CreateDough()
	cp.sauce = cp.ingredinetFactory.CreateSauce()
	cp.cheese = cp.ingredinetFactory.CreateCheese()
}
