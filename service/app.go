package service

import (
	"strings"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/SHshzik/potionomics_go/domain/gen"
	"github.com/tomcraven/goga"
)

type App struct {
	GenAlgo          *goga.GeneticAlgorithm
	EliteConsumer    *gen.EliteConsumer
	BrewSimulator    *gen.BrewSimulator
	PotionsRecords   [][]string
	CauldronsRecords [][]string
}

func NewApp(genAlgo *goga.GeneticAlgorithm, eliteConsumer *gen.EliteConsumer, brewSimulator *gen.BrewSimulator, potionsRecords [][]string, cauldronsRecords [][]string) *App {
	return &App{
		GenAlgo:          genAlgo,
		EliteConsumer:    eliteConsumer,
		BrewSimulator:    brewSimulator,
		PotionsRecords:   potionsRecords,
		CauldronsRecords: cauldronsRecords,
	}
}

func (s *App) Generate(r domain.GenerateRequest) string {
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

func (s *App) GetPotions() []domain.Potion {
	potions := make([]domain.Potion, 0, len(s.PotionsRecords))
	for i, record := range s.PotionsRecords[1:] {
		potions = append(potions, domain.Potion{ID: i, Name: record[0]})
	}
	return potions
}

func (s *App) GetCauldrons() []domain.Cauldron {
	cauldrons := make([]domain.Cauldron, 0, len(s.CauldronsRecords))
	for i, record := range s.CauldronsRecords[1:] {
		cauldrons = append(cauldrons, domain.Cauldron{ID: i, Name: record[0]})
	}
	return cauldrons
}
