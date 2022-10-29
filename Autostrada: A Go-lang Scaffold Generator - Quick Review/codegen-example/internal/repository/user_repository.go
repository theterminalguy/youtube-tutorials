package repository

import (
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (*UserRepository) GetAll() (*[]User, error) {
	return nil, nil
}

func (*UserRepository) Get(id uuid.UUID) (*User, error) {
	return nil, nil
}

func (*UserRepository) Create(e *User) (*User, error) {
	return nil, nil
}

func (r *UserRepository) Update(e *User) (*User, error) {
	return nil, nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	return nil
}
