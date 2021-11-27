package ny

import (
	"github.com/projectred/head—first/abstract_factory/ingredinent"
)

type NYPizzaIngredientFactory struct{}

func (f *NYPizzaIngredientFactory) CreateDough() ingredinent.Dough {
	return ingredinent.ThinCrustDough{}
}
func (f *NYPizzaIngredientFactory) CreateSauce() ingredinent.Sauce {
	return ingredinent.MarinaraSauce{}
}

func (f *NYPizzaIngredientFactory) CreateCheese() ingredinent.Cheese {
	return ingredinent.ReggianoCheese{}
}

func (f *NYPizzaIngredientFactory) CreateVeggies() []ingredinent.Veggie {
	return []ingredinent.Veggie{ingredinent.Garlic{}, ingredinent.Onion{}, ingredinent.Mushroom{}, ingredinent.RedPepper{}}
}

func (f *NYPizzaIngredientFactory) CreatePepperoni() ingredinent.Pepperoni {
	return ingredinent.SlicedPepperoni{}
}

func (f *NYPizzaIngredientFactory) CreateClam() ingredinent.Clams {
	return ingredinent.FreshClams{}
}
