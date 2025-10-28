package router

import (
	"gin-todo-list/internal/router/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	group := r.Group("/api")
	
	group.GET("/health", api.CheckHealth)

	return r
}
