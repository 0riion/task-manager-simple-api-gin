package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully", "id": id})
}
