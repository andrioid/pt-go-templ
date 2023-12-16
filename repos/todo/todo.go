package todo

type TodoRepo struct {
	Items []string
}

func (tr *TodoRepo) GetItems() []string {
	return tr.Items
}

func (tr *TodoRepo) AddItem(newItem string) {
	tr.Items = append(tr.Items, newItem)
}
