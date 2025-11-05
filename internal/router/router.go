package router

import (
	"gin-todo-list/internal/router/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	group := r.Group("/api")
	{
		group.GET("/health", api.CheckHealth)

		// Job API
		group.GET("/job", api.GetJobList)
		group.GET("/job/:name", api.GetJob)
		group.POST("/job", api.PostJob)
		group.PUT("/job/:name/doing", api.DoingJob)
		group.PUT("/job/:name/done", api.DoneJob)
		group.DELETE("/job/:name", api.DeleteJob)
	}

	// Job View
	r.StaticFile("/", "internal/view/index.html")
	r.StaticFile("/job", "internal/view/job.html")

	// CSS
	r.Static("/assets", "internal/view/assets")

	return r
}
