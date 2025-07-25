package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SHshzik/potionomics_go/adapter/save"
	"github.com/SHshzik/potionomics_go/domain"
)

func GetBDIngredients(ingredientsRecords [][]string) domain.BDIngredients {
	allIngredients := make(domain.BDIngredients, len(ingredientsRecords))

	for _, ingredientRecord := range ingredientsRecords[1:] {
		baseName := ingredientRecord[0]
		name := toLower(baseName)
		a, _ := strconv.Atoi(ingredientRecord[1])
		b, _ := strconv.Atoi(ingredientRecord[2])
		c, _ := strconv.Atoi(ingredientRecord[3])
		d, _ := strconv.Atoi(ingredientRecord[4])
		e, _ := strconv.Atoi(ingredientRecord[5])

		allIngredients[name] = domain.Ingredient{
			Name:     baseName,
			Translit: ingredientRecord[42],
			A:        a,
			B:        b,
			C:        c,
			D:        d,
			E:        e,
		}
	}

	return allIngredients
}

func GetIngredientsInInventory(BDIngredients domain.BDIngredients) []domain.InventoryCell {
	filepath, _ := FindLastUpdated()
	fmt.Println(filepath)
	saveIngredients := save.FetchIngredientsInInventory(filepath)
	var length int
	for _, saveIngredient := range saveIngredients {
		length += int(saveIngredient.Count)
	}
	ingredients := make([]domain.InventoryCell, 0, length)

	for _, saveIngredient := range saveIngredients {
		name := toLower(saveIngredient.Name)
		ing, ok := BDIngredients[name]
		if !ok {
			fmt.Println("ERROR NAME - ", name)
			panic("ERROR NAME")
		}
		for i := 0; i < int(saveIngredient.Count); i++ {
			ingredients = append(ingredients, domain.InventoryCell{
				Ingredient: ing,
				CellNumber: i,
			})
		}
	}
	return ingredients
}

func GetIngredientsInShop(BDIngredients domain.BDIngredients) []domain.InventoryCell {
	filepath, _ := FindLastUpdated()
	fmt.Println(filepath)
	saveIngredients := save.FetchIngredientsInShop(filepath)
	var length int
	for _, saveIngredient := range saveIngredients {
		length += int(saveIngredient.Count)
	}
	ingredients := make([]domain.InventoryCell, 0, length)

	for _, saveIngredient := range saveIngredients {
		name := toLower(saveIngredient.Name)
		ing, ok := BDIngredients[name]
		if !ok {
			fmt.Println("ERROR NAME - ", saveIngredient)
			panic("ERROR NAME")
		}
		for i := 0; i < int(saveIngredient.Count); i++ {
			ingredients = append(ingredients, domain.InventoryCell{
				Ingredient: ing,
				CellNumber: i,
			})
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
