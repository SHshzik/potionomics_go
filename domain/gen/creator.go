package gen

import (
	"math/rand"
	"slices"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/tomcraven/goga"
)

type BitsetCreator struct {
	IngredientsInInventory []domain.InventoryCell
	Capacity               int
}

func (bc *BitsetCreator) Go() goga.Bitset {
	b := goga.Bitset{}
	b.Create(len(bc.IngredientsInInventory))
	randomSelectedRows := make([]domain.InventoryCell, 0, bc.Capacity)
	for i := 0; i < bc.Capacity; i++ {
		randomIndex := rand.Intn(len(bc.IngredientsInInventory))
		randomSelectedRows = append(randomSelectedRows, bc.IngredientsInInventory[randomIndex])
	}
	for i := 0; i < len(bc.IngredientsInInventory); i++ {
		result := slices.ContainsFunc(randomSelectedRows, func(r domain.InventoryCell) bool {
			return r == bc.IngredientsInInventory[i]
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
