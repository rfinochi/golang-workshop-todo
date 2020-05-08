package mongo

import (
	"context"
	"os"

	"github.com/rfinochi/golang-workshop-todo/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	ID     int    `json:"id,omitempty bson:"id,omitempty" datastore:"id"`
	Title  string `json:"title,omitempty" bson:"title,omitempty" datastore:"title"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty" datastore:"isdone"`
}

var uri string

// MongoRepository godoc
type MongoRepository struct {
}

func init() {
	var ok bool

	if uri, ok = os.LookupEnv("TODO_MONGO_URI"); !ok {
		uri = "mongodb://localhost:27017"
	}
}

// CreateItem godoc
func (MongoRepository) CreateItem(newItem models.Item) (err error) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	_, err = collection.InsertOne(ctx, newItem)

	disconnect(ctx, client)

	return
}

// UpdateItem godoc
func (MongoRepository) UpdateItem(item models.Item) (err error) {
	update := bson.M{"$set": bson.M{"title": item.Title, "isdone": item.IsDone}}

	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	_, err = collection.UpdateOne(ctx, Item{ID: item.ID}, update)

	disconnect(ctx, client)

	return
}

// GetItems godoc
func (MongoRepository) GetItems() (items []models.Item, err error) {
	ctx, client := connnect()

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
func (MongoRepository) GetItem(id int) (item models.Item, err error) {
	ctx, client := connnect()

	options := options.Find()
	options.SetLimit(1)

	collection := client.Database("todo").Collection("items")
	cursor, err := collection.Find(ctx, Item{ID: id}, options)

	if err == nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			cursor.Decode(&item)
		}
	}

	disconnect(ctx, client)

	return
}

// DeleteItem godoc
func (MongoRepository) DeleteItem(id int) (err error) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	_, err = collection.DeleteMany(ctx, Item{ID: id})

	disconnect(ctx, client)

	return nil
}

func connnect() (context.Context, *mongo.Client) {
	ctx := context.Background()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return ctx, client
}

func disconnect(ctx context.Context, client *mongo.Client) {
	defer client.Disconnect(ctx)
}
