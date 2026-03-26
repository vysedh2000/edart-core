package repositories

import (
	"edart-core/app/models"
	"edart-core/database"
	"fmt"

	"gorm.io/gorm"
)

type UserRespository struct {
	DB * gorm.DB
}

func NewUserRepo() *UserRespository{
	db := database.DB
	return &UserRespository{DB: db}
}

func (u*UserRespository) CreateUser(user models.UserInfo) error {
	return u.DB.Create(&user).Error
}

func (r *UserRespository) GetNextUserId() (string, error) {
	var seq int64

	err := r.DB.Raw("SELECT nextval('user_sequence')").Scan(&seq).Error
	if err != nil {
		return "", err
	}

	acc := fmt.Sprintf("U%08d", seq)
	return acc, nil
}