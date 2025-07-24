package service

import (
	"github.com/SHshzik/potionomics_go/domain"
	"github.com/SHshzik/potionomics_go/domain/gen"
	"github.com/google/uuid"
	"github.com/tomcraven/goga"
)

type App struct {
	bdPotions              domain.BDPotions
	bdCauldrons            domain.BDCauldrons
	bdIngredients          domain.BDIngredients
	ingredientsInInventory []domain.Ingredient
}

func NewApp(bdPotions domain.BDPotions, bdCauldrons domain.BDCauldrons, bdIngredients domain.BDIngredients, ingredientsInInventory []domain.Ingredient) *App {
	return &App{
		bdPotions:              bdPotions,
		bdCauldrons:            bdCauldrons,
		bdIngredients:          bdIngredients,
		ingredientsInInventory: ingredientsInInventory,
	}
}

func (s *App) Generate(r domain.GenerateRequest) []domain.BrewResult {
	resultChannel := make(chan []domain.Ingredient, 10)

	simulator := &gen.BrewSimulator{
		IngredientsInInventory: s.ingredientsInInventory,
		Capacity:               r.Cauldron.Capacity,
		ResultChannel:          resultChannel,
	}
	creator := &gen.BitsetCreator{
		IngredientsInInventory: s.ingredientsInInventory,
		Capacity:               r.Cauldron.Capacity,
	}
	eliteConsumer := &gen.EliteConsumer{
		IngredientsInInventory: s.ingredientsInInventory,
		ResultChannel:          resultChannel,
	}

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

	// Move in gui input?
	genAlgo.Init(50000, 4)
	go func() {
		genAlgo.Simulate()
	}()

	top := make([]domain.BrewResult, 0, 10)

	for r := range resultChannel {
		if len(top) > 10 {
			top = top[1:]
		}
		top = append(top, domain.BrewResult{
			ID:      uuid.New().String(),
			Receipt: r,
		})
	}

	return top
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
