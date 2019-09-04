package main

var items = []Item{}

type Memory struct {
}

func (Memory) CreateItem(newItem Item) {
	items = append(items, newItem)
}

func (Memory) UpdateItem(updatedItem Item) {
	for i, item := range items {
		if item.ID == updatedItem.ID {
			item.Title = updatedItem.Title
			item.IsDone = updatedItem.IsDone
			items = append(items[:i], item)
		}
	}
}

func (Memory) GetItems() []Item {
	return items
}

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

func (Memory) DeleteItem(id int) {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
}
