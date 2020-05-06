package memory

import "github.com/rfinochi/golang-workshop-todo/pkg/models"

var items = []models.Item{}

type Item struct {
	ID     int    `json:"id,omitempty bson:"id,omitempty" datastore:"id"`
	Title  string `json:"title,omitempty" bson:"title,omitempty" datastore:"title"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty" datastore:"isdone"`
}

// Memory godoc
type MemoryRepository struct {
}

// CreateItem godoc
func (MemoryRepository) CreateItem(newItem models.Item) {
	items = append(items, newItem)
}

// UpdateItem godoc
func (MemoryRepository) UpdateItem(updatedItem models.Item) {
	for i, item := range items {
		if item.ID == updatedItem.ID {
			item.Title = updatedItem.Title
			item.IsDone = updatedItem.IsDone
			items = append(items[:i], item)
		}
	}
}

// GetItems godoc
func (MemoryRepository) GetItems() []models.Item {
	return items
}

// GetItem godoc
func (MemoryRepository) GetItem(id int) models.Item {
	var result models.Item

	for _, item := range items {
		if item.ID == id {
			result = item
			break
		}
	}

	return result
}

// DeleteItem godoc
func (MemoryRepository) DeleteItem(id int) {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
}
