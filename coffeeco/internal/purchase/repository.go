package purchase

import (
	"context"
	"fmt"
	"time"

	coffeeco "coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/store"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Store(context.Context, Purchase) error
}

type MongoRepository struct {
	purchases *mongo.Collection
}

type mongoPurchase struct {
	id                 uuid.UUID
	store              store.Store
	productsToPurchase []coffeeco.Product
	total              money.Money
	paymentMeans       payment.Means
	timeOfPurchase     time.Time
	cardToken          *string
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	purchases := client.Database("coffeeco").Collection("purchases")

	return &MongoRepository{
		purchases: purchases,
	}, nil
}

func (mr *MongoRepository) Store(ctx context.Context, purchase Purchase) error {
	mongoP := toMongoPurchase(purchase)
	_, err := mr.purchases.InsertOne(ctx, mongoP)
	if err != nil {
		return fmt.Errorf("failed to persist purchaseL %w", err)
	}
	return nil
}

func toMongoPurchase(p Purchase) mongoPurchase {
	return mongoPurchase{
		id:                 p.id,
		store:              p.Store,
		productsToPurchase: p.ProductsToPurchase,
		total:              p.total,
		paymentMeans:       p.PaymentMeans,
		timeOfPurchase:     p.timeOfPurchase,
		cardToken:          p.cardToken,
	}
}
