package repositories

import (
	"rest-api/internal/database"
	"rest-api/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.GetDB().Create(user).Error

}
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.GetDB().Where("email = ?", email).First(&user).Error
	return &user, err
}
