package routes

import (
	"net/http"
	"server/internal/authorization"
	"server/internal/authorizationdata"
	"server/internal/database"

	"github.com/gin-gonic/gin"
)

func CreateRoutes() *gin.Engine {
	result := gin.Default()
	result.POST("/adminLogin", getAdminToken)
	result.POST("/login", getToken)
	result.Use(RequestPermission)
	authorized := result.Group("/api", authMiddleware)
	authorized.GET("/getSubjects", getSubjectsHandler)
	authorized.GET("/getLectorSubjects", getLectorSubjectsHandler)

	setAdminRoutes(authorized)
	return result
}

func authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("JWT")
	if tokenString == "" || !authorization.ValidateToken(tokenString) {
		c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
		return
	}
	c.Next()
}

func getToken(c *gin.Context) {
	login, password := getInput(c)
	loginStruct := authorizationdata.Set{
		Login:     login,
		Password:  password,
		AccessLvl: authorization.Lector,
	}
	token, err := authorization.GetLectorToken(loginStruct)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"JWT": token})
}

func getSubjectsHandler(c *gin.Context) {
	result := database.GetSubjectsList()
	c.JSON(http.StatusOK, result)
}

func getLectorSubjectsHandler(c *gin.Context) {
	id := authorization.GetIDFromToken(c.GetHeader("JWT"))
	result := database.GetLectorSubjects(id)
	c.JSON(http.StatusOK, result)
}
