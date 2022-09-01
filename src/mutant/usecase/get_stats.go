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

	ramon := model.Stats{
		CountMutantDNA: mutantCount,
		CountHumanDNA:  humanCount,
		Ratio:          getRatio(humanCount, mutantCount),
	}

	return ramon, nil
}

func getRatio(h int, m int) float64 {
	var ratio float64 = 0
	if h == 0 && m == 0 {
		return ratio
	}
	ratio = float64(h) / float64(m)
	return ratio
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
