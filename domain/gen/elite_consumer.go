package gen

import (
	"fmt"
	"strings"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/tomcraven/goga"
)

type EliteConsumer struct {
	currentIter   int
	Ingredients   []domain.Ingredient
	ResultChannel chan string
}

func (ec *EliteConsumer) OnElite(g goga.Genome) {
	bits := g.GetBits().GetAll()
	ec.currentIter++
	rString := strings.Builder{}
	for i, selected := range bits {
		if selected == 1 {
			ingredient := ec.Ingredients[i]
			rString.WriteString(fmt.Sprintf("%s,", ingredient.Name))
		}
	}

	ec.ResultChannel <- rString.String()
}
