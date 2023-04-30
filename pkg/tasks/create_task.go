package tasks

import (
	"net/http"

	"github.com/0riion/task-manager-api-golang/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	err := c.BindJSON(&task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// title validation
	if task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	// email validation
	if task.User.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": " task.User.Email is required"})
		return
	}

	// start_date validation
	if task.StartDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date is required"})
		return
	}

	// end_date validation
	if task.EndDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end_date is required"})
		return
	}

	// create new task
	newTask := models.NewTask(task.Title, task.Details, task.StartDate, task.EndDate, task.User)

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully", "task": newTask})
}
