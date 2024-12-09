package repository

import (
	"database/sql"

	"github.com/romulosm/go-modular-backend-template/internal/user/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", user.ID, user.Name, user.Email)
	return err
}

func (r *UserRepository) GetByID(id string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
