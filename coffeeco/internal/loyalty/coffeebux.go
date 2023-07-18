package loyalty

import (
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
