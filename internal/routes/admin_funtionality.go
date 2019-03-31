package routes

import (
	"fmt"
	"net/http"
	"server/internal/authorization"
	"server/internal/authorizationdata"
	"server/internal/database"
	"server/internal/registration"
	"server/internal/validator"

	"github.com/gin-gonic/gin"
)

type idBind struct {
	Id int `json:"id"`
}

func setAdminRoutes(group *gin.RouterGroup) {
	adminGroup := group.Group("/admin", adminMiddleware)

	adminGroup.POST("/addSubject", addSubjectHandler)
	adminGroup.DELETE("/deleteSubject", deleteSubjectHandler)
	adminGroup.GET("/getSubjects", getSubjectsHandler)

	adminGroup.POST("addLector", addLectorHandler)
	adminGroup.DELETE("/deleteLector", deleteLectorHandler)
	adminGroup.GET("/getLectors", getLectorsHandler)
	adminGroup.PATCH("/changeLector", changeLectorHandler)

	adminGroup.POST("/addSubject/", addSubjectHandler)
	adminGroup.DELETE("/deleteSubject/", deleteSubjectHandler)
	adminGroup.GET("/getSubjects/", getSubjectsHandler)

	adminGroup.POST("addLector/", addLectorHandler)
	adminGroup.DELETE("/deleteLector/", deleteLectorHandler)
	adminGroup.GET("/getLectors/", getLectorsHandler)
	adminGroup.PATCH("/changeLector/", changeLectorHandler)
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
		c.AbortWithStatus(401)
		return
	}
	c.JSON(http.StatusOK, gin.H{"JWT": token})
}

func deleteSubjectHandler(c *gin.Context) {
	var idStruct idBind
	if c.Bind(&idStruct) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(idStruct.Id)
	database.DeleteSubject(idStruct.Id)
	c.Status(http.StatusOK)
	return
}

func addSubjectHandler(c *gin.Context) {
	var subject registration.SubjectData
	//&& database.IsExistsLector(subject.LectorID)
	if c.Bind(&subject) == nil && len(subject.Title) > 3 {
		fmt.Println(subject)
		database.AddSubject(subject)
		c.Status(http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusBadRequest)
}

func deleteLectorHandler(c *gin.Context) {
	var idStruct idBind
	if c.Bind(&idStruct) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if database.IsExistsLector(idStruct.Id) {
		database.DeleteLector(idStruct.Id)
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
	var lector registration.LectorDataEdit
	if c.Bind(&lector) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(lector)
	if len(lector.Password) == 0 {
		lector.Password = database.GetLectorPassword(lector.Id)
	}
	valid, errValid := validator.Validate(lector)
	if !database.IsExistsLector(lector.Id) {
		c.JSON(http.StatusBadRequest, gin.H{"code": "INVALI_ID"})
		return
	}
	if !valid && errValid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errValid.Error()})
		return
	}
	database.UpdateLector(lector, lector.Id)
	c.Status(http.StatusOK)
}
