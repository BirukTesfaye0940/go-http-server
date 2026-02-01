package storage

import (
	"context"
	"errors"
	"sync"

	"go-http-server/models"
)

type UserRepository struct {
	mu    sync.Mutex
	users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []models.User{},
	}
}

func (r *UserRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = len(r.users) + 1
	r.users = append(r.users, user)

	return user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return models.User{}, errors.New("user not found")
}
