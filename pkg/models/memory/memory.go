package memory

import "github.com/rfinochi/golang-workshop-todo/pkg/models"

var items = []models.Item{}

// ItemRepository godoc
type ItemRepository struct {
}

// GetItems godoc
func (ItemRepository) GetItems() ([]models.Item, error) {
	return items, nil
}

// GetItem godoc
func (ItemRepository) GetItem(id int) (models.Item, error) {
	var result models.Item

	for _, item := range items {
		if item.ID == id {
			result = item
			break
		}
	}

	return result, nil
}

// CreateItem godoc
func (ItemRepository) CreateItem(newItem models.Item) error {
	items = append(items, newItem)

	return nil
}

// UpdateItem godoc
func (ItemRepository) UpdateItem(updatedItem models.Item) error {
	for i, item := range items {
		if item.ID == updatedItem.ID {
			item.Title = updatedItem.Title
			item.IsDone = updatedItem.IsDone
			items = append(items[:i], item)
		}
	}

	return nil
}

// DeleteItem godoc
func (ItemRepository) DeleteItem(id int) error {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}

	return nil
}
