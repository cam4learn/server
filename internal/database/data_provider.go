package database

import (
	"database/sql"
	"reflect"
)

func GetLectorSubjects(LectorID int) []Subject {
	resultRows, _ := dbInstance.Query("select ID, Title from Subject where LectorID=(?)", LectorID)
	result := dbRowsToObjects(resultRows, &Subject{}).([]Subject)
	return result
}

func GetSubjectsList() []Subject {
	resultRows, _ := dbInstance.Query("select ID, Title from Subject")
	result := dbRowsToObjects(resultRows, &Subject{}).([]Subject)
	return result
}

func dbRowsToObjects(rows *sql.Rows, destinationType DataStructureBinder) interface{} {
	resultElementType := reflect.TypeOf(destinationType).Elem()
	result := reflect.MakeSlice(reflect.SliceOf(resultElementType), 0, 0)
	for rows.Next() {
		data := reflect.New(resultElementType)
		data.MethodByName("BindToFields").Call([]reflect.Value{reflect.ValueOf(rows)})
		result = reflect.Append(result, reflect.Indirect(data))
	}
	return result.Interface()
}

func GetLectorsList() []Lector {
	resultRows, _ := dbInstance.Query("select ID, Name, Surname from Lector")
	result := dbRowsToObjects(resultRows, &Lector{}).([]Lector)
	return result
}
