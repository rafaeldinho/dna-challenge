package repository

import (
	ctx "context"
	"github/meli/src/domain/model"

	"github.com/spf13/viper"
)

func (r *mutantRepository) Save(mutant *model.Mutant) error {
	_, err := r.firestore.Collection(viper.GetString("COLLECTION")).Doc(mutant.DNA).Set(ctx.Background(), mutant)
	if err != nil {
		return err
	}
	return nil
}
