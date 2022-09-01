package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github/meli/src/domain/model"
	"github/meli/src/mutant/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMutantHandler(t *testing.T) {
	t.Parallel()
	t.Run("when calling stats should return stats object", func(t *testing.T) {

		useCase := new(mocks.MockUseCase)
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/mutant/stats", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

		useCase.On("GetStats").Return(mocks.MockObjectStats(), nil)

		healthCheck := NewMutantHandler(e, useCase)

		_ = healthCheck.GetStats(echoContext)

		assert.Equal(t, http.StatusOK, rec.Code)
		useCase.AssertCalled(t, "GetStats")
	})

	t.Run("when checking is dna IsMutant should return status ok", func(t *testing.T) {

		useCase := new(mocks.MockUseCase)
		e := echo.New()

		data, _ := json.Marshal(mocks.MockRequest())
		req := httptest.NewRequest(http.MethodPost, "/mutant/stats", bytes.NewReader(data))
		req.Header.Set("Content-Type", echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

		useCase.On("ProcessMutant", mock.Anything).Return(mocks.MockMutant(true), nil)

		healthCheck := NewMutantHandler(e, useCase)

		_ = healthCheck.IsMutant(echoContext)

		assert.Equal(t, http.StatusOK, rec.Code)
		useCase.AssertCalled(t, "ProcessMutant", mock.Anything)
	})

	t.Run("when checking is dna Is not Mutant should return an forbidden status", func(t *testing.T) {

		useCase := new(mocks.MockUseCase)
		e := echo.New()

		data, _ := json.Marshal(mocks.MockRequest())
		req := httptest.NewRequest(http.MethodPost, "/mutant/stats", bytes.NewReader(data))
		req.Header.Set("Content-Type", echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

		useCase.On("ProcessMutant", mock.Anything).Return(mocks.MockMutant(false), nil)

		healthCheck := NewMutantHandler(e, useCase)

		_ = healthCheck.IsMutant(echoContext)

		assert.Equal(t, http.StatusForbidden, rec.Code)
		useCase.AssertCalled(t, "ProcessMutant", mock.Anything)
	})
}

func TestMutantHandlerEror(t *testing.T) {
	t.Parallel()
	t.Run("when calling stats with erros should return an error", func(t *testing.T) {

		useCase := new(mocks.MockUseCase)
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/mutant/stats", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

		useCase.On("GetStats").Return(model.Stats{}, errors.New("something went wront"))

		healthCheck := NewMutantHandler(e, useCase)

		_ = healthCheck.GetStats(echoContext)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		useCase.AssertCalled(t, "GetStats")
	})

	t.Run("when checking is dna Is not Mutant should return an error", func(t *testing.T) {

		useCase := new(mocks.MockUseCase)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/mutant/stats", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

		useCase.On("ProcessMutant", mock.Anything).Return(mocks.MockMutant(false), nil)

		healthCheck := NewMutantHandler(e, useCase)

		_ = healthCheck.IsMutant(echoContext)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("when checking is dna Is not Mutant should return an error", func(t *testing.T) {

		useCase := new(mocks.MockUseCase)
		e := echo.New()

		data, _ := json.Marshal(mocks.MockRequest())
		req := httptest.NewRequest(http.MethodPost, "/mutant/stats", bytes.NewReader(data))
		req.Header.Set("Content-Type", echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

		useCase.On("ProcessMutant", mock.Anything).Return(mocks.MockMutant(false), errors.New("something went worng"))

		healthCheck := NewMutantHandler(e, useCase)

		_ = healthCheck.IsMutant(echoContext)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}
