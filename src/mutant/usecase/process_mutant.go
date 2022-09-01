package usecase

import (
	"strings"

	"github/meli/src/domain/model"
)

func (m *mutantUsecase) ProcessMutant(dna *model.RequestMutant) (*model.Mutant, error) {
	logger.Info("processing mutant")
	mutant := &model.Mutant{
		DNA:      strings.Join(dna.DNA[:], "-"),
		IsMutant: isMutant(dna.DNA),
	}

	if err := m.repository.Save(mutant); err != nil {
		return &model.Mutant{}, err
	}

	return mutant, nil
}

func isMutant(DNA []string) bool {
	var counter int = 0
	var a = [][]string{DNA}
	var n int = len(a)
	var m int = len(a[0])

	for i := 0; i <= m+n-2; i++ {
		for x := 0; x <= i && (x < n) && i-x < m; x++ {
			counter++
		}
		for y := m - 1; y >= 0 && i-y >= 0 && i-y < n && i-y > 0; y-- {
			counter++
		}
	}

	return counter >= 4
}
