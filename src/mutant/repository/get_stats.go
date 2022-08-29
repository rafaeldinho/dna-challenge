package repository

import (
	"github/meli/src/shared"
)

func (r *mutantRepository) GetCountByFlag(flag bool) (*int64, error) {
	db := r.Conn.GetConnection()

	var count *int64

	if tx := db.Table(shared.MutantTable).Count(count).Where("is_mutant = ?", flag); tx.Error != nil {
		return count, tx.Error
	}
	return count, nil
}
