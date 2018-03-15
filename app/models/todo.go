package models

type Todo struct {
	Id       int
	Title    string `json:"title"`
	Content  string `json:"content"`
	Finished bool   `json:"finished"`
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

func (db *DB) AddTodo(title, content string) (*Todo, error) {
	stmt, err := db.Prepare("INSERT INTO todo (title, content) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var id int
	if err = stmt.QueryRow(title, content).Scan(&id); err != nil {
		return nil, err
	}

	return db.GetTodo(id)
}

func (db *DB) GetTodo(id int) (*Todo, error) {
	todo := Todo{}
	row := db.QueryRow("SELECT id, title, content FROM todo WHERE id=$1", id)
	if err := row.Scan(&todo.Id, &todo.Title, &todo.Content); err != nil {
		return nil, nil
	}

	return &todo, nil
}
