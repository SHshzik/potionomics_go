package gen

import (
	"fmt"
	"strings"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/tomcraven/goga"
)

type EliteConsumer struct {
	currentIter int
	Ingredients []domain.Ingredient
}

func (ec *EliteConsumer) OnElite(g goga.Genome) {
	bitset := g.GetBits()
	bits := bitset.GetAll()
	ec.currentIter++
	rString := strings.Builder{}
	for i, selected := range bits {
		if selected == 1 {
			ingredient := ec.Ingredients[i]
			rString.WriteString(fmt.Sprintf("%s,", ingredient.Name))
		}
	}

	fmt.Println(ec.currentIter, "\t", rString.String(), "\t", g.GetFitness())
}
