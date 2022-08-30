package mocks

import (
	"github/meli/src/domain/model"

	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) GetStats() (model.Stats, error) {
	mocked := m.Called()
	return mocked.Get(0).(model.Stats), mocked.Error(1)
}

func (m *MockUseCase) IsMutant(dna *model.RequestMutant) (*model.Mutant, error) {
	mocked := m.Called(dna)
	return mocked.Get(0).(*model.Mutant), mocked.Error(1)
}

func MockObjectStats() model.Stats {
	return model.Stats{
		CountMutantDNA: 40,
		CountHumanDNA:  100,
		Ratio:          0.4,
	}
}

func MockMutant(flag bool) *model.Mutant {
	return &model.Mutant{
		DNA:      "CCCCTA,TCACTG",
		IsMutant: flag,
	}
}

func MockRequest() *model.RequestMutant {
	s := make([]string, 0)
	s = append(s, "ATGCGA")
	s = append(s, "CAGTGC")
	return &model.RequestMutant{DNA: s}
}
