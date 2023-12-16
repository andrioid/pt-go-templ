package todo

var items []string = make([]string, 0)

func GetItems() []string {
	return items
}

func AddItem(newItem string) []string {
	items = append(items, newItem)
	return items
}
