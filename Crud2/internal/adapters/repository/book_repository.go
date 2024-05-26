package repository

import (
	db "example.com/crudbasics/internal/adapters/repository/database"
	"example.com/crudbasics/internal/domain"
)

type BookRepository interface {
	New(b domain.Book) error
}

type localBookRepository struct{}

func NewLocalBookRepository() *localBookRepository {
	return &localBookRepository{}
}

func (r *localBookRepository) New(b domain.Book) error {
	query := "insert into tblBook (booName, booIdAuthor) values ($1, $2)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(b.BooName, b.BooIdAuthor)

	if err != nil {
		return err
	}

	return nil
}
