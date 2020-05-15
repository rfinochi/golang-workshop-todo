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
	ErrNoRecord = errors.New("No matching record found")
	// ErrRecordExist godoc
	ErrRecordExist = errors.New("Record already exists")
)

// ItemRepository godoc
type ItemRepository interface {
	GetItems() ([]Item, error)
	GetItem(int) (Item, error)
	CreateItem(Item) error
	UpdateItem(Item) error
	DeleteItem(int) error
}

// ItemModel godoc
type ItemModel struct {
	Repository ItemRepository
}

// GetItems godoc
func (model ItemModel) GetItems() (i []Item, e error) {
	i, e = model.Repository.GetItems()

	if i == nil && e == nil {
		i = []Item{}
	}

	return
}

// GetItem godoc
func (model ItemModel) GetItem(id int) (i Item, e error) {
	i, e = model.Repository.GetItem(id)

	if i == (Item{}) && e == nil {
		e = ErrNoRecord
	}

	return
}

// CreateItem godoc
func (model ItemModel) CreateItem(newItem Item) error {
	i, _ := model.Repository.GetItem(newItem.ID)

	if i.ID == newItem.ID {
		return ErrRecordExist
	}

	return model.Repository.CreateItem(newItem)
}

// UpdateItem godoc
func (model ItemModel) UpdateItem(updatedItem Item) error {
	i, e := model.Repository.GetItem(updatedItem.ID)

	if i == (Item{}) && e == nil {
		return ErrNoRecord
	}

	return model.Repository.UpdateItem(updatedItem)
}

// DeleteItem godoc
func (model ItemModel) DeleteItem(id int) error {
	i, e := model.Repository.GetItem(id)

	if i == (Item{}) && e == nil {
		return ErrNoRecord
	}

	return model.Repository.DeleteItem(id)
}
