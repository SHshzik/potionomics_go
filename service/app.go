package service

import (
	"strings"

	"github.com/SHshzik/potionomics_go/domain/gen"
	"github.com/tomcraven/goga"
)

type App struct {
	GenAlgo       *goga.GeneticAlgorithm
	EliteConsumer *gen.EliteConsumer
	BrewSimulator *gen.BrewSimulator
}

func NewApp(genAlgo *goga.GeneticAlgorithm, eliteConsumer *gen.EliteConsumer, brewSimulator *gen.BrewSimulator) *App {
	return &App{
		GenAlgo:       genAlgo,
		EliteConsumer: eliteConsumer,
		BrewSimulator: brewSimulator,
	}
}

func (s *App) Generate() string {
	resultChannel := make(chan string, 10)
	s.EliteConsumer.ResultChannel = resultChannel
	s.BrewSimulator.ResultChannel = resultChannel
	// Move in gui input?
	s.GenAlgo.Init(50000, 4)
	s.GenAlgo.Simulate()

	top := make([]string, 0, 10)

	for r := range resultChannel {
		if len(top) > 10 {
			top = top[1:]
		}
		top = append(top, r)
	}
	sr := strings.Builder{}
	for _, t := range top {
		sr.WriteString(t)
		sr.WriteString("<br>")
	}

	return sr.String()
}
