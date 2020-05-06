package mongo

import (
	"context"
	"os"

	models "github.com/rfinochi/golang-workshop-todo/pkg/models"

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
func (MongoRepository) CreateItem(newItem models.Item) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	collection.InsertOne(ctx, newItem)

	disconnect(ctx, client)
}

// UpdateItem godoc
func (MongoRepository) UpdateItem(item models.Item) {
	update := bson.M{"$set": bson.M{"title": item.Title, "isdone": item.IsDone}}

	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	collection.UpdateOne(ctx, Item{ID: item.ID}, update)

	disconnect(ctx, client)
}

// GetItems godoc
func (MongoRepository) GetItems() (items []models.Item) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	cursor, _ := collection.Find(ctx, bson.M{})

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var oneItem models.Item
		cursor.Decode(&oneItem)
		items = append(items, oneItem)
	}

	disconnect(ctx, client)

	return
}

// GetItem godoc
func (MongoRepository) GetItem(id int) (item models.Item) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	collection.FindOne(ctx, Item{ID: id}).Decode(&item)

	disconnect(ctx, client)

	return
}

// DeleteItem godoc
func (MongoRepository) DeleteItem(id int) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	collection.DeleteMany(ctx, Item{ID: id})

	disconnect(ctx, client)
}

func connnect() (context.Context, *mongo.Client) {
	ctx := context.Background()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return ctx, client
}

func disconnect(ctx context.Context, client *mongo.Client) {
	defer client.Disconnect(ctx)
}
