package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RetrieveTask(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Task retrieved successfully"})
}
