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

func (m *MockUseCase) ProcessMutant(dna *model.RequestMutant) (*model.Mutant, error) {
	mocked := m.Called(dna)
	return mocked.Get(0).(*model.Mutant), mocked.Error(1)
}

type MockRespository struct {
	mock.Mock
}

func (m *MockRespository) Save(mutant *model.Mutant) error {
	mocked := m.Called(mutant)
	return mocked.Error(0)
}

func (m *MockRespository) GetCountByFlag() ([]model.Mutant, error) {
	mocked := m.Called()
	return mocked.Get(0).([]model.Mutant), mocked.Error(1)
}

func MockObjectStats() model.Stats {
	return model.Stats{
		CountMutantDNA: 40,
		CountHumanDNA:  100,
		Ratio:          0.4,
	}
}

func MockSliceStast() []model.Mutant {
	s := make([]model.Mutant, 0)
	s = append(s, model.Mutant{DNA: "dumy", IsMutant: true})
	s = append(s, model.Mutant{DNA: "dumy", IsMutant: false})
	return s
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
