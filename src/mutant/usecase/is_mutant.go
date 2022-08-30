package usecase

import (
	"github/meli/src/domain/model"
	"strings"
)

func (m *mutantUsecase) IsMutant(dna *model.RequestMutant) (*model.Mutant, error) {
	// logica loca
	mutant := &model.Mutant{
		DNA:      strings.Join(dna.DNA[:], ","),
		IsMutant: true,
	}

	if err := m.repository.Save(mutant); err != nil {
		return &model.Mutant{}, err
	}

	return mutant, nil
}
