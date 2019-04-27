package routes

import (
	"fmt"
	"net/http"
	"server/internal/authorization"
	"server/internal/authorizationdata"
	"server/internal/database"
	"server/internal/registration"
	"server/internal/validator"
	"strconv"

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
	adminGroup.GET("/getLectors", getLectorsAdminHandler)
	adminGroup.PATCH("/changeLector", changeLectorHandler)

	adminGroup.POST("/addSubject/", addSubjectHandler)
	adminGroup.DELETE("/deleteSubject/", deleteSubjectHandler)
	adminGroup.GET("/getSubjects/", getSubjectsHandler)

	adminGroup.POST("addLector/", addLectorHandler)
	adminGroup.DELETE("/deleteLector/", deleteLectorHandler)
	adminGroup.GET("/getLectors/", getLectorsAdminHandler)
	adminGroup.PATCH("/changeLector/", changeLectorHandler)

	adminGroup.DELETE("/device", deleteDeviceHandler)
	adminGroup.DELETE("/device/", deleteDeviceHandler)
	adminGroup.GET("/device", getDevicesHandler)
	adminGroup.GET("/device/", getDevicesHandler)
	adminGroup.POST("/device", addDeviceHandler)
	adminGroup.POST("/device/", addDeviceHandler)
	adminGroup.PATCH("/device", editDeviceHandler)
	adminGroup.PATCH("/device/", editDeviceHandler)

	adminGroup.DELETE("/group", deleteGroupHandler)
	adminGroup.DELETE("/group/", deleteGroupHandler)
	adminGroup.GET("/group", getGroupsHandler)
	adminGroup.GET("/group/", getGroupsHandler)
	adminGroup.POST("/group", addGroupHandler)
	adminGroup.POST("/group/", addGroupHandler)
	adminGroup.PATCH("/group", editGroupHandler)
	adminGroup.PATCH("/group/", editGroupHandler)
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
	fmt.Println(idStruct.Id)
	if database.IsExistsLector(idStruct.Id) {
		database.DeleteLector(idStruct.Id)
		c.Status(http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusBadRequest)
}

func getLectorsAdminHandler(c *gin.Context) {
	data := database.GetLectorsListAdmin()
	c.JSON(http.StatusOK, data)
}

func addLectorHandler(c *gin.Context) {
	var lector registration.LectorData
	if c.Bind(&lector) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(lector)
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

func addDeviceHandler(c *gin.Context) {
	var device registration.DeviceData
	if c.Bind(&device) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	valid, errValid := validator.Validate(device)
	if !valid && errValid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errValid.Error()})
		return
	}
	database.AddDevice(device)
	c.Status(http.StatusOK)
}

func deleteDeviceHandler(c *gin.Context) {
	var idStruct idBind
	if c.Bind(&idStruct) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(idStruct.Id)
	database.DeleteDevice(idStruct.Id)
	c.Status(http.StatusOK)
}

func getDevicesHandler(c *gin.Context) {
	data := database.GetDevicesListAdmin()
	c.JSON(http.StatusOK, data)
}

func editDeviceHandler(c *gin.Context) {
	var device registration.DeviceDataEdit
	if c.Bind(&device) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(device)
	valid, errValid := validator.Validate(device)
	if !database.IsExistsDevice(device.Id) {
		c.JSON(http.StatusBadRequest, gin.H{"code": "INVALI_ID"})
		return
	}
	if !valid && errValid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": errValid.Error()})
		return
	}
	database.UpdateDevice(device, device.Id)
	c.Status(http.StatusOK)
}

func getGroupsHandler(c *gin.Context) {
	result := database.GetGroupsAdmin()
	c.JSON(http.StatusOK, result)
}

func addGroupHandler(c *gin.Context) {
	var group registration.GroupAddData
	if c.Bind(&group) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if len(group.Name) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"code": "TO_SHORT_NAME"})
		return
	}
	database.AddGroup(group)
	c.Status(http.StatusOK)
}

func deleteGroupHandler(c *gin.Context) {
	//if c.BindQuery(&idStruct) != nil {
	//	c.AbortWithStatus(http.StatusBadRequest)
	//	return
	//}
	idStr, _ := c.GetQuery("id")
	Id, err := strconv.Atoi(idStr)
	if err != nil || Id == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(Id)
	database.DeleteGroup(Id)
	c.Status(http.StatusOK)
}

func editGroupHandler(c *gin.Context) {
	var group registration.GroupEditData
	if c.Bind(&group) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(group)
	if !database.IsExistsGroup(group.Id) {
		c.JSON(http.StatusBadRequest, gin.H{"code": "INVALI_ID"})
		return
	}
	if len(group.Name) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"code": "BAD_NAME"})
		return
	}
	database.UpdateGroup(group, group.Id)
	c.Status(http.StatusOK)
}
