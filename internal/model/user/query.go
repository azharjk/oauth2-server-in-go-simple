package user

import (
	"github.com/estradax/exater/internal/model"
)

func Create(user model.User) (model.User, error) {
	data := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	result := model.DB.Create(&data)
	return data, result.Error
}

func FindByEmail(email string) (model.User, bool, error) {
	user := new(model.User)
	result := model.DB.Limit(1).Find(user, model.User{Email: email})

	if result.RowsAffected != 1 {
		return *user, false, result.Error
	}

	return *user, true, result.Error
}
