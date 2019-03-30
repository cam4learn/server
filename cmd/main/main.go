package main

import (
	"fmt"
	"server/internal/database"
	"server/internal/routes"
)

func main() {
	//testType := registration.StudentData{}
	//testType.Name = "test"
	//type1 := reflect.TypeOf(testType)
	//field1, _ := type1.FieldByName("Name")
	//fmt.Println(field1.Tag)
	//gin.SetMode(gin.ReleaseMode)
	database.InitializeDB("admin:12345678@/courseProjectDB")
	//result := database.GetSubjectsList()
	//fmt.Println(gin.H{"data": result})
	//recognizer.InitializeRecognizor()
	defer database.CloseDB()
	r := routes.CreateRoutes()
	fmt.Println(r.RedirectTrailingSlash)
	//recognizer.Teach("", 5)
	r.Run(":8030")
}
