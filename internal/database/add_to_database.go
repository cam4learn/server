package database

import "server/internal/registration"

func AddSubject(form registration.SubjectData) {
	dbInstance.Exec("insert into Subject (LectorID, Title) values (?),(?)", form.LectorID, form.Title)
}
