package usecase

import (
	"github/meli/src/domain/model"
	"strings"
)

func (m *mutantUsecase) IsMutant(dna []string) (*model.Mutant, error) {
	// logica loca
	mutant := &model.Mutant{
		ADN:      strings.Join(dna[:], ","),
		IsMutant: true,
	}

	if err := m.repository.Save(mutant); err != nil {
		return &model.Mutant{}, err
	}

	return mutant, nil
}
