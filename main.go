package main

import (
	"embed"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// numThreads := 4
	// runtime.GOMAXPROCS(numThreads)

	// // ingredient

	// csvClient := csv.NewClient()
	// saveClient := save.NewClient()
	// service := service.NewService(csvClient, saveClient, "./PotionomicsSaveData9.sav")

	// ingredients := service.GetIngredients()

	// simulator := &gen.BrewSimulator{Ingredients: ingredients, Capacity: 6}
	// creator := &gen.BitsetCreator{Ingredients: ingredients, Capacity: 6}
	// eliteConsumer := &gen.EliteConsumer{Ingredients: ingredients}

	// genAlgo := goga.NewGeneticAlgorithm()

	// genAlgo.Simulator = simulator
	// genAlgo.BitsetCreate = creator
	// genAlgo.EliteConsumer = eliteConsumer

	// // Maybe change.
	// genAlgo.Mater = goga.NewMater(
	// 	[]goga.MaterFunctionProbability{
	// 		{P: 1.0, F: goga.TwoPointCrossover},
	// 		{P: 1.0, F: goga.Mutate},
	// 		{P: 1.0, F: goga.UniformCrossover, UseElite: true},
	// 	},
	// )

	// // Maybe change.
	// genAlgo.Selector = goga.NewSelector(
	// 	[]goga.SelectorFunctionProbability{
	// 		{P: 1.0, F: goga.Roulette},
	// 	},
	// )

	// genAlgo.Init(50000, 4)

	// startTime := time.Now()
	// genAlgo.Simulate()
	// fmt.Println(time.Since(startTime))

	// Create an instance of the app structure
	app := domain.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "potionomics_go",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
