package database

import "server/internal/registration"

func UpdateLector(form registration.LectorDataEdit, ID int) {
	dbInstance.Exec("update Lector set Name = (?), Surname = (?), Login = (?), Password = (?) where ID = (?)",
		form.Name, form.Surname, form.Login, form.Password, ID)
}

func UpdateDevice(form registration.DeviceDataEdit, ID int) {
	dbInstance.Exec("update Device set Room = (?), MACAdress = (?) where ID = (?)",
		form.Room, form.MACAdress, ID)
}
