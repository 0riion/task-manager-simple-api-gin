package main

import (
	"github.com/0riion/task-manager-api-golang/pkg/tasks"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	port := viper.Get("PORT").(string)

	router := gin.Default()
	tasks.RegisterRoutes(router.Group("/api/v1"))

	router.Run(":" + port)
}
