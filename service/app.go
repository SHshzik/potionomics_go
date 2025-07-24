package service

import (
	"github.com/SHshzik/potionomics_go/domain"
)

type App struct {
	bdPotions     domain.BDPotions
	bdCauldrons   domain.BDCauldrons
	bdIngredients domain.BDIngredients
	// GenAlgo          *goga.GeneticAlgorithm
	// EliteConsumer    *gen.EliteConsumer
	// BrewSimulator    *gen.BrewSimulator
	// PotionsRecords   [][]string
	// CauldronsRecords [][]string
	// Ingredients      []domain.Ingredient
}

func NewApp(bdPotions domain.BDPotions, bdCauldrons domain.BDCauldrons, bdIngredients domain.BDIngredients) *App {
	return &App{
		bdPotions:     bdPotions,
		bdCauldrons:   bdCauldrons,
		bdIngredients: bdIngredients,
	}
}

func (s *App) Generate(r domain.GenerateRequest) [][]domain.Ingredient {
	return [][]domain.Ingredient{}
	// resultChannel := make(chan string, 10)
	// s.EliteConsumer.ResultChannel = resultChannel
	// s.BrewSimulator.ResultChannel = resultChannel
	// // Move in gui input?
	// s.GenAlgo.Init(50000, 4)
	// s.GenAlgo.Simulate()

	// top := make([]domain.Ingredient, 0, 10)

	// for r := range resultChannel {
	// 	if len(top) > 10 {
	// 		top = top[1:]
	// 	}
	// 	top = append(top, s.Ingredients[r])
	// }
	// sr := strings.Builder{}
	// for _, t := range top {
	// 	sr.WriteString(t)
	// 	sr.WriteString("<br>")
	// }

	// return sr.String()
}

func (s *App) GetPotions() []domain.Potion {
	potions := make([]domain.Potion, 0, len(s.bdPotions))
	for _, potion := range s.bdPotions {
		potions = append(potions, potion)
	}
	return potions
}

func (s *App) GetCauldrons() []domain.Cauldron {
	cauldrons := make([]domain.Cauldron, 0, len(s.bdCauldrons))
	for _, cauldron := range s.bdCauldrons {
		cauldrons = append(cauldrons, cauldron)
	}
	return cauldrons
}
