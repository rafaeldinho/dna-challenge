package repository

import (
	"github/meli/src/domain/model"
)

func (r *mutantRepository) Save(mutant *model.Mutant) error {
	db := r.Conn.GetConnection()

	if err := db.Create(&mutant).Error; err != nil {
		return err
	}
	return nil
}
