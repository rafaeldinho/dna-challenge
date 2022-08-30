package repository

import (
	ctx "context"
	"github/meli/src/domain/model"

	"github.com/spf13/viper"
)

func (r *mutantRepository) GetCountByFlag() ([]model.Mutant, error) {
	var listResult []model.Mutant

	iter := r.firestore.Collection(viper.GetString("COLLECTION")).Documents(ctx.Background())
	docs, err := iter.GetAll()
	if err != nil {
		return listResult, err
	}

	for _, doc := range docs {
		var mutant model.Mutant
		doc.DataTo(&mutant)
		listResult = append(listResult, mutant)
	}

	return listResult, nil

}
