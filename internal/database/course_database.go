package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryId string) (Course, error) {
	id := uuid.New().String()
	stmt, err := c.db.Prepare("INSERT INTO courses (id, name, description, categoryId) VALUES (?, ?, ?, ?)")
	if err != nil {
		return Course{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description, categoryId)
	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description}, nil

}
