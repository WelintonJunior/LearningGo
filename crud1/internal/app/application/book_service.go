package application

import (
	"example.com/crud1/internal/adapters/repository"
	"example.com/crud1/internal/app/domain"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) NewBook(b domain.Book) error {
	return s.repo.NewBook(b)
}
