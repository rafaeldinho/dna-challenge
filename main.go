package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	healthHandler "github/meli/src/health/handler"
	healthUseCase "github/meli/src/health/usecase"

	"github/meli/src/infrastructure/mongo"
	envs "github/meli/src/infrastructure/viper"
	"github/meli/src/infrastructure/web"
	mutantHandler "github/meli/src/mutant/handler"
	"github/meli/src/mutant/repository"
	mutantUsecase "github/meli/src/mutant/usecase"
	"github/meli/src/shared"
)

var logger = log.WithFields(log.Fields{
	"layer": shared.MainLayer,
})

func init() {
	logger.Info("Staring app...")
	envs.SetValuesOnEnvironment()
	envs.PrintEnvs()
}

func main() {

	app := web.ServerInstance()
	mongoInstance := mongo.CreateMongoClient()

	// HEALTH SERVICES
	healthHandler.NewHealthHandler(app, healthUseCase.NewHealthUseCase())

	// MUTANT SERVICES
	mutantRepository := repository.NewMutantRepository(mongoInstance)
	mutantService := mutantUsecase.NewMutantUsecase(mutantRepository)
	mutantHandler.NewMutantHandler(app, mutantService)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", port)))
}
