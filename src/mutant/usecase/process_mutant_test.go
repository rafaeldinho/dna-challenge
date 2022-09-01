package usecase

import (
	"errors"
	"github/meli/src/domain/model"
	"github/meli/src/mutant/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIsMutant(t *testing.T) {

	t.Run("When calling ProcessMutant should create and detect dna person without error", func(t *testing.T) {

		repository := new(mocks.MockRespository)
		usecase := NewMutantUsecase(repository)

		repository.On("Save", mock.Anything).Return(nil)

		s := make([]string, 0)
		s = append(s, "ATGCGA")
		s = append(s, "CAGTGC")
		request := &model.RequestMutant{DNA: s}

		_, err := usecase.ProcessMutant(request)

		assert.NoError(t, err)
		repository.AssertCalled(t, "Save", mock.Anything)
	})

	t.Run("When calling ProcessMutant should create and detect dna person without error", func(t *testing.T) {

		repository := new(mocks.MockRespository)
		usecase := NewMutantUsecase(repository)

		repository.On("Save", mock.Anything).Return(nil)

		s := make([]string, 0)
		s = append(s, "ATGCGA")
		s = append(s, "CAGTGC")
		s = append(s, "TTATGT")
		s = append(s, "AGAAGG")
		s = append(s, "CCCCTA")
		s = append(s, "TCACTG")
		request := &model.RequestMutant{DNA: s}

		_, err := usecase.ProcessMutant(request)

		assert.NoError(t, err)
		repository.AssertCalled(t, "Save", mock.Anything)
	})

	t.Run("When calling ProcessMutant should create and detect dna person with error", func(t *testing.T) {

		repository := new(mocks.MockRespository)
		usecase := NewMutantUsecase(repository)

		repository.On("Save", mock.Anything).Return(errors.New("something went worng"))

		s := make([]string, 0)
		s = append(s, "ATGCGA")
		request := &model.RequestMutant{DNA: s}

		_, err := usecase.ProcessMutant(request)

		assert.Error(t, err)
		repository.AssertCalled(t, "Save", mock.Anything)
	})
}
