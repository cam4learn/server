package database

func DeleteSubject(ID int) {
	dbInstance.Exec("delete from Subject where ID=(?)", ID)
}

func DeleteLector(ID int) {
	dbInstance.Exec("delete from Lector where ID=(?)", ID)
}
