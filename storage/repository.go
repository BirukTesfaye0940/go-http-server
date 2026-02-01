package storage

import (
	"context"
	"go-http-server/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	result := r.db.WithContext(ctx).Create(&user)
	return user, result.Error
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	result := r.db.WithContext(ctx).Find(&users)
	return users, result.Error
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).First(&user, id)
	return user, result.Error
}
