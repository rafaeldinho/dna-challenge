package usecase

import (
	"github/meli/src/domain/model"
	"github/meli/src/mutant/repository"
	"github/meli/src/shared"

	log "github.com/sirupsen/logrus"
)

type MutantUsecase interface {
	IsMutant(dna *model.RequestMutant) (*model.Mutant, error)
	GetStats() (model.Stats, error)
}

type mutantUsecase struct {
	repository repository.MutantRepository
}

var logger = log.WithFields(log.Fields{
	"layer": shared.Mutant,
})

func NewMutantUsecase(repository repository.MutantRepository) *mutantUsecase {
	return &mutantUsecase{repository}
}
