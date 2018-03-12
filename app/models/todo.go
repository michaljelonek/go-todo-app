package models

import "database/sql"

type Todo struct {
	Id       int
	Title    string
	Content  string
	Finished bool
}

func (t *Todo) MarkFinished() {
	t.Finished = true
}

func (t *Todo) MarkUnfished() {
	t.Finished = false
}

func AllTodos(db *sql.DB) []*Todo {
	var todos []*Todo

	rows, err := db.Query("SELECT id, title, content FROM todo")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		todo := &Todo{}
		rows.Scan(&todo.Id, &todo.Title, &todo.Content)
		todos = append(todos, todo)
	}

	return todos
}
