package models

import (
	"errors"
)

// Item godoc
type Item struct {
	ID     int    `json:"id,omitempty" bson:"id,omitempty" datastore:"id"`
	Title  string `json:"title,omitempty" bson:"title,omitempty" datastore:"title"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty" datastore:"isdone"`
}

var (
	// ErrNoRecord godoc
	ErrNoRecord = errors.New("models: no matching record found")
)

// ItemRepository godoc
type ItemRepository interface {
	CreateItem(Item) error
	UpdateItem(Item) error
	GetItems() ([]Item, error)
	GetItem(int) (Item, error)
	DeleteItem(int) error
}
