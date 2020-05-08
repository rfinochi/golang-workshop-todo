package models

// Item godoc
type Item struct {
	ID     int    `json:"id,omitempty bson:"id,omitempty" datastore:"id"`
	Title  string `json:"title,omitempty" bson:"title,omitempty" datastore:"title"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty" datastore:"isdone"`
}

// TodoRepository godoc
type TodoRepository interface {
	CreateItem(Item) error
	UpdateItem(Item) error
	GetItems() ([]Item, error)
	GetItem(int) (Item, error)
	DeleteItem(int) error
}
