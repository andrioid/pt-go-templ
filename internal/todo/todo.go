package todo

import "net/http"

var items []string = make([]string, 0)

func GetItems() []string {
	return items
}

func AddItem(newItem string) []string {
	items = append(items, newItem)
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
		list := AddItem(r.FormValue("add-item"))
		component := Items(list)
		component.Render(r.Context(), w)
		return
	}
}
