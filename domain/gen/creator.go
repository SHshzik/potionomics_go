package gen

import (
	"math/rand"
	"slices"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/tomcraven/goga"
)

type BitsetCreator struct {
	Ingredients []domain.Ingredient
	Capacity    int
}

func (bc *BitsetCreator) Go() goga.Bitset {
	b := goga.Bitset{}
	b.Create(len(bc.Ingredients))
	randomSelectedRows := make([]domain.Ingredient, 0, bc.Capacity)
	for i := 0; i < bc.Capacity; i++ {
		randomIndex := rand.Intn(len(bc.Ingredients))
		randomSelectedRows = append(randomSelectedRows, bc.Ingredients[randomIndex])
	}
	for i := 0; i < len(bc.Ingredients); i++ {
		result := slices.ContainsFunc(randomSelectedRows, func(r domain.Ingredient) bool {
			return r == bc.Ingredients[i]
		})
		var r int
		if result {
			r = 1
		} else {
			r = 0
		}
		b.Set(i, r)
	}
	return b
}
