package database

import (
	"database/sql"
	"reflect"
)

func GetLectorSubjects(LectorID int) []Subject {
	resultRows, _ := dbInstance.Query("select Subject.ID, Subject.Title, Lector.Name, Lector.Surname from Subject join Lector on Lector.ID = Subject.LectoID where LectorID=(?)", LectorID)
	result := dbRowsToObjects(resultRows, &Subject{}).([]Subject)
	return result
}

func GetSubjectsList() []Subject {
	resultRows, _ := dbInstance.Query("select Subject.ID, Subject.Title, Lector.Name, Lector.Surname from Subject inner join Lector on Lector.ID = Subject.LectorID")
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

func GetLectorsListAdmin() []LectorAdminData {
	resultRows, _ := dbInstance.Query("select ID, Name, Surname, Login from Lector")
	result := dbRowsToObjects(resultRows, &LectorAdminData{}).([]LectorAdminData)
	return result
}

func GetLectorPassword(ID int) string {
	var result string
	dbInstance.QueryRow("select Password from Lector where ID = (?)", ID).Scan(&result)
	return result
}

func GetDevicesList() []Device {
	resultRows, _ := dbInstance.Query("select ID, Room from Device")
	result := dbRowsToObjects(resultRows, &Device{}).([]Device)
	return result
}

func GetDevicesListAdmin() []DeviceAdminData {
	resultRows, _ := dbInstance.Query("Select ID, Room, MACAdress from Device")
	result := dbRowsToObjects(resultRows, &DeviceAdminData{}).([]DeviceAdminData)
	return result
}
