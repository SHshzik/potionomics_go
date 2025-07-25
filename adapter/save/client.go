package save

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	startIngredientsInInventoryB = []byte("IngredientsInInventory")
	startPotionsInInventoryB     = []byte("PotionsInInventory")
	ingredientCollectionB        = []byte("IngredientCollection")
	ingredintB                   = append([]byte("Ingredient"), 0x00)
	countB                       = []byte("Count")
)

func FetchIngredientsInInventory(filepath string) []Ingredient {
	data, err := ioutil.ReadFile(filepath) // "PotionomicsSaveData11.sav"
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		os.Exit(1)
	}

	startIngredientsInInventoryOffset := bytes.Index(data, startIngredientsInInventoryB)
	idx := bytes.Index(data[startIngredientsInInventoryOffset+1:], ingredientCollectionB)
	offset := startIngredientsInInventoryOffset + idx + 2
	startPotionsInInventoryOffset := bytes.Index(data, startPotionsInInventoryB)
	ingredients := make([]Ingredient, 0)
	for {
		idx := bytes.Index(data[offset:], ingredintB)
		if idx == -1 {
			break
		}
		if offset+idx > startPotionsInInventoryOffset {
			break
		}

		ingredient := Ingredient{}
		startNameOffset := offset + idx + 117
		nameBytes := make([]byte, 0)

		for i := startNameOffset; i < len(data); i++ {
			if data[i] == 0 {
				break
			}
			nameBytes = append(nameBytes, data[i])
		}
		ingredient.Name = string(nameBytes)
		ingredient.NameOffset = startNameOffset

		idx = bytes.Index(data[offset:], countB)
		if idx == -1 {
			break
		}
		if offset+idx > startPotionsInInventoryOffset {
			break
		}

		count := binary.LittleEndian.Uint16(data[offset+idx+31 : offset+idx+31+2])
		ingredient.Count = count
		ingredient.CountOffset = offset + idx + 31
		ingredients = append(ingredients, ingredient)
		offset += idx + 1
	}

	return ingredients
}

var (
	startIngredientsInShopB = []byte("DailyIngredients")
	startShopCauldronsB     = []byte("ShopCauldrons")
)

func FetchIngredientsInShop(filepath string) []Ingredient {
	data, err := ioutil.ReadFile(filepath) // "PotionomicsSaveData11.sav"
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		os.Exit(1)
	}

	startIngredientsInShopOffset := bytes.Index(data, startIngredientsInShopB)
	idx := bytes.Index(data[startIngredientsInShopOffset+1:], ingredientCollectionB)
	offset := startIngredientsInShopOffset + idx + 2
	startShopCauldronsOffset := bytes.Index(data, startShopCauldronsB)
	ingredients := make([]Ingredient, 0)
	for {
		idx := bytes.Index(data[offset:], ingredintB)
		if idx == -1 {
			break
		}
		if offset+idx > startShopCauldronsOffset {
			break
		}

		ingredient := Ingredient{Shop: true}
		startNameOffset := offset + idx + 117
		nameBytes := make([]byte, 0)

		for i := startNameOffset; i < len(data); i++ {
			if data[i] == 0 {
				break
			}
			nameBytes = append(nameBytes, data[i])
		}
		ingredient.Name = string(nameBytes)
		ingredient.NameOffset = startNameOffset

		idx = bytes.Index(data[offset:], countB)
		if idx == -1 {
			break
		}
		if offset+idx > startShopCauldronsOffset {
			break
		}

		count := binary.LittleEndian.Uint16(data[offset+idx+31 : offset+idx+31+2])
		ingredient.Count = count
		ingredient.CountOffset = offset + idx + 31
		ingredients = append(ingredients, ingredient)
		offset += idx + 1
	}

	return ingredients
}
