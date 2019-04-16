package database

import (
	"database/sql"
	"fmt"
	"server/internal/authorizationdata"

	_ "github.com/go-sql-driver/mysql"
)

var dbInstance *sql.DB

type DataStructureBinder interface {
	BindToFields(*sql.Rows)
}

func InitializeDB(DSN string) {
	var err error
	dbInstance, err = sql.Open("mysql", DSN)
	if err != nil {
		fmt.Println("db wasn't opened")
		panic("DB wasn't opened")
	}
}

func CloseDB() {
	dbInstance.Close()
}

func IsAuthenticatedLector(loginData authorizationdata.Set) bool {
	var result bool
	err := dbInstance.QueryRow("select count(*) from Lector where Login=(?) and Password=(?)", loginData.Login, loginData.Password).Scan(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return result
}

func IsAuthenticatedAdmin(loginData authorizationdata.Set) bool {
	var result bool
	err := dbInstance.QueryRow("select count(*) from Admin where Login=(?) and Password=(?)", loginData.Login, loginData.Password).Scan(&result)
	if err != nil {
		return false
	}
	return result
}

func GetLectorID(loginData authorizationdata.Set) int {
	var result int
	dbInstance.QueryRow("select ID from Lector where Login=(?) and Password=(?)", loginData.Login, loginData.Password).Scan(&result)
	return result
}

func GetAdminID(loginData authorizationdata.Set) int {
	var result int
	dbInstance.QueryRow("select ID from Admin where Login=(?) and Password=(?)", loginData.Login, loginData.Password).Scan(&result)
	return result
}

func IsExistsLector(ID int) bool {
	var result bool
	err := dbInstance.QueryRow("select count(*) from Lector where ID = (?)", ID).Scan(&result)
	if err != nil {
		return false
	}
	return result
}

func IsExistsDevice(ID int) bool {
	var result bool
	err := dbInstance.QueryRow("select count(*) from Device where ID = (?)", ID).Scan(&result)
	if err != nil {
		return false
	}
	return result
}

func GenerateJSONForLecuteCourse(subjectID int) []SecondExport {
	result := make([]SecondExport, 0)
	var subjectName string
	dbInstance.QueryRow("select Title from Subject where ID=(?)", subjectID).Scan(&subjectName)
	rowsLectures := getLecturesForOutput(subjectID)
	for rowsLectures.Next() {
		var lectureID int
		var data SecondExport
		data.Subject = subjectName
		rowsLectures.Scan(&lectureID, &data.Date)
		rowsStudents := getStudentsForOutput(lectureID)
		data.AttendatnceList = dbRowsToObjects(rowsStudents, &StudentExport{}).([]StudentExport)
		result = append(result, data)
	}
	return result
}

func getLecturesForOutput(subjectID int) *sql.Rows {
	result, _ := dbInstance.Query(
		"select ID, Date "+
			"from Lecture "+
			"where SubjectID=(?) "+
			"order by Date", subjectID)
	return result
}

func getStudentsForOutput(lectureID int) *sql.Rows {
	result, _ := dbInstance.Query(
		"select Student.Surname, Student.GroupName, Mark.IsPresent, Mark.Value "+
			"from Mark "+
			"inner join Student on Mark.StudentID = Student.ID "+
			"where Mark.LectureID = (?) "+
			"order by Mark.LectureID, Student.GroupName ", lectureID)
	return result
}

func GenerateJSONForLecuteCourse1(subjectID int) []SecondExport1 {
	result := make([]SecondExport1, 0)
	var subjectName string
	dbInstance.QueryRow("select Title from Subject where ID=(?)", subjectID).Scan(&subjectName)
	rowsLectures := getLecturesForOutput(subjectID)
	for rowsLectures.Next() {
		var lectureID int
		var data SecondExport1
		data.Subject = subjectName
		rowsLectures.Scan(&lectureID, &data.Date)
		rowsStudents := getStudentsForOutput1(lectureID)
		data.AttendatnceList = dbRowsToObjects(rowsStudents, &StudentExport1{}).([]StudentExport1)
		result = append(result, data)
	}
	return result
}

func getStudentsForOutput1(lectureID int) *sql.Rows {
	result, _ := dbInstance.Query(
		"select Student.ID, Student.Surname, Student.GroupName, Mark.IsPresent, Mark.Value "+
			"from Mark "+
			"inner join Student on Mark.StudentID = Student.ID "+
			"where Mark.LectureID = (?) "+
			"order by Mark.LectureID, Student.GroupName ", lectureID)
	return result
}
