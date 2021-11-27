package ny

import (
	"testing"
)

func TestPizzaStore(t *testing.T) {
	p := NewNYPizzaStore()
	p.OrderPizza("cheese")
}
