package handlers

import (
	"net/http"

	"github.com/weldsh2535/task-manager/models"

	"github.com/gin-gonic/gin"
)

type ProjectInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create Project
func CreateProject(c *gin.Context) {
	var input ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project := models.Project{Name: input.Name, Description: input.Description}
	DB.Create(&project)
	c.JSON(http.StatusOK, project)
}

// Get All Projects
func GetProjects(c *gin.Context) {
	var projects []models.Project
	DB.Preload("Tasks").Preload("Users").Find(&projects)
	c.JSON(http.StatusOK, projects)
}
