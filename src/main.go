package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	healthHandler "github/meli/src/health/handler"
	mutantHandler "github/meli/src/mutant/handler"
	healthUseCase "github/meli/src/health/usecase"
	mutantUsecase "github/meli/src/mutant/usecase"
	"github/meli/src/infrastructure/sql"
	"github/meli/src/infrastructure/web"
	"github/meli/src/mutant/repository"
	"github/meli/src/shared"
)

var logger = log.WithFields(log.Fields{
	"layer": shared.MainLayer,
})

func init () {
	logger.Info("Staring app...")
}

func main() {

	sqlInstance := sql.NewSqlInstance("root" + ":" + "uCvyM0$SCMmj5R`H" + "@tcp" + "(" + "35.237.33.131" + ":" + "3306" + ")/" + "mutants" + "?" + "parseTime=true&loc=Local")
	sqlInstance.Migration()

	app := web.ServerInstance()

   mutantRepository := repository.NewMutantRepository(sqlInstance)
   
	healthHandler.NewHealthHandler(app, healthUseCase.NewHealthUseCase())

	adasd := mutantUsecase.NewMutantUsecase(mutantRepository)

	mutantHandler.NewMutantHandler(app,adasd)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	app.Logger.Fatal(app.StartServer(server))
}
