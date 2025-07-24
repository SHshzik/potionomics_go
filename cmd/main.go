package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/SHshzik/potionomics_go/adapter/csv"
	"github.com/SHshzik/potionomics_go/adapter/save"
	"github.com/SHshzik/potionomics_go/domain/gen"
	"github.com/SHshzik/potionomics_go/service"
	"github.com/tomcraven/goga"
)

func main() {
	numThreads := 4
	runtime.GOMAXPROCS(numThreads)

	// ingredient

	csvClient := csv.NewClient()
	saveClient := save.NewClient()
	service := service.NewService(csvClient, saveClient, "./PotionomicsSaveData9.sav")

	ingredients := service.GetIngredients()

	simulator := &gen.BrewSimulator{Ingredients: ingredients, Capacity: 6}
	creator := &gen.BitsetCreator{Ingredients: ingredients, Capacity: 6}
	eliteConsumer := &gen.EliteConsumer{Ingredients: ingredients}

	genAlgo := goga.NewGeneticAlgorithm()

	genAlgo.Simulator = simulator
	genAlgo.BitsetCreate = creator
	genAlgo.EliteConsumer = eliteConsumer

	// Maybe change.
	genAlgo.Mater = goga.NewMater(
		[]goga.MaterFunctionProbability{
			{P: 1.0, F: goga.TwoPointCrossover},
			{P: 1.0, F: goga.Mutate},
			{P: 1.0, F: goga.UniformCrossover, UseElite: true},
		},
	)

	// Maybe change.
	genAlgo.Selector = goga.NewSelector(
		[]goga.SelectorFunctionProbability{
			{P: 1.0, F: goga.Roulette},
		},
	)

	genAlgo.Init(50000, 4)

	startTime := time.Now()
	genAlgo.Simulate()
	fmt.Println(time.Since(startTime))
}
