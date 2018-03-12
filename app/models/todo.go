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

func AddTodo(db *sql.DB) *Todo {
	// add sample todo for now
	result, err := db.Exec("INSERT INTO todo (title, content) VALUES ($1, $2)", "Test Title", "Some test content")
	if err != nil {
		panic(err.Error())
	}
	id64, err := result.LastInsertId()
	id := int(id64)

	return GetTodo(db, id)
}

func GetTodo(db *sql.DB, id int) *Todo {
	todo := &Todo{Id: id}
	row := db.QueryRow("SELECT title, content FROM todo WHERE id=$1", id)
	row.Scan(&todo.Title, &todo.Content)
	return todo
}
