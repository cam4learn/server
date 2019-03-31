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
	result.Use(RequestPermission)
	result.POST("/adminLogin", getAdminToken)
	result.POST("/adminLogin/", getAdminToken)
	result.POST("/login", getToken)
	result.POST("/login/", getToken)
	result.Static("/swaggerui/", "./swaggerui")
	authorized := result.Group("/api", authMiddleware)
	authorized.GET("/getSubjects", getSubjectsHandler)
	authorized.GET("/getLectorSubjects", getLectorSubjectsHandler)
	authorized.GET("/getSubjects/", getSubjectsHandler)
	authorized.GET("/getLectorSubjects/", getLectorSubjectsHandler)
	authorized.GET("lectors", getLectorsHandler)

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
	login, password, errInput := getInput(c)
	if errInput != nil {
		c.AbortWithError(http.StatusBadRequest, errInput)
		return
	}
	loginStruct := authorizationdata.Set{
		Login:     login,
		Password:  password,
		AccessLvl: authorization.Lector,
	}
	token, err := authorization.GetLectorToken(loginStruct)
	if err != nil {
		c.AbortWithStatus(401)
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
