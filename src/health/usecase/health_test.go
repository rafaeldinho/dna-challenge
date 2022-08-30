package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github/meli/src/health/mocks"
)

func TestNewHealthUseCase(t *testing.T) {

	t.Run("When calling GetCheck should return health status object", func(t *testing.T) {

		healthCheck := NewHealthUseCase()

		resutl := healthCheck.GetCheck()
		mockResult := mocks.MockHealthObject()

		assert.Equal(t, resutl, mockResult)
	})
}
