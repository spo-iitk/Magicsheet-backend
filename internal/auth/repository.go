package auth

import (
	"context"

	"github.com/spo-iitk/Magicsheet-backend/internal/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*database.User, error) {

	var user database.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id uint) (*database.User, error) {
	var user database.User

	err := r.db.WithContext(ctx).Where("ID = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *database.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
