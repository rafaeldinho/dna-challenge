package usecase

import "github/meli/src/domain/model"

func (m *mutantUsecase) GetStats() (model.Stats, error) {
	var err error
	var humanCount *int64
	var mutantCount *int64

	mutantCount, err = m.repository.GetCountByFlag(true)

	if err != nil {
		return model.Stats{}, err
	}

	humanCount, err = m.repository.GetCountByFlag(false)

	if err != nil {
		return model.Stats{}, err
	}

	return model.Stats{
		CountMutantADN: mutantCount,
		CountHumanADN:  humanCount,
		Ratio:          2.6,
	}, nil
}
