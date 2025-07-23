package main

import (
	"fmt"
	"runtime"

	"github.com/SHshzik/potionomics_go/adapter/csv"
	"github.com/SHshzik/potionomics_go/adapter/save"
	"github.com/SHshzik/potionomics_go/service"
)

func main() {
	numThreads := 4
	runtime.GOMAXPROCS(numThreads)

	// ingredient

	csvClient := csv.NewClient()
	saveClient := save.NewClient()
	service := service.NewService(csvClient, saveClient, "PotionomicsSaveData8.sav")

	ingredients := service.GetIngredients()

	fmt.Println(ingredients)

	// records := csvClient.ReadCsvFile("i.csv")
	// fmt.Println(len(records))
	// records := readCsvFile("i.csv")

	// ingredients := make([]domain.Ingredient, 0, 207)
	// for _, record := range records {
	// 	a, _ := strconv.Atoi(record[1])
	// 	b, _ := strconv.Atoi(record[2])
	// 	c, _ := strconv.Atoi(record[3])
	// 	d, _ := strconv.Atoi(record[4])
	// 	e, _ := strconv.Atoi(record[5])
	// 	ingredients = append(ingredients, domain.Ingredient{
	// 		Name: record[0],
	// 		A:    a,
	// 		B:    b,
	// 		C:    c,
	// 		D:    d,
	// 		E:    e,
	// 	})
	// }
	// ingredients = domain.Filter(ingredients, domain.AvailableIngredients)

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
	// 		{P: 1.0, F: goga.UniformCrossover, UseElite: false},
	// 	},
	// )

	// // Maybe change.
	// genAlgo.Selector = goga.NewSelector(
	// 	[]goga.SelectorFunctionProbability{
	// 		{P: 1.0, F: goga.Roulette},
	// 	},
	// )

	// genAlgo.Init(600, 4)

	// startTime := time.Now()
	// genAlgo.Simulate()
	// fmt.Println(time.Since(startTime))
}
