package main

import (
	"fmt"
	"server/internal/database"
	"server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//testType := registration.StudentData{}
	//testType.Name = "test"
	//type1 := reflect.TypeOf(testType)
	//field1, _ := type1.FieldByName("Name")
	//fmt.Println(field1.Tag)
	database.InitializeDB("admin:admin@/courseProjectDB")
	result := database.GetSubjectsList()
	fmt.Println(gin.H{"data": result})
	//recognizer.InitializeRecognizor()
	defer database.CloseDB()
	r := routes.CreateRoutes()
	//recognizer.Teach("", 5)
	r.Run(":8030")
}
