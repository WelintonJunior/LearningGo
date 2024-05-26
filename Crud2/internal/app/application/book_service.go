package application

import (
	"example.com/crudbasics/internal/adapters/repository"
	"example.com/crudbasics/internal/domain"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) New(b domain.Book) error {
	return s.repo.New(b)
}
