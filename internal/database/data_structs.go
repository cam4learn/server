package database

import (
	"database/sql"
)

type Subject struct {
	ID            int    `json:"id" csv:"id"`
	Title         string `json:"title" csv:"title"`
	LectorName    string `json:"name" csv:"name"`
	LectorSurname string `json:"surname" csv:"surname"`
}

func (s *Subject) BindToFields(row *sql.Rows) {
	row.Scan(&s.ID, &s.Title, &s.LectorName, &s.LectorSurname)
}

type Lector struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (l *Lector) BindToFields(row *sql.Rows) {
	row.Scan(&l.ID, &l.Name, &l.Surname)
}

type LectorAdminData struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Login   string `json:"login"`
}

func (l *LectorAdminData) BindToFields(row *sql.Rows) {
	row.Scan(&l.ID, &l.Name, &l.Surname, &l.Login)
}

type Device struct {
	ID   int    `json:"id"`
	Room string `json:"room"`
}

func (d *Device) BindToFields(row *sql.Rows) {
	row.Scan(&d.ID, &d.Room)
}

type DeviceAdminData struct {
	ID   int    `json:"id"`
	Room string `json:"room"`
	Mac  string `json:"address"`
}

func (d *DeviceAdminData) BindToFields(row *sql.Rows) {
	row.Scan(&d.ID, &d.Room, &d.Mac)
}

type StudentExport struct {
	Surname   string `json:"surname" csv:"surname"`
	Group     string `json:"group" csv:"group"`
	Value     int64  `json:"value" csv:"value"`
	IsPresent bool   `json:"isPresent" csv:"isPresent"`
}

func (s *StudentExport) BindToFields(row *sql.Rows) {
	row.Scan(&s.Surname, &s.Group, &s.IsPresent, &s.Value)
}

type SecondExport struct {
	Date            string          `json:"date" csv:"date"`
	Subject         string          `json:"subject" csv:"subject"`
	AttendatnceList []StudentExport `json:"attendanceList" csv:"attendanceList"`
}

type SecondExportCSV struct {
	Date      string `json:"date" csv:"date"`
	Surname   string `json:"surname" csv:"surname"`
	Group     string `json:"group" csv:"group"`
	Value     int64  `json:"value" csv:"value"`
	IsPresent bool   `json:"isPresent" csv:"isPresent"`
}

func FromSecondExportToCSVStruct(input []SecondExport) []SecondExportCSV {
	result := make([]SecondExportCSV, 0)
	for _, a := range input {
		date := a.Date
		for _, attendance := range a.AttendatnceList {
			var toAdd SecondExportCSV
			toAdd.Date = date
			toAdd.Group = attendance.Group
			toAdd.IsPresent = attendance.IsPresent
			toAdd.Surname = attendance.Surname
			toAdd.Value = attendance.Value
			result = append(result, toAdd)
		}
	}
	return result
}

type StudentExport1 struct {
	Id        int    `json:"id" csv:"id"`
	Surname   string `json:"surname" csv:"surname"`
	Group     string `json:"group" csv:"group"`
	Value     int64  `json:"value" csv:"value"`
	IsPresent bool   `json:"isPresent" csv:"isPresent"`
}

func (s *StudentExport1) BindToFields(row *sql.Rows) {
	row.Scan(&s.Id, &s.Surname, &s.Group, &s.IsPresent, &s.Value)
}

type SecondExport1 struct {
	Date            string           `json:"date" csv:"date"`
	Subject         string           `json:"subject" csv:"subject"`
	AttendatnceList []StudentExport1 `json:"attendanceList" csv:"attendanceList"`
}
