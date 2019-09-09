package main

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

const entityName string = "todoitem"

// GoogleDatastoreRepository godoc
type GoogleDatastoreRepository struct {
}

// CreateItem godoc
func (GoogleDatastoreRepository) CreateItem(newItem Item) {
	ctx, client := connnectToDatastore()

	key := datastore.IDKey(entityName, int64(newItem.ID), nil)
	_, err := client.Put(ctx, key, &newItem)
	if err != nil {
		log.Printf("CreateItem Error: '%s'", err.Error())
	}
}

// UpdateItem godoc
func (GoogleDatastoreRepository) UpdateItem(item Item) {
	ctx, client := connnectToDatastore()

	key := datastore.IDKey(entityName, int64(item.ID), nil)
	_, err := client.Put(ctx, key, &item)
	if err != nil {
		log.Printf("UpdateItem Error: '%s'", err.Error())
	}
}

// GetItems godoc
func (GoogleDatastoreRepository) GetItems() (items []Item) {
	ctx, client := connnectToDatastore()

	query := datastore.NewQuery("todoitem").Order("ID")
	it := client.Run(ctx, query)
	for {
		var item Item
		if _, err := it.Next(&item); err == iterator.Done {
			break
		} else if err != nil {
			log.Printf("GetItems Error: '%s'", err.Error())
		}
		items = append(items, item)
	}

	return
}

// GetItem godoc
func (GoogleDatastoreRepository) GetItem(id int) (item Item) {
	ctx, client := connnectToDatastore()

	key := datastore.IDKey(entityName, int64(id), nil)
	err := client.Get(ctx, key, &item)
	if err != nil {
		log.Printf("GetItem Error: '%s'", err.Error())
	}

	return
}

// DeleteItem godoc
func (GoogleDatastoreRepository) DeleteItem(id int) {
	ctx, client := connnectToDatastore()

	key := datastore.IDKey(entityName, int64(id), nil)
	err := client.Delete(ctx, key)
	if err != nil {
		log.Printf("DeleteItem Error: '%s'", err.Error())
	}
}

func connnectToDatastore() (context.Context, *datastore.Client) {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "golang-workshop-todo")
	if err != nil {
		log.Fatal(err.Error())
	}

	return ctx, client
}
