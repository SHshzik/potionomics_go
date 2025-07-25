package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/SHshzik/potionomics_go/adapter/csv"
	"github.com/SHshzik/potionomics_go/domain"
	"github.com/SHshzik/potionomics_go/handlers"
	"github.com/SHshzik/potionomics_go/pkg/httpserver"
	"github.com/SHshzik/potionomics_go/pkg/logger"
	"github.com/SHshzik/potionomics_go/service"
	"github.com/gofiber/fiber/v2"
)

//go:embed all:data
var data embed.FS

func main() {
	numThreads := 4
	runtime.GOMAXPROCS(numThreads)

	// TODO: use config
	l := logger.New("debug")

	var bdPotions domain.BDPotions
	{
		potionsRecords := csv.ReadCsvFile(data, "data/potions.csv")
		bdPotions = service.GetBDPotions(potionsRecords)
	}

	var bdCauldrons domain.BDCauldrons
	{
		cauldronsRecords := csv.ReadCsvFile(data, "data/cauldrons.csv")
		bdCauldrons = service.GetBDCauldrons(cauldronsRecords)
	}

	var bdIngredients domain.BDIngredients
	{
		ingredientsRecords := csv.ReadCsvFile(data, "data/ingredients.csv")
		bdIngredients = service.GetBDIngredients(ingredientsRecords)
	}

	// TODO: use not path, last updated file from directory.
	ingredientsInInventory := service.GetIngredientsInInventory("./PotionomicsSaveData9.sav", bdIngredients)

	server := httpserver.New(httpserver.Port("8080"))

	{
		server.App.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })
	}

	{
		app := service.NewApp(bdPotions, bdCauldrons, bdIngredients, ingredientsInInventory)
		myServer := handlers.NewHTTPServer(app, l)
		v1Ground := server.App.Group("")
		v1Ground.Get("/get_potions", myServer.GetPotions)
		v1Ground.Get("/get_cauldrons", myServer.GetCauldrons)
		v1Ground.Post("/generate", myServer.Generate)
	}

	server.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-server.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := server.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
