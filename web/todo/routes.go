package todo

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoId}", GetATodo)
	router.Delete("/{todoId}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetTodos)

	return router
}

func GetATodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")
	todos := Todo{
		Slug:  todoId,
		Title: "hello world",
		Body:  "Hello world from planet earth",
	}

	render.JSON(w, r, todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "deleted TODO successfully"
	render.JSON(w, r, response)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Create TODO successfully"
	render.JSON(w, r, response)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{
			Slug:  "slug",
			Title: "hello world",
			Body:  "Hello worl d from planet earth",
		},
	}

	render.JSON(w, r, todos)
}
