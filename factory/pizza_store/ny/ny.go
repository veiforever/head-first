package ny

import (
	"github.com/projectred/head—first/factory/pizza"
	pizzastore "github.com/projectred/head—first/factory/pizza_store"
)

type NYPizzaStore struct {
	pizzastore.DefaultPizzaStore
}

func NewNYPizzaStore() *NYPizzaStore {
	pizzaStore := &NYPizzaStore{}
	pizzaStore.AbstractCreatePizza = pizzaStore
	return pizzaStore
}

func (ny *NYPizzaStore) CreatePizza(t string) pizza.Pizza {
	if t == "cheese" {
		return pizza.NewNYStyleCheesePizza()
	} else if t == "beff" {
		return pizza.NewNYStypeBeffPizza()
	}
	return &pizza.DefaultPizza{}
}
