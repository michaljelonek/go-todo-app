package app

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mjelonek92/go-todo-app/app/models"
	"github.com/mjelonek92/go-todo-app/app/utils"
	"github.com/mjelonek92/go-todo-app/config"
)

type App struct {
	router *mux.Router
	db     models.Datastore
}

func (app *App) Start(conf *config.Config) {
	db, err := models.InitDB(conf.DBConfig)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	app.db = db
	app.router = mux.NewRouter()
	app.initRouters()
	app.run(":8080")
}

func (app *App) initRouters() {
	app.router.HandleFunc("/", app.status).Methods("Get")
	app.router.HandleFunc("/todo", app.listTodos).Methods("Get")
	app.router.HandleFunc("/todo/{id:[0-9]+}", app.getTodo).Methods("Get")
	app.router.HandleFunc("/todo/create", app.addTodo).Methods("Post")
}

func (app *App) run(addr string) {
	loggedRouter := handlers.LoggingHandler(os.Stdout, app.router)
	http.ListenAndServe(addr, loggedRouter)
}

func (app *App) listTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.db.AllTodos()
	if err != nil {
		utils.ServerError(w)
		return
	}

	utils.RespondJson(w, http.StatusOK, todos)
}

func (app *App) addTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := app.db.AddTodo()
	if err != nil {
		utils.ServerError(w)
		return
	}

	utils.RespondJson(w, http.StatusCreated, todo)
}

func (app *App) getTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.BadRequest(w, "ID must be an int")
	}

	todo, err := app.db.GetTodo(id)
	if err != nil {
		utils.ServerError(w)
		return
	}

	utils.RespondJson(w, http.StatusOK, todo)
}

func (app *App) status(w http.ResponseWriter, r *http.Request) {
	utils.RespondJson(w, http.StatusOK, "API is up and working!")
}
