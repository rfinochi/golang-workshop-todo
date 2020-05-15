package mongo

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rfinochi/golang-workshop-todo/pkg/common"
	"github.com/rfinochi/golang-workshop-todo/pkg/models"
)

// ItemRepository godoc
type ItemRepository struct {
}

// GetItems godoc
func (ItemRepository) GetItems() (items []models.Item, err error) {
	ctx, client, err := connnect()
	if err != nil {
		return
	}

	collection := client.Database("todo").Collection("items")
	cursor, err := collection.Find(ctx, bson.M{})

	if err == nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var oneItem models.Item
			cursor.Decode(&oneItem)
			items = append(items, oneItem)
		}
	}

	disconnect(ctx, client)

	return
}

// GetItem godoc
func (ItemRepository) GetItem(id int) (item models.Item, err error) {
	ctx, client, err := connnect()
	if err != nil {
		return
	}

	options := options.Find()
	options.SetLimit(1)

	collection := client.Database("todo").Collection("items")
	cursor, err := collection.Find(ctx, models.Item{ID: id}, options)

	if err == nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			cursor.Decode(&item)
		}
	}

	disconnect(ctx, client)

	return
}

// CreateItem godoc
func (ItemRepository) CreateItem(newItem models.Item) (err error) {
	ctx, client, err := connnect()
	if err != nil {
		return
	}

	collection := client.Database("todo").Collection("items")
	_, err = collection.InsertOne(ctx, newItem)

	disconnect(ctx, client)

	return
}

// UpdateItem godoc
func (ItemRepository) UpdateItem(item models.Item) (err error) {
	update := bson.M{"$set": bson.M{"title": item.Title, "isdone": item.IsDone}}

	ctx, client, err := connnect()
	if err != nil {
		return
	}

	collection := client.Database("todo").Collection("items")
	_, err = collection.UpdateOne(ctx, models.Item{ID: item.ID}, update)

	disconnect(ctx, client)

	return
}

// DeleteItem godoc
func (ItemRepository) DeleteItem(id int) (err error) {
	ctx, client, err := connnect()
	if err != nil {
		return
	}

	collection := client.Database("todo").Collection("items")
	_, err = collection.DeleteMany(ctx, models.Item{ID: id})

	disconnect(ctx, client)

	return
}

func connnect() (ctx context.Context, client *mongo.Client, err error) {
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(getURI()))
	if err != nil {
		return
	}

	return
}

func disconnect(ctx context.Context, client *mongo.Client) {
	defer client.Disconnect(ctx)
}

func getURI() (uri string) {
	var ok bool

	if uri, ok = os.LookupEnv(common.RepositoryMongoURIEnvVarName); !ok {
		uri = common.RepositoryMongoURIDefault
	}

	return
}
