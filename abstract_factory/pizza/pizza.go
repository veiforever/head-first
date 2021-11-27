package pizza

import "github.com/projectred/head—first/abstract_factory/ingredinent"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
	SetName(string)
}

type DefaultPizza struct {
	name      string
	dough     ingredinent.Dough
	sauce     ingredinent.Sauce
	veggies   []ingredinent.Veggie
	cheese    ingredinent.Cheese
	pepperoni ingredinent.Pepperoni
	clam      ingredinent.Clams
	AbstractPrepare
}

type AbstractPrepare interface {
	Prepare()
}

// func (d *DefaultPizza) Prepare() {
// 	println("Preparing", d.name)
// 	println("Tossing dough...")
// 	println("kAdding sauce...")
// 	println("Adding toppings:")
// 	for _, topping := range d.toppings {
// 		println("    ", topping)
// 	}
// }
func (d *DefaultPizza) Bake()               { println("Bake for 25 minutes at 350") }
func (d *DefaultPizza) Cut()                { println("Cutting the pizza into diagonal slices") }
func (d *DefaultPizza) Box()                { println("Place pizza in officaial PizzaSotore box") }
func (d *DefaultPizza) GetName() string     { return d.name }
func (d *DefaultPizza) SetName(name string) { d.name = name }
