package repository

import (
	ctx "context"
	"github/meli/src/domain/model"
)

func (r *mutantRepository) Save(mutant *model.Mutant) error {
	_, err := r.firestore.Collection("stats").Doc(mutant.DNA).Set(ctx.Background(), mutant)
	if err != nil {
		return err
	}
	return nil
}



