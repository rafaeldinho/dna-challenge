package usecase

import (
	"github/meli/src/domain/model"
	"github/meli/src/shared"

	log "github.com/sirupsen/logrus"
)

type HealthUseCase interface {
	GetCheck() model.Health
}

type healthUseCase struct{}

var logger = log.WithFields(log.Fields{
	"layer": shared.HealthLayer,
})

func NewHealthUseCase() *healthUseCase {
	return &healthUseCase{}
}

func (h *healthUseCase) GetCheck() model.Health {
	logger.Info("getting service status")
	return model.Health{
		Status:  "UP",
		Version: "1.0.0",
	}
}
