package service

import (
	"fmt"

	"github.com/SHshzik/potionomics_go/adapter/csv"
	"github.com/SHshzik/potionomics_go/adapter/save"
	"github.com/SHshzik/potionomics_go/domain"
)

type Service struct {
	csvClient    *csv.Client
	saveClient   *save.Client
	saveFilePath string
}

func NewService(csvClient *csv.Client, saveClient *save.Client, filepath string) *Service {
	return &Service{
		csvClient:    csvClient,
		saveClient:   saveClient,
		saveFilePath: filepath,
	}
}

func (s *Service) GetIngredients() []domain.Ingredient {
	saveIngredientsInInventory := s.saveClient.FetchIngredientsInInventory(s.saveFilePath)
	var length int
	for _, saveIngredient := range saveIngredientsInInventory {
		length += int(saveIngredient.Count)
	}
	ingredients := make([]domain.Ingredient, 0, length)
	// TODO: to domain
	_ = ingredients
	fmt.Println(saveIngredientsInInventory)
	return []domain.Ingredient{}
}
