package models

import (
	"database/sql"
	"fmt"

	"github.com/mjelonek92/go-todo-app/config"
)

func InitDB(dbConfig *config.DBConfig) *sql.DB {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Name,
	)

	db, err := sql.Open(dbConfig.Dialect, dbURI)
	if err != nil {
		panic(err.Error())
	}

	return db
}
