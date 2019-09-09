package main

var items = []Item{}

// Memory godoc
type Memory struct {
}

// CreateItem godoc
func (Memory) CreateItem(newItem Item) {
	items = append(items, newItem)
}

// UpdateItem godoc
func (Memory) UpdateItem(updatedItem Item) {
	for i, item := range items {
		if item.ID == updatedItem.ID {
			item.Title = updatedItem.Title
			item.IsDone = updatedItem.IsDone
			items = append(items[:i], item)
		}
	}
}

// GetItems godoc
func (Memory) GetItems() []Item {
	return items
}

// GetItem godoc
func (Memory) GetItem(id int) Item {
	var result Item

	for _, item := range items {
		if item.ID == id {
			result = item
			break
		}
	}

	return result
}

// DeleteItem godoc
func (Memory) DeleteItem(id int) {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
}
