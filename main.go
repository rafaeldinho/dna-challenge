package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	healthHandler "github/meli/src/health/handler"
	healthUseCase "github/meli/src/health/usecase"
	"github/meli/src/infrastructure/firestore"
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
	godotenv.Load("app.env")
	envs.SetValuesOnEnvironment()
	envs.PrintEnvs()
}

func main() {

	app := web.ServerInstance()
	fireStoreInstance := firestore.CreateClient()

	mutantRepository := repository.NewMutantRepository(fireStoreInstance)

	healthHandler.NewHealthHandler(app, healthUseCase.NewHealthUseCase())

	adasd := mutantUsecase.NewMutantUsecase(mutantRepository)

	mutantHandler.NewMutantHandler(app, adasd)

	port := viper.GetString("PORT")

	if port == "" {
		port = ":8080"
	} else {
		port = fmt.Sprintf(":%s", viper.GetString("PORT"))
	}

	server := &http.Server{
		Addr:         port,
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	app.Logger.Fatal(app.StartServer(server))
}
