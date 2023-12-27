package todo

import (
	"app/internal/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoModel struct {
	id       int64
	name     string
	finished bool
}

func GetItems() []TodoModel {
	rows, err := db.DB.Query("SELECT id, name, finished FROM todo")
	if err != nil {
		panic(err)
	}
	var items []TodoModel

	for rows.Next() {
		var item TodoModel
		rows.Scan(&item.id, &item.name, &item.finished)
		items = append(items, item)
	}
	return items
}

func AddItem(newItem string) (TodoModel, error) {
	res, err := db.DB.Exec("INSERT INTO todo (name) VALUES (?)", newItem)
	if err != nil {
		return TodoModel{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return TodoModel{}, err
	}
	return TodoModel{
		id:   id,
		name: newItem,
	}, nil
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("todo handler called", r.URL)
	// if r.URL.Path != "/" {
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// 	return
	// }
	if r.Method == http.MethodGet {
		component := TodoIndex()
		component.Render(r.Context(), w)
		return
	}
	if r.Method == http.MethodPost {
		item, err := AddItem(r.FormValue("add-item"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		component := TodoItem(item)
		component.Render(r.Context(), w)

		return
	}
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", RootHandler)
}
