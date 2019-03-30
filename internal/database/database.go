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
