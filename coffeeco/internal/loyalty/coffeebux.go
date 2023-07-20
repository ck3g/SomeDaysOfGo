package loyalty

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	coffeeco "coffeeco/internal"
	"coffeeco/internal/store"
)

type CoffeeBux struct {
	ID                                    uuid.UUID
	Store                                 store.Store
	CoffeeLover                           coffeeco.CoffeeLover
	FreeDrinkAvailable                    int
	RemainingDrinkPurchasesUntilFreeDrink int
}

func (c *CoffeeBux) AddStamp() {
	if c.RemainingDrinkPurchasesUntilFreeDrink == 1 {
		c.RemainingDrinkPurchasesUntilFreeDrink = 10
		c.FreeDrinkAvailable += 1
	} else {
		c.RemainingDrinkPurchasesUntilFreeDrink--
	}
}

func (c *CoffeeBux) Pay(ctx context.Context, purchases []coffeeco.Product) error {
	lp := len(purchases)
	if lp == 0 {
		return errors.New("nothing to buy")
	}

	if c.FreeDrinkAvailable < lp {
		return fmt.Errorf("not enough coffeeBux to cover entire purchase. Have %d need %d", lp, c.FreeDrinkAvailable)
	}

	c.FreeDrinkAvailable = c.FreeDrinkAvailable - lp

	return nil
}
