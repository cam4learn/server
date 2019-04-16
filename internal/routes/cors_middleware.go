package routes

import (
	"github.com/gin-gonic/gin"
)

func RequestPermission(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, JWT")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, OPTIONS, POST, PATCH, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
	} else {
		c.Next()
	}
}
