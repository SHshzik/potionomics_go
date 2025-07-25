package service

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/SHshzik/potionomics_go/domain/gen"
	"github.com/google/uuid"
	"github.com/tomcraven/goga"
)

type App struct {
	bdPotions              domain.BDPotions
	bdCauldrons            domain.BDCauldrons
	bdIngredients          domain.BDIngredients
	ingredientsInInventory []domain.InventoryCell
}

func NewApp(bdPotions domain.BDPotions, bdCauldrons domain.BDCauldrons, bdIngredients domain.BDIngredients) *App {
	return &App{
		bdPotions:     bdPotions,
		bdCauldrons:   bdCauldrons,
		bdIngredients: bdIngredients,
	}
}

func (s *App) GetIngredientsInInventory() []domain.InventoryCell {
	return GetIngredientsInInventory(s.bdIngredients)
}

func (s *App) GetIngredientsInShop() []domain.InventoryCell {
	return GetIngredientsInShop(s.bdIngredients)
}

func (s *App) Generate(r domain.GenerateRequest) []domain.BrewResult {
	s.ingredientsInInventory = s.GetIngredientsInInventory()

	resultChannel := make(chan []domain.Ingredient, 10)

	maxA, maxB, maxC, maxD, maxE := gen.CalculateMaxValues(r.Cauldron.Magmin, r.Potion.Proportions)
	minFitness := r.Cauldron.Magmin * 75 / 100

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	simulator := &gen.BrewSimulator{
		IngredientsInInventory: s.ingredientsInInventory,
		Capacity:               r.Cauldron.Capacity,
		Proportions:            r.Potion.Proportions,
		ResultChannel:          resultChannel,
		MaxA:                   maxA,
		MaxB:                   maxB,
		MaxC:                   maxC,
		MaxD:                   maxD,
		MaxE:                   maxE,
		MinFitness:             minFitness,
		Ctx:                    ctx,
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

func (s *App) GetPotion(id string) (domain.Potion, error) {
	potion, ok := s.bdPotions[id]
	if !ok {
		return domain.Potion{}, errors.New("potion not found")
	}
	return potion, nil
}

func (s *App) GetCauldron(id string) (domain.Cauldron, error) {
	cauldron, ok := s.bdCauldrons[id]
	if !ok {
		return domain.Cauldron{}, errors.New("cauldron not found")
	}
	return cauldron, nil
}

func (s *App) GetPotions() []domain.Potion {
	potions := make([]domain.Potion, 0, len(s.bdPotions))
	for _, potion := range s.bdPotions {
		potions = append(potions, potion)
	}
	sort.Slice(potions, func(i, j int) bool {
		return potions[i].Order < potions[j].Order
	})
	return potions
}

func (s *App) GetCauldrons() []domain.Cauldron {
	cauldrons := make([]domain.Cauldron, 0, len(s.bdCauldrons))
	for _, cauldron := range s.bdCauldrons {
		cauldrons = append(cauldrons, cauldron)
	}
	sort.Slice(cauldrons, func(i, j int) bool {
		return cauldrons[i].Order < cauldrons[j].Order
	})
	return cauldrons
}
