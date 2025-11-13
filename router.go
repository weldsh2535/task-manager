// package main

// import (
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"gorm.io/gorm"
// )

// func NewRouter(db *gorm.DB) http.Handler {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("OK"))
// 	}).Methods("GET")

// 	r.HandleFunc("/tasks", GetTasks(db)).Methods("GET")
// 	r.HandleFunc("/tasks", CreateTask(db)).Methods("POST")
// 	r.HandleFunc("/tasks/{id:[0-9]+}", GetTask(db)).Methods("GET")
// 	r.HandleFunc("/tasks/{id:[0-9]+}", UpdateTask(db)).Methods("PUT")
// 	r.HandleFunc("/tasks/{id:[0-9]+}", DeleteTask(db)).Methods("DELETE")
// 	r.HandleFunc("/tasks/{id:[0-9]+}/toggle", ToggleTask(db)).Methods("POST")

// 	return r
// }

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weldsh2535/task-manager/handlers"
)

func SetupRouter() *gin.Engine {
	// Inject DB into handlers package
	handlers.SetDB(DB)

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/users", handlers.GetUsers)

	r.POST("/projects", handlers.CreateProject)
	r.GET("/projects", handlers.GetProjects)

	r.POST("/tasks", handlers.CreateTask)
	r.GET("/projects/:project_id/tasks", handlers.GetTasksByProject)

	return r
}
