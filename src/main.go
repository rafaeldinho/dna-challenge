package main

import (
	"github/meli/src/shared/entity"

	log "github.com/sirupsen/logrus"
)


var logger = log.WithFields(log.Fields{
	"layer": entity.MainLayer,
})

func main()  {
	logger.Info("Staring app...")
}