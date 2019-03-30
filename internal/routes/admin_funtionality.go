package routes

import (
	"net/http"
	"server/internal/authorization"
	"server/internal/authorizationdata"
	"server/internal/database"
	"server/internal/registration"
	"server/internal/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setAdminRoutes(group *gin.RouterGroup) {
	adminGroup := group.Group("/admin", adminMiddleware)

	adminGroup.POST("/addSubject", addSubjectHandler)
	adminGroup.DELETE("/deleteSubject", deleteSubjectHandler)
	adminGroup.GET("/getSubjects", getSubjectsHandler)

	adminGroup.POST("addLector", addLectorHandler)
	adminGroup.DELETE("/deleteLector", deleteLectorHandler)
	adminGroup.GET("/getLectors", getLectorsHandler)
	adminGroup.PUT("/changeLector", changeLectorHandler)

	adminGroup.POST("/addSubject/", addSubjectHandler)
	adminGroup.DELETE("/deleteSubject/", deleteSubjectHandler)
	adminGroup.GET("/getSubjects/", getSubjectsHandler)

	adminGroup.POST("addLector/", addLectorHandler)
	adminGroup.DELETE("/deleteLector/", deleteLectorHandler)
	adminGroup.GET("/getLectors/", getLectorsHandler)
	adminGroup.PUT("/changeLector/", changeLectorHandler)
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
	login, password, errInput := getInput(c)
	if errInput != nil {
		c.AbortWithError(http.StatusBadRequest, errInput)
		return
	}
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

func addSubjectHandler(c *gin.Context) {
	var subject registration.SubjectData
	if c.Bind(&subject) == nil && len(subject.Title) > 3 && database.IsExistsLector(subject.LectorID) {
		database.AddSubject(subject)
		c.Status(http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusBadRequest)
}

func deleteLectorHandler(c *gin.Context) {
	idString, ok := c.GetPostForm("id")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idString)
	if err == nil && database.IsExistsLector(id) {
		database.DeleteLector(id)
		c.Status(http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusBadRequest)
}

func getLectorsHandler(c *gin.Context) {
	data := database.GetLectorsList()
	c.JSON(http.StatusOK, data)
}

func addLectorHandler(c *gin.Context) {
	var lector registration.LectorData
	if c.Bind(&lector) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ok, err := validator.Validate(lector)
	if ok && err == nil {
		database.AddLector(lector)
		c.Status(http.StatusOK)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"code": err.Error()})
}

func changeLectorHandler(c *gin.Context) {
	idString, ok := c.GetPostForm("lectorId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": "NO_ID"})
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil || !database.IsExistsLector(id) {
		c.JSON(http.StatusBadRequest, gin.H{"code": "WRONG_ID"})
		return
	}
	var lector registration.LectorData
	if c.Bind(&lector) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if len(lector.Password) == 0 {
		lector.Password = database.GetLectorPassword(id)
	}
	valid, errValid := validator.Validate(lector)
	if !valid && errValid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": err.Error()})
		return
	}
	database.UpdateLector(lector, id)
	c.Status(http.StatusOK)
}
