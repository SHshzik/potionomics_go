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
		basePrice, _ := strconv.Atoi(ingredientRecord[9])
		perfectForHealtPotion, _ := strconv.ParseBool(ingredientRecord[22])
		perfectForManaPotion, _ := strconv.ParseBool(ingredientRecord[23])
		perfectForStaminaPotion, _ := strconv.ParseBool(ingredientRecord[24])
		perfectForSpeedPotion, _ := strconv.ParseBool(ingredientRecord[25])
		perfectForTolerancePotion, _ := strconv.ParseBool(ingredientRecord[26])
		perfectForFireTonic, _ := strconv.ParseBool(ingredientRecord[27])
		perfectForIceTonic, _ := strconv.ParseBool(ingredientRecord[28])
		perfectForThunderTonic, _ := strconv.ParseBool(ingredientRecord[29])
		perfectForShadowTonic, _ := strconv.ParseBool(ingredientRecord[30])
		perfectForRadiationTonic, _ := strconv.ParseBool(ingredientRecord[31])
		perfectForSightEnhancer, _ := strconv.ParseBool(ingredientRecord[32])
		perfectForAlertnessEnhancer, _ := strconv.ParseBool(ingredientRecord[33])
		perfectForInsightEnhancer, _ := strconv.ParseBool(ingredientRecord[34])
		perfectForDowsingEnhancer, _ := strconv.ParseBool(ingredientRecord[35])
		perfectForSeekingEnhancer, _ := strconv.ParseBool(ingredientRecord[36])
		perfectForPoisonCure, _ := strconv.ParseBool(ingredientRecord[37])
		perfectForDrowsinessCure, _ := strconv.ParseBool(ingredientRecord[38])
		perfectForPetrificationCure, _ := strconv.ParseBool(ingredientRecord[39])
		perfectForSilenceCure, _ := strconv.ParseBool(ingredientRecord[40])
		perfectForCurseCure, _ := strconv.ParseBool(ingredientRecord[41])

		allIngredients[name] = domain.Ingredient{
			Name:      baseName,
			Translit:  ingredientRecord[42],
			A:         a,
			B:         b,
			C:         c,
			D:         d,
			E:         e,
			BasePrice: basePrice,

			PerfectForHealtPotion:       perfectForHealtPotion,
			PerfectForManaPotion:        perfectForManaPotion,
			PerfectForStaminaPotion:     perfectForStaminaPotion,
			PerfectForSpeedPotion:       perfectForSpeedPotion,
			PerfectForTolerancePotion:   perfectForTolerancePotion,
			PerfectForFireTonic:         perfectForFireTonic,
			PerfectForIceTonic:          perfectForIceTonic,
			PerfectForThunderTonic:      perfectForThunderTonic,
			PerfectForShadowTonic:       perfectForShadowTonic,
			PerfectForRadiationTonic:    perfectForRadiationTonic,
			PerfectForSightEnhancer:     perfectForSightEnhancer,
			PerfectForAlertnessEnhancer: perfectForAlertnessEnhancer,
			PerfectForInsightEnhancer:   perfectForInsightEnhancer,
			PerfectForDowsingEnhancer:   perfectForDowsingEnhancer,
			PerfectForSeekingEnhancer:   perfectForSeekingEnhancer,
			PerfectForPoisonCure:        perfectForPoisonCure,
			PerfectForDrowsinessCure:    perfectForDrowsinessCure,
			PerfectForPetrificationCure: perfectForPetrificationCure,
			PerfectForSilenceCure:       perfectForSilenceCure,
			PerfectForCurseCure:         perfectForCurseCure,
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

func Optimizate(ingredients []domain.InventoryCell, potion domain.Potion) []domain.InventoryCell {
	result := make([]domain.InventoryCell, 0, len(ingredients))
	for _, ing := range ingredients {
		if ing.Ingredient.PerfectForPotion(potion) {
			result = append(result, ing)
		}
	}

	return result
}

func WithShop(inv, shop []domain.InventoryCell) []domain.InventoryCell {
	return append(inv, shop...)
}

func UpdateCellNumber(inv []domain.InventoryCell) {
	for i, ing := range inv {
		ing.CellNumber = i
	}
}
