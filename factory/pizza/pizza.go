package pizza

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type DefaultPizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func (d *DefaultPizza) Prepare() {
	println("Preparing", d.name)
	println("Tossing dough...")
	println("kAdding sauce...")
	println("Adding toppings:")
	for _, topping := range d.toppings {
		println("    ", topping)
	}
}
func (d *DefaultPizza) Bake()           { println("Bake for 25 minutes at 350") }
func (d *DefaultPizza) Cut()            { println("Cutting the pizza into diagonal slices") }
func (d *DefaultPizza) Box()            { println("Place pizza in officaial PizzaSotore box") }
func (d *DefaultPizza) GetName() string { return d.name }
