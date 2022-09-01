package usecase

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github/meli/src/domain/model"
	"github/meli/src/shared"
)

var matchMap = map[string][]string{}
var result = make([]string, 0)

func (m *mutantUsecase) ProcessMutant(dna *model.RequestMutant) (*model.Mutant, error) {
	logger.Info("processing mutant")
	mutant := &model.Mutant{
		DNA:      strings.Join(dna.DNA[:], "-"),
		IsMutant: isMutant(dna.DNA),
	}

	if err := m.repository.Save(mutant); err != nil {
		return &model.Mutant{}, err
	}

	cleanMap()
	cleanSliceResult()
	logger.WithFields(log.Fields{"DNA": mutant.DNA, "isMutant": mutant.IsMutant}).Info("DNA saved sucess")
	return mutant, nil
}

func cleanMap() {
	for k := range matchMap {
		delete(matchMap, k)
	}
}

func cleanSliceResult() {
	result = nil
}
func setValue(key string) {
	match := make([]string, 0)

	if value, ok := matchMap[key]; ok {
		match = append(match, value...)
	} else {
		cleanMap()
	}

	match = append(match, key)
	matchMap[key] = match
	if len(match) >= 4 {
		result = append(result, match...)
	}
}
func isMutant(DNA []string) bool {
	dnaMatris := make([]string, 0)
	for _, v := range DNA {
		s := make([]string, 0)
		for i := 0; i < len(v); i++ {
			s = append(s, string(v[i]))
		}
		dnaMatris = append(dnaMatris, s...)
		logger.Infoln(s)
	}

	for _, k := range dnaMatris {
		if _, ok := shared.DNAMAP[k]; ok {
			setValue(k)
		}
	}

	logger.Infoln((result))
	return len(result) > 0
}
