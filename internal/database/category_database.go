package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()
	stmt, err := c.db.Prepare("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)")

	if err != nil {
		return Category{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id, name, description)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
