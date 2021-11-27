package pizza

type NYStyleCheesePizza struct {
	DefaultPizza
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{
		DefaultPizza{
			name:     "NY Style Sauce and Cheese Pizza",
			dough:    "Thin Crust Dough",
			sauce:    "Martinara Sauce",
			toppings: []string{"Grated Reggiano Cheese"},
		}}
}

type NYStyleBeffPizza struct {
	DefaultPizza
}

func NewNYStypeBeffPizza() *NYStyleBeffPizza {
	return &NYStyleBeffPizza{
		DefaultPizza{
			name:     "NY Style Beff Pizza",
			dough:    "Thin Crust Dough",
			sauce:    "Martinara Sauce",
			toppings: []string{"Greate Beff"},
		}}
}
