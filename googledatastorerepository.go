package main

import (
	"context"

	"cloud.google.com/go/datastore"
)

type GoogleDatastoreRepository struct {
}

func (GoogleDatastoreRepository) CreateItem(newItem Item) {
	ctx := context.Background()

	client, _ := datastore.NewClient(ctx, "golang-workshop-todo")
	key := datastore.NameKey("TodoItem", string(newItem.ID), nil)
	client.Put(ctx, key, newItem);
}

func (GoogleDatastoreRepository) UpdateItem(item Item) {
	// update := bson.M{"$set": bson.M{"title": item.Title,"isdone": item.IsDone }}

	// ctx, client := connnect()

	// collection := client.Database("todo").Collection("items")
	// collection.UpdateOne(ctx, Item{ID: item.ID}, update)

	// disconnect(ctx,client)
}

func (GoogleDatastoreRepository) GetItems() (items []Item) {
	// ctx, client := connnect()

	// collection := client.Database("todo").Collection("items")
	// cursor, _ := collection.Find(ctx, bson.M{})

	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	var oneItem Item
	// 	cursor.Decode(&oneItem)
	// 	items = append(items, oneItem)
	// }

	// disconnect(ctx,client)

	return
}

func (GoogleDatastoreRepository) GetItem(id int) (item Item) {
	// ctx, client := connnect()

	// collection := client.Database("todo").Collection("items")
	// collection.FindOne(ctx, Item{ID: id}).Decode(&item)
	
	// disconnect(ctx,client)

	return
}

func (GoogleDatastoreRepository) DeleteItem(id int) {
	// ctx, client := connnect()

	// collection := client.Database("todo").Collection("items")
	// collection.DeleteMany(ctx, Item{ID: id})
	
	// disconnect(ctx,client)
}

// func connnect() (context.Context, *mongo.Client) {
// 	ctx := context.Background()
// 	client, _ := mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017"))

// 	return ctx, client
// }

// func disconnect(ctx context.Context, client *mongo.Client) {
// 	defer client.Disconnect(ctx)
// }
