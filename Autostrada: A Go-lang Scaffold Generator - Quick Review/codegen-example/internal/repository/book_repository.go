package repository

import (
	"github.com/google/uuid"
)

type Book struct {
	ID uuid.UUID
}

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (*BookRepository) GetAll() (*[]Book, error) {
	return nil, nil
}

func (*BookRepository) Get(id uuid.UUID) (*Book, error) {
	return nil, nil
}

func (*BookRepository) Create(e *Book) (*Book, error) {
	return nil, nil
}

func (r *BookRepository) Update(e *Book) (*Book, error) {
	return nil, nil
}

func (r *BookRepository) Delete(id uuid.UUID) error {
	return nil
}
