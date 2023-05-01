package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "id": id})
}
