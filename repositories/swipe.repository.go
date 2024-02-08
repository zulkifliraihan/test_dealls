package repositories

import (
	"test_dealls/models"

	"gorm.io/gorm"
)

type SwipeRepository interface {
	FindAll() ([]models.Swipe, error)
	QueryWhere(query string) ([]models.Swipe, error)
	FindByEmail(email string) ([]models.Swipe, error)
	FindById(id int) ([]models.Swipe, error)
	Save(swipe *models.Swipe) (*models.Swipe, error)
}

type swipeRepository struct {
	db *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) SwipeRepository {
	return &swipeRepository{
		db: db,
	}
}

func (r *swipeRepository) Save(swipe *models.Swipe) (*models.Swipe, error) {
	if err := r.db.Preload("SwipingUser").Preload("SwipedUser").Create(&swipe).Error; err != nil {
		return nil, err
	}
	return swipe, nil
}

func (r *swipeRepository) FindAll() ([]models.Swipe, error) {
	var swipes []models.Swipe
	if err := r.db.Preload("SwipingUser").Preload("SwipedUser").Find(&swipes).Error; err != nil {
		return nil, err
	}
	return swipes, nil
}

func (r *swipeRepository) QueryWhere(query string) ([]models.Swipe, error) {
	var swipes []models.Swipe
	if err := r.db.Preload("SwipingUser").Preload("SwipedUser").Where(query).Find(&swipes).Error; err != nil {
		return nil, err
	}
	return swipes, nil
}

func (r *swipeRepository) FindByEmail(email string) ([]models.Swipe, error) {
	var swipes []models.Swipe
	if err := r.db.Preload("SwipingUser").Preload("SwipedUser").Where("email = ?", email).First(&swipes).Error; err != nil {
		return nil, err
	}
	return swipes, nil
}

func (r *swipeRepository) FindById(id int) ([]models.Swipe, error) {
	var swipes []models.Swipe
	if err := r.db.Preload("SwipingUser").Preload("SwipedUser").Where("id = ?", id).First(&swipes).Error; err != nil {
		return nil, err
	}
	return swipes, nil
}

