package repository

import (
	"github/meli/src/domain/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type mutantRepository struct {
	client *mongo.Client
}

type MutantRepository interface {
	Save(mutant *model.Mutant) error
	GetCountByFlag() ([]model.Mutant, error)
}

func NewMutantRepository(client *mongo.Client) *mutantRepository {
	return &mutantRepository{client: client}
}
