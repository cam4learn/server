package routes

import (
	"net/http"
	"server/internal/authorization"
	"server/internal/authorizationdata"
	"server/internal/database"
	"server/registration"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setAdminRoutes(group *gin.RouterGroup) {
	newGroup := group.Group("/admin", adminMiddleware)
	newGroup.POST("/addSubject", addSubjectHandler)
	newGroup.DELETE("/deleteSubject", deleteSubjectHandler)
	newGroup.GET("/info/subjects", getSubjectsToDelete)
}

func adminMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("JWT")
	if !authorization.IsAdmin(tokenString) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Next()
}

func getAdminToken(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	login, password := getInput(c)
	loginStruct := authorizationdata.Set{
		Login:     login,
		Password:  password,
		AccessLvl: authorization.Admin,
	}
	token, err := authorization.GetAdminToken(loginStruct)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"JWT": token})
}

func deleteSubjectHandler(c *gin.Context) {
	idString, ok := c.GetPostForm("id")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idString)
	if err == nil {
		database.DeleteSubject(id)
		c.Status(http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusBadRequest)
}

func getSubjectsToDelete(c *gin.Context) {
	data := database.GetSubjectsList()
	c.JSON(http.StatusOK, data)
}

func addSubjectHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var subject registration.SubjectData
	if c.Bind(&subject) == nil && len(subject.Title) > 3 && database.IsExistsLector(subject.LectorID) {
		database.AddSubject(subject)
		c.Status(http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusBadRequest)
}
