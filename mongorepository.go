package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoRepository godoc
type MongoRepository struct {
}

// CreateItem godoc
func (MongoRepository) CreateItem(newItem Item) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	collection.InsertOne(ctx, newItem)

	disconnect(ctx, client)
}

// UpdateItem godoc
func (MongoRepository) UpdateItem(item Item) {
	update := bson.M{"$set": bson.M{"title": item.Title, "isdone": item.IsDone}}

	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	collection.UpdateOne(ctx, Item{ID: item.ID}, update)

	disconnect(ctx, client)
}

// GetItems godoc
func (MongoRepository) GetItems() (items []Item) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	cursor, _ := collection.Find(ctx, bson.M{})

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var oneItem Item
		cursor.Decode(&oneItem)
		items = append(items, oneItem)
	}

	disconnect(ctx, client)

	return
}

// GetItem godoc
func (MongoRepository) GetItem(id int) (item Item) {
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
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	return ctx, client
}

func disconnect(ctx context.Context, client *mongo.Client) {
	defer client.Disconnect(ctx)
}
