package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github/meli/src/domain/model"
	"github/meli/src/mutant/mocks"
)

func TestGetStats(t *testing.T) {

	t.Run("When calling GetStats should return stats", func(t *testing.T) {

		repository := new(mocks.MockRespository)
		usecase := NewMutantUsecase(repository)

		repository.On("GetCountByFlag").Return(mocks.MockSliceStast(), nil)

		_, err := usecase.GetStats()

		assert.NoError(t, err)
		repository.AssertCalled(t, "GetCountByFlag")
	})

	t.Run("When calling GetStats with empty data should return stats", func(t *testing.T) {

		repository := new(mocks.MockRespository)
		usecase := NewMutantUsecase(repository)

		repository.On("GetCountByFlag").Return([]model.Mutant{}, nil)

		_, err := usecase.GetStats()

		assert.NoError(t, err)
		repository.AssertCalled(t, "GetCountByFlag")
	})

	t.Run("When calling GetStats should with error should return an error", func(t *testing.T) {

		repository := new(mocks.MockRespository)
		usecase := NewMutantUsecase(repository)

		repository.On("GetCountByFlag").Return(mocks.MockSliceStast(), errors.New("something went worng"))

		_, err := usecase.GetStats()

		assert.Error(t, err)
		repository.AssertCalled(t, "GetCountByFlag")
	})
}
