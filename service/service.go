package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SHshzik/potionomics_go/adapter/csv"
	"github.com/SHshzik/potionomics_go/adapter/save"
	"github.com/SHshzik/potionomics_go/domain"
)

func GetIngredients(filepath string) []domain.Ingredient {
	csvClient := csv.NewClient()
	saveClient := save.NewClient()
	csvIngredients := csvClient.ReadCsvFile("i.csv")
	allIngredients := make(map[string]domain.Ingredient, 250)

	for _, csvIngredient := range csvIngredients[1:] {
		baseName := csvIngredient[0]
		name := toLower(baseName)
		a, _ := strconv.Atoi(csvIngredient[1])
		b, _ := strconv.Atoi(csvIngredient[2])
		c, _ := strconv.Atoi(csvIngredient[3])
		d, _ := strconv.Atoi(csvIngredient[4])
		e, _ := strconv.Atoi(csvIngredient[5])
		allIngredients[name] = domain.Ingredient{
			Name: baseName,
			A:    a,
			B:    b,
			C:    c,
			D:    d,
			E:    e,
		}
	}
	saveIngredients := saveClient.FetchIngredientsInInventory(filepath)
	var length int
	for _, saveIngredient := range saveIngredients {
		length += int(saveIngredient.Count)
	}
	ingredients := make([]domain.Ingredient, 0, length)

	for _, saveIngredient := range saveIngredients {
		name := toLower(saveIngredient.Name)
		ing, ok := allIngredients[name]
		if !ok {
			fmt.Println("ERROR NAME - ", name)
			panic("ERROR NAME")
		}
		for i := 0; i < int(saveIngredient.Count); i++ {
			ingredients = append(ingredients, ing)
		}
	}
	return ingredients
}

func toLower(s string) string {
	r := strings.ReplaceAll(s, " ", "")
	r = strings.ReplaceAll(r, "'", "")
	r = strings.ReplaceAll(r, "-", "")
	r = strings.ToLower(r)
	return r
}
