package handlers

import (
	"net/http"

	"github.com/weldsh2535/task-manager/models"

	"github.com/gin-gonic/gin"
)

type TaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	ProjectID   uint   `json:"project_id"`
	Completed   bool   `json:"completed"`
}

// Create Task
func CreateTask(c *gin.Context) {
	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		UserID:      input.UserID,
		ProjectID:   input.ProjectID,
		Completed:   input.Completed,
	}
	DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

// Get Tasks by Project
func GetTasksByProject(c *gin.Context) {
	var tasks []models.Task
	projectID := c.Param("project_id")
	DB.Where("project_id = ?", projectID).Preload("User").Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}
