package repository

import (
	"log"
	"sync"

	db "example.com/crud1/internal/adapters/repository/database"
	"example.com/crud1/internal/app/domain"
)

type BookRepository interface {
	NewBook(b domain.Book) error
	ListBooks() ([]domain.Book, error)
	ListBooksNoRoutine() ([]domain.Book, error)
}

type localBookRepository struct{}

func NewLocalBookRepository() *localBookRepository {
	return &localBookRepository{}
}

func (r *localBookRepository) NewBook(b domain.Book) error {
	query := "insert into tblBook (booName, booIdAuthor) values ($1, $2)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	for i := 0; i < 10000; i++ {
		_, err = stmt.Exec(b.BooName, b.BooIdAuthor)

		if err != nil {
			return err
		}
	}

	return nil
}
func (r *localBookRepository) ListBooks() ([]domain.Book, error) {
	query := "SELECT BooId, BooName, BooIdAuthor, BooCreatedAt FROM tblBook"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book
	var wg sync.WaitGroup

	bookChan := make(chan domain.Book)
	done := make(chan struct{})
	const numWorkers = 10

	worker := func() {
		for b := range bookChan {
			books = append(books, b)
			wg.Done()
		}
	}

	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	go func() {
		defer close(done)
		for rows.Next() {
			var b domain.Book
			if err := rows.Scan(&b.BooId, &b.BooName, &b.BooIdAuthor, &b.BooCreatedAt); err != nil {
				log.Println("Error scanning row:", err)
				continue
			}
			wg.Add(1)
			bookChan <- b
		}
		wg.Wait()
		close(bookChan)
	}()

	<-done
	return books, nil
}

func (r *localBookRepository) ListBooksNoRoutine() ([]domain.Book, error) {
	query := "SELECT BooId, BooName, BooIdAuthor, BooCreatedAt FROM tblBook"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var b domain.Book
		if err := rows.Scan(&b.BooId, &b.BooName, &b.BooIdAuthor, &b.BooCreatedAt); err != nil {
			log.Println("Error scanning row: ", err)
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
