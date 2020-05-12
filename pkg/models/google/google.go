package google

import (
	"context"

	"github.com/rfinochi/golang-workshop-todo/pkg/models"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

const entityName string = "todoitem"

// ItemRepository godoc
type ItemRepository struct {
}

// CreateItem godoc
func (ItemRepository) CreateItem(newItem models.Item) (err error) {
	ctx, client, err := connnectToDatastore()
	if err != nil {
		return
	}

	key := datastore.IDKey(entityName, int64(newItem.ID), nil)
	_, err = client.Put(ctx, key, &newItem)
	if err != nil {
		return
	}

	err = client.Close()

	return
}

// UpdateItem godoc
func (ItemRepository) UpdateItem(item models.Item) (err error) {
	ctx, client, err := connnectToDatastore()
	if err != nil {
		return
	}

	key := datastore.IDKey(entityName, int64(item.ID), nil)
	_, err = client.Put(ctx, key, &item)
	if err != nil {
		return
	}

	err = client.Close()

	return
}

// GetItems godoc
func (ItemRepository) GetItems() (items []models.Item, err error) {
	ctx, client, err := connnectToDatastore()
	if err != nil {
		return
	}

	query := datastore.NewQuery("todoitem").Order("ID")
	it := client.Run(ctx, query)
	for {
		var item models.Item
		if _, err = it.Next(&item); err == iterator.Done {
			err = nil
			break
		} else if err != nil {
			return
		}
		items = append(items, item)
	}

	err = client.Close()

	return
}

// GetItem godoc
func (ItemRepository) GetItem(id int) (item models.Item, err error) {
	ctx, client, err := connnectToDatastore()
	if err != nil {
		return
	}

	key := datastore.IDKey(entityName, int64(id), nil)
	err = client.Get(ctx, key, &item)
	if err == datastore.ErrNoSuchEntity {
		err = nil
	} else if err != nil {
		return
	}

	err = client.Close()

	return
}

// DeleteItem godoc
func (ItemRepository) DeleteItem(id int) (err error) {
	ctx, client, err := connnectToDatastore()
	if err != nil {
		return
	}

	key := datastore.IDKey(entityName, int64(id), nil)
	err = client.Delete(ctx, key)
	if err != nil {
		return
	}

	err = client.Close()

	return
}

func connnectToDatastore() (ctx context.Context, client *datastore.Client, err error) {
	ctx = context.Background()

	client, err = datastore.NewClient(ctx, "golang-workshop-todo")
	if err != nil {
		return
	}

	err = client.Close()

	return
}
