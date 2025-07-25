package gen

import (
	"github.com/SHshzik/potionomics_go/domain"
	"github.com/tomcraven/goga"
)

type EliteConsumer struct {
	IngredientsInInventory []domain.InventoryCell
	ResultChannel          chan []domain.Ingredient
}

func (ec *EliteConsumer) OnElite(g goga.Genome) {
	if g.GetFitness() > 0 {
		bits := g.GetBits().GetAll()
		result := make([]domain.Ingredient, 0, len(bits))
		for i, selected := range bits {
			if selected == 1 {
				result = append(result, ec.IngredientsInInventory[i].Ingredient)
			}
		}

		ec.ResultChannel <- result
	}
}
