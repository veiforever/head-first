package pizzastore

import "github.com/projectred/head—first/abstract_factory/pizza"

type PizzaStore interface {
	OrderPizza(t string) pizza.Pizza
	AbstractCreatePizza
}

type AbstractCreatePizza interface {
	CreatePizza(t string) pizza.Pizza
}

type DefaultPizzaStore struct{ AbstractCreatePizza }

func (ps *DefaultPizzaStore) OrderPizza(t string) pizza.Pizza {
	pizza := ps.CreatePizza(t)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
