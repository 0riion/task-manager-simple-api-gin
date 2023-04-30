package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "List of tasks"})
}
