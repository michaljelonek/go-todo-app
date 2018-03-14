package models

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

func (db *DB) AllTodos() ([]*Todo, error) {
	rows, err := db.Query("SELECT id, title, content FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]*Todo, 0)
	for rows.Next() {
		todo := &Todo{}
		rows.Scan(&todo.Id, &todo.Title, &todo.Content)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (db *DB) AddTodo() (*Todo, error) {
	title, content := "Title", "Content"
	// add sample todo for now
	result, err := db.Exec("INSERT INTO todo (title, content) VALUES ($1, $2)", title, content)
	if err != nil {
		return nil, err
	}
	id64, err := result.LastInsertId()
	id := int(id64)

	return &Todo{Id: id, Title: title, Content: content}, nil
}

func (db *DB) GetTodo(id int) (*Todo, error) {
	todo := Todo{}
	row := db.QueryRow("SELECT id, title, content FROM todo WHERE id=$1", id)
	if err := row.Scan(&todo.Id, &todo.Title, &todo.Content); err != nil {
		return nil, nil
	}

	return &todo, nil
}
