package api

import "github.com/gin-gonic/gin"

func CheckHealth(c *gin.Context) {
	c.String(200, "OK")
}
