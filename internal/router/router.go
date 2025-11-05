package router

import (
	"gin-todo-list/internal/router/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	group := r.Group("/api")

	group.GET("/health", api.CheckHealth)

	group.GET("/job", api.GetJobList)
	group.GET("/job/:name", api.GetJob)
	group.POST("/job", api.PostJob)
	group.PUT("/job/:name/doing", api.DoingJob)
	group.PUT("/job/:name/done", api.DoneJob)
	group.DELETE("/job/:name", api.DeleteJob)

	return r
}
