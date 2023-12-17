package todo

import "net/http"

type Todo string

var items []Todo = make([]Todo, 0)

func GetItems() []Todo {
	return items
}

func AddItem(newItem string) []Todo {
	items = append(items, Todo(newItem))
	return items
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method == http.MethodGet && r.URL.Path == "/" {
		component := TodoIndex()
		component.Render(r.Context(), w)
		return
	}
	if r.Method == http.MethodPost {
		AddItem(r.FormValue("add-item"))
		item := Todo(r.FormValue("add-item"))
		component := TodoItem(item)
		component.Render(r.Context(), w)

		return
	}
}
