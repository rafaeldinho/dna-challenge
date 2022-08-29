package repository

import (
	"github/meli/src/domain/model"
	"github/meli/src/infrastructure/sql"
)

type mutantRepository struct {
	Conn sql.SqlRepository
}

type MutantRepository interface {
	Save(mutant *model.Mutant) error
	GetCountByFlag(flag bool) (*int64, error)
}

func NewMutantRepository(Conn sql.SqlRepository) *mutantRepository {
	return &mutantRepository{Conn: Conn}
}
