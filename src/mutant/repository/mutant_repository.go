package repository

import (
	"github/meli/src/domain/model"

	"cloud.google.com/go/firestore"
)

type mutantRepository struct {
	firestore *firestore.Client
}

type MutantRepository interface {
	Save(mutant *model.Mutant) error
	GetCountByFlag() ([]model.Mutant, error)
}

func NewMutantRepository(firestore *firestore.Client) *mutantRepository {
	return &mutantRepository{firestore: firestore}
}
