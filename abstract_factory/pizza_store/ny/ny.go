package ny

import (
	"github.com/projectred/head—first/abstract_factory/pizza"
	pizzaingredinentfactory "github.com/projectred/head—first/abstract_factory/pizza_ingredinent_factory"
	"github.com/projectred/head—first/abstract_factory/pizza_ingredinent_factory/ny"
	pizzastore "github.com/projectred/head—first/abstract_factory/pizza_store"
)

type NYPizzaStore struct {
	pizzastore.DefaultPizzaStore
	ingredinetFactory pizzaingredinentfactory.PizzaIngredientFactory
}

func NewNYPizzaStore() *NYPizzaStore {
	pizzaStore := &NYPizzaStore{ingredinetFactory: &ny.NYPizzaIngredientFactory{}}
	pizzaStore.AbstractCreatePizza = pizzaStore
	return pizzaStore
}

func (ny *NYPizzaStore) CreatePizza(t string) pizza.Pizza {
	var p pizza.Pizza
	if t == "cheese" {
		p = pizza.NewCheesePizza(ny.ingredinetFactory)
		p.SetName("New York Style Cheese Pizza")
	} else if t == "claim" {
		p = pizza.NewClamPizza(ny.ingredinetFactory)
		p.SetName("New York Style Clam Pizza")
	}
	return p
}
