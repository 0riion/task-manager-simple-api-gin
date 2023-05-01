// https://www.makeuseof.com/golang-crud-api-mongodb-gin/
package tasks

import (
	"net/http"

	"github.com/0riion/task-manager-api-golang/pkg/common/db"
	"github.com/0riion/task-manager-api-golang/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// CreateTask creates a new task in the database.
func CreateTask(c *gin.Context) {
	// Connect to the database.
	dbClient, dbErr := db.ConnectDB()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	defer db.CloseDB(dbClient)

	// Get the "tasks" collection from the database.
	tasksCollection := dbClient.Database("task_manager").Collection("tasks")

	// Bind the task data from the request body.
	var task models.Task
	if bindErr := c.BindJSON(&task); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	// Validate the task data.
	if task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}
	if task.User.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user email is required"})
		return
	}
	if task.StartDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start date is required"})
		return
	}
	if task.EndDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end date is required"})
		return
	}

	// Create a new task object.
	newTask := models.NewTask(task.Title, task.Details, task.StartDate, task.EndDate, task.User)

	// Insert the new task into the database.
	insertResult, insertErr := tasksCollection.InsertOne(c, newTask)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		return
	}

	// Retrieve the inserted task from the database.
	var insertedTask models.Task
	findErr := tasksCollection.FindOne(c, map[string]interface{}{"_id": insertResult.InsertedID}).Decode(&insertedTask)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
		return
	}

	// Return the new task to the client.
	c.JSON(http.StatusOK, gin.H{"message": "task created successfully", "task": insertedTask})
}
