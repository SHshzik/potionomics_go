package main

import (
	"embed"
	"runtime"

	"github.com/SHshzik/potionomics_go/adapter/csv"
	"github.com/SHshzik/potionomics_go/domain"
	"github.com/SHshzik/potionomics_go/service"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed all:data
var data embed.FS

func main() {
	numThreads := 4
	runtime.GOMAXPROCS(numThreads)

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
	_ = ingredientsInInventory

	// simulator := &gen.BrewSimulator{Ingredients: ingredientsInInventory}
	// creator := &gen.BitsetCreator{Ingredients: ingredientsInInventory}
	// eliteConsumer := &gen.EliteConsumer{Ingredients: ingredientsInInventory}

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
	app := service.NewApp(bdPotions, bdCauldrons, bdIngredients)

	// Create an instance of the app structure
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "potionomics_go",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		// OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
