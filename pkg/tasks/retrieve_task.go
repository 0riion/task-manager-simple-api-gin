package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RetrieveTask(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "Task retrieved successfully", "id": id})
}
