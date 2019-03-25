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
