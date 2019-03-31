package database

import "server/internal/registration"

func UpdateLector(form registration.LectorDataEdit, ID int) {
	dbInstance.Exec("update Lector set Name = (?), Surname = (?), Login = (?), Password = (?) where ID = (?)",
		form.Name, form.Surname, form.Login, form.Password, ID)
}
