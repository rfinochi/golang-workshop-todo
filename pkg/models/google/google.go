package google

import (
	"context"
	"log"

	"github.com/rfinochi/golang-workshop-todo/pkg/models"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

const entityName string = "todoitem"

// GoogleRepository godoc
type GoogleRepository struct {
}

// CreateItem godoc
func (GoogleRepository) CreateItem(newItem models.Item) (err error) {
	ctx, client := connnectToDatastore()

	key := datastore.IDKey(entityName, int64(newItem.ID), nil)
	_, err = client.Put(ctx, key, &newItem)

	return err
}

// UpdateItem godoc
func (GoogleRepository) UpdateItem(item models.Item) (err error) {
	ctx, client := connnectToDatastore()

	key := datastore.IDKey(entityName, int64(item.ID), nil)
	_, err = client.Put(ctx, key, &item)

	return
}

// GetItems godoc
func (GoogleRepository) GetItems() (items []models.Item, err error) {
	ctx, client := connnectToDatastore()

	query := datastore.NewQuery("todoitem").Order("ID")
	it := client.Run(ctx, query)
	for {
		var item models.Item
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
func (GoogleRepository) GetItem(id int) (item models.Item, err error) {
	ctx, client := connnectToDatastore()

	key := datastore.IDKey(entityName, int64(id), nil)
	err = client.Get(ctx, key, &item)
	if err != nil {
		log.Printf("GetItem Error: '%s'", err.Error())
	}

	return
}

// DeleteItem godoc
func (GoogleRepository) DeleteItem(id int) error {
	ctx, client := connnectToDatastore()

	key := datastore.IDKey(entityName, int64(id), nil)
	return client.Delete(ctx, key)
}

func connnectToDatastore() (context.Context, *datastore.Client) {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "golang-workshop-todo")
	if err != nil {
		log.Fatal(err.Error())
	}

	return ctx, client
}
