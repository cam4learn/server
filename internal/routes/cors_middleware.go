package routes

import (
	"github.com/gin-gonic/gin"
)

func RequestPermission(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "AuthID,Origin, X-Requested-With, Content-Type, Accept, JWT")
	c.Header("Access-Control-Allow-Methods", "PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
	} else {
		c.Next()
	}
}
