package client

import "github.com/estradax/exater/internal/model"

func FindByID(id string) (model.Client, bool, error) {
	c := new(model.Client)
	result := model.DB.Limit(1).Find(c, "id = ?", id)

	if result.RowsAffected != 1 {
		return *c, false, result.Error
	}

	return *c, true, result.Error
}
