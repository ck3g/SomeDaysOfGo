package purchase

import (
	"context"
)

type Repository interface {
	Store(context.Context, Purchase) error
}
