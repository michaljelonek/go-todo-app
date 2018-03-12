package app

import (
	"database/sql"
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mjelonek92/go-todo-app/app/handlers"
	"github.com/mjelonek92/go-todo-app/app/models"
	"github.com/mjelonek92/go-todo-app/config"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Init(conf *config.Config) {
	db := models.InitDB(conf.DBConfig)

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	a.DB = db
	a.Router = mux.NewRouter()
	a.initRouters()
}

func (a *App) initRouters() {
	a.Router.HandleFunc("/", handlers.Healthcheck).Methods("Get")
	a.Router.HandleFunc("/health", handlers.Healthcheck).Methods("Get")
	a.Router.HandleFunc("/todo", a.listTodos).Methods("Get")
	a.Router.HandleFunc("/todo/add", a.addTodo).Methods("Post")
}

func (a *App) Run(port string) {
	defer a.DB.Close()
	loggedRouter := ghandlers.LoggingHandler(os.Stdout, a.Router)
	http.ListenAndServe(port, loggedRouter)
}

func (a *App) listTodos(w http.ResponseWriter, r *http.Request) {
	handlers.ListTodos(a.DB, w, r)
}

func (a *App) addTodo(w http.ResponseWriter, r *http.Request) {
	handlers.AddTodo(a.DB, w, r)
}
