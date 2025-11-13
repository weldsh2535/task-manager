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
