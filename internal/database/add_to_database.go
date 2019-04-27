package database

import (
	"server/internal/registration"
)

func AddSubject(form registration.SubjectData) {
	dbInstance.Exec("insert into Subject (LectorID, Title) values ((?), (?))", form.LectorID, form.Title)
}

func AddLector(form registration.LectorData) {
	dbInstance.Exec("insert into Lector (Name, Surname, Login, Password) values ((?), (?), (?), (?))",
		form.Name, form.Surname, form.Login, form.Password)
}

func AddDevice(form registration.DeviceData) {
	dbInstance.Exec("insert into Device (Room, MACAdress) values ((?), (?))",
		form.Room, form.MACAdress)
}

func AddGroup(form registration.GroupAddData) {
	dbInstance.Exec("insert into StudentGroup (Name) values ((?))", form.Name)
}
