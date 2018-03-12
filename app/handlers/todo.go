package handlers

import (
	"database/sql"
	"net/http"

	"github.com/mjelonek92/go-todo-app/app/models"
)

func ListTodos(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	respondJson(w, http.StatusOK, models.AllTodos(db))
}
