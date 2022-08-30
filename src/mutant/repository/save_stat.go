package repository

import (
	ctx "context"
	"github/meli/src/domain/model"
	"github/meli/src/shared"
)

func (r *mutantRepository) Save(mutant *model.Mutant) error {
	collection := r.client.Database(shared.DatabaseCollection).Collection(shared.MutantCollection)

	if _, err := collection.InsertOne(ctx.Background(), mutant); err != nil {
		return err
	}

	return nil

}
