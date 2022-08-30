package usecase

import "github/meli/src/domain/model"

func (m *mutantUsecase) GetStats() (model.Stats, error) {
	var err error
	var humanCount int
	var mutantCount int

	mutantList, err := m.repository.GetCountByFlag()

	if err != nil {
		return model.Stats{}, err
	}

	humanCount, mutantCount = getCountByFlag(mutantList)

	return model.Stats{
		CountMutantDNA: mutantCount,
		CountHumanDNA:  humanCount,
		Ratio:          2.6,
	}, nil
}

func getCountByFlag(list []model.Mutant) (int, int) {
	var humanCount int
	var mutantCount int
	for _, m := range list {
		if m.IsMutant {
			humanCount++
		}
		if !m.IsMutant {
			mutantCount++
		}
	}

	return humanCount, mutantCount
}
