package routes

import (
	"errors"
	"net/http"
	"server/internal/authorization"
	"server/internal/authorizationdata"
	"server/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func CreateRoutes() *gin.Engine {
	result := gin.Default()
	result.Use(RequestPermission)
	result.POST("/adminLogin", getAdminToken)
	result.POST("/adminLogin/", getAdminToken)
	result.POST("/login", getToken)
	result.POST("/login/", getToken)
	result.GET("/getSubjects", getSubjectsHandler)
	result.GET("/subjectStatisticJson", getSubjectStatisticJsonHandler1)
	result.GET("/subjectStatisticJson/", getSubjectStatisticJsonHandler1)
	result.GET("/subjectStatisticCsv", getSubjectStatisticCsvHandler)
	result.GET("/subjectStatisticCsv/", getSubjectStatisticCsvHandler)
	result.GET("/getSubjectsCsv", getSubjectsCsvHandler)
	result.Static("/swaggerui/", "./swaggerui")
	authorized := result.Group("/api", authMiddleware)
	authorized.GET("/getSubjects", getSubjectsHandler)
	authorized.GET("/getLectorSubjects", getLectorSubjectsHandler)
	authorized.GET("/getSubjects/", getSubjectsHandler)
	authorized.GET("/getLectorSubjects/", getLectorSubjectsHandler)

	authorized.GET("/lectors", getLectorsHandler)

	authorized.GET("/subjectStatisticJson", getSubjectStatisticJsonHandler)
	authorized.GET("/subjectStatisticJson/", getSubjectStatisticJsonHandler)
	authorized.GET("/subjectStatisticCsv", getSubjectStatisticCsvHandler)
	authorized.GET("/subjectStatisticCsv/", getSubjectStatisticCsvHandler)
	authorized.GET("/getSubjectsCsv/", getSubjectsCsvHandler)

	authorized.GET("/getLectorSubjectsCsv/", getSubjectsCsvHandler)
	setAdminRoutes(authorized)
	return result
}

func authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("JWT")
	if tokenString == "" || !authorization.ValidateToken(tokenString) {
		c.AbortWithError(400, errors.New("BAD_TOKEN"))
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

func getLectorsHandler(c *gin.Context) {
	data := database.GetLectorsList()
	c.JSON(http.StatusOK, data)
}

// func getSubjectStatisticHandler(c *gin.Context) {
// 	idString := c.Query("subjectId")
// 	format := c.Param("format")
// 	id, _ := strconv.Atoi(idString)
// 	fmt.Println(id)
// 	result := database.GenerateJSONForLecuteCourse(id)
// 	fmt.Println(result)
// 	if format == "JSON" {
// 		c.JSON(http.StatusOK, result)
// 	} else if format == "CSV" {
// 		csv, _ := gocsv.MarshalString(database.FromSecondExportToCSVStruct(result))
// 		c.Header("Content-Type", "text/csv")
// 		c.Writer.Write([]byte(csv))
// 		c.Writer.Flush()
// 		c.Status(http.StatusOK)
// 	}
// }

func getSubjectStatisticJsonHandler(c *gin.Context) {
	idString := c.Query("subjectId")
	id, _ := strconv.Atoi(idString)
	result := database.GenerateJSONForLecuteCourse(id)
	c.JSON(http.StatusOK, result)
}

func getSubjectStatisticCsvHandler(c *gin.Context) {
	idString := c.Query("subjectId")
	id, _ := strconv.Atoi(idString)
	result := database.GenerateJSONForLecuteCourse(id)
	csv, _ := gocsv.MarshalString(database.FromSecondExportToCSVStruct(result))
	c.Header("Content-Type", "text/csv")
	c.Writer.Write([]byte(csv))
	c.Writer.Flush()
	c.Status(http.StatusOK)
}

func getSubjectsCsvHandler(c *gin.Context) {
	result := database.GetSubjectsList()
	csv, _ := gocsv.MarshalString(result)
	c.Header("Content-Type", "text/csv")
	c.Writer.Write([]byte(csv))
	c.Writer.Flush()
	c.Status(http.StatusOK)
}

func getSubjectStatisticJsonHandler1(c *gin.Context) {
	idString := c.Query("subjectId")
	id, _ := strconv.Atoi(idString)
	result := database.GenerateJSONForLecuteCourse1(id)
	c.JSON(http.StatusOK, result)
}

func getSubjectStatisticCsvHandler1(c *gin.Context) {
	//idString := c.Query("subjectId")
	//id, _ := strconv.Atoi(idString)
	//result := database.GenerateJSONForLecuteCourse1(id)
	//csv, _ := gocsv.MarshalString(database.FromSecondExportToCSVStruct(result))
	//c.Header("Content-Type", "text/csv")
	//c.Writer.Write([]byte(csv))
	//c.Writer.Flush()
	//c.Status(http.StatusOK)
}
