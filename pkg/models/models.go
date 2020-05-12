package models

import "errors"

// Item godoc
type Item struct {
	ID     int    `json:"id,omitempty" bson:"id,omitempty" datastore:"id"`
	Title  string `json:"title,omitempty" bson:"title,omitempty" datastore:"title"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty" datastore:"isdone"`
}

var (
	// ErrNoRecord godoc
	ErrNoRecord = errors.New("no matching record found")
)

// ItemRepository godoc
type ItemRepository interface {
	CreateItem(Item) error
	UpdateItem(Item) error
	GetItems() ([]Item, error)
	GetItem(int) (Item, error)
	DeleteItem(int) error
}

// ItemModel godoc
type ItemModel struct {
	Repository ItemRepository
}

// CreateItem godoc
func (model ItemModel) CreateItem(newItem Item) error {
	return model.Repository.CreateItem(newItem)
}

// UpdateItem godoc
func (model ItemModel) UpdateItem(updatedItem Item) error {
	return model.Repository.UpdateItem(updatedItem)
}

// GetItems godoc
func (model ItemModel) GetItems() ([]Item, error) {
	return model.Repository.GetItems()
}

// GetItem godoc
func (model ItemModel) GetItem(id int) (i Item, e error) {
	i, e = model.Repository.GetItem(id)

	if i == (Item{}) && e == nil {
		e = ErrNoRecord
	}

	return
}

// DeleteItem godoc
func (model ItemModel) DeleteItem(id int) error {
	return model.Repository.DeleteItem(id)
}
