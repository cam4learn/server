package database

import (
	"database/sql"
	"server/internal/authorizationdata"
	"server/registration"

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

func DeleteSubject(ID int) {
	dbInstance.Exec("delete from Subject where ID=(?)", ID)
}

func IsExistsLector(ID int) bool {
	var result bool
	err := dbInstance.QueryRow("select count(*) from Lector where ID = (?)", ID).Scan(&result)
	if err != nil {
		return false
	}
	return result
}

func AddSubject(form registration.SubjectData) {
	dbInstance.Exec("insert into Subject (LectorID, Title) values (?),(?)", form.LectorID, form.Title)
}