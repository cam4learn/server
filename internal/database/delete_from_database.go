package database

func DeleteSubject(ID int) {
	dbInstance.Exec("delete from Subject where ID=(?)", ID)
}
