package repositories

import (
	"fmt"
	"log"

	"github.com/iufb/goth/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Create(user models.User) error
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r *GormUserRepository) FindByEmail(email string) (*models.User, error) {
	u := new(models.User)
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User not found")
		}
		return nil, fmt.Errorf("Db error")
	}
	log.Println(u)
	return u, nil
}

func (r *GormUserRepository) Create(user models.User) error {
	res := r.db.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
