package tasks

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/tasks/create", CreateTask)
	r.GET("/tasks", ListTasks)
	r.GET("/tasks/:id", RetrieveTask)
	r.PUT("/tasks/:id", UpdateTask)
	r.DELETE("/tasks/:id", DeleteTask)
}
