package repository

import (
	ctx "context"
	"github/meli/src/domain/model"
	"github/meli/src/shared"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutantRepository) GetCountByFlag() ([]model.Mutant, error) {
	collection := r.client.Database(shared.DatabaseCollection).Collection(shared.MutantCollection)

	var results []model.Mutant

	rstSearch, err := collection.Find(ctx.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	for rstSearch.Next(ctx.Background()) {
		var mutant model.Mutant
		if err := rstSearch.Decode(&mutant); err != nil {
			return nil, err
		}
		results = append(results, mutant)
	}

	if err := rstSearch.Err(); err != nil {
		return nil, err
	}

	rstSearch.Close(ctx.Background())

	return results, nil

}
