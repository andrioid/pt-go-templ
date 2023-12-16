package repos

import "app/repos/todo"

func CreateRepositories() struct {
	Todo todo.TodoRepo
} {
	return struct{ Todo todo.TodoRepo }{
		Todo: todo.TodoRepo{
			Items: []string{"initial item"},
		},
	}
}
