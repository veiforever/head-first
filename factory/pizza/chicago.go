package pizza

type ChicagoStyleCheesePizza struct {
	DefaultPizza
}

func (p *ChicagoStyleCheesePizza) Cut() { println("Cutting; the pizza into square lices") }

func NewChicagoStyleCheesePizza() Pizza {
	return &ChicagoStyleCheesePizza{
		DefaultPizza{
			name:     "Chicago Style Deep Dish Cheese Pizza",
			dough:    "Extra Thick Crust Dough",
			sauce:    "Plum Tomato Sauce",
			toppings: []string{"Shredded Mozzarella Cheese"},
		}}
}
