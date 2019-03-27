package database

import (
	"database/sql"
)

type Subject struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (s *Subject) BindToFields(row *sql.Rows) {
	row.Scan(&s.ID, &s.Title)
}

type Lector struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (l *Lector) BindToFields(row *sql.Rows) {
	row.Scan(&l.ID, &l.Name, &l.Surname)
}
