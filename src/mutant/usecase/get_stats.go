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
		Ratio:          float64(mutantCount) / float64(humanCount),
	}, nil
}

func getCountByFlag(list []model.Mutant) (int, int) {
	var humanCount = 0
	var mutantCount = 0
	for _, m := range list {
		if m.IsMutant {
			mutantCount++
			continue
		}
		humanCount++
	}

	return humanCount, mutantCount
}
