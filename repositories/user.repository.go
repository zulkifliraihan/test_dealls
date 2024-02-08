package repositories

import (
	"test_dealls/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	QueryWhere(query string, args ...interface{}) ([]models.User, error)
	FindByEmail(email string) ([]models.User, error)
	FindById(id int) ([]models.User, error)
	Save(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Save(user *models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) QueryWhere(query string, args ...interface{}) ([]models.User, error) {
	var users []models.User
	if err := r.db.Where(query, args...).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByEmail(email string) ([]models.User, error) {
	var users []models.User
	if err := r.db.Where("email = ?", email).First(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindById(id int) ([]models.User, error) {
	var users []models.User
	if err := r.db.Where("id = ?", id).First(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

