package gen

import (
	"fmt"
	"strings"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/tomcraven/goga"
)

type BrewSimulator struct {
	Ingredients []domain.Ingredient
	Capacity      int
	ResultChannel chan string
}

func (bs *BrewSimulator) OnBeginSimulation() {
}

func (bs *BrewSimulator) OnEndSimulation() {
}

func (bs *BrewSimulator) Simulate(g goga.Genome) {
	maxA := 80
	maxB := 80
	maxC := 0
	maxD := 0
	maxE := 0

	bits := g.GetBits().GetAll()
	a, b, c, d, e, weight, value := 0, 0, 0, 0, 0, 0, 0
	if countOnes(bits) <= bs.Capacity {
		for i, selected := range bits {
			if selected == 1 {
				item := bs.Ingredients[i]
				a += item.A
				b += item.B
				c += item.C
				d += item.D
				e += item.E
				m := item.A + item.B + item.C + item.D + item.E
				weight += m
				value += m
			}
		}
		mixins := calculateMixins(a, b, c, d, e, maxA, maxB, maxC, maxD, maxE)
		if weight > 0 {
			valueF := float64(value) * (1 - (float64(mixins) / float64(weight)))
			if (maxA > 0 && a > maxA) ||
				(maxB > 0 && b > maxB) ||
				(maxC > 0 && c > maxC) ||
				(maxD > 0 && d > maxD) ||
				(maxE > 0 && e > maxE) {
				valueF = 0
			}
			if valueF > 0 && weight > (maxA+maxB+maxC+maxD+maxE) {
				valueF = 0
			}
			if valueF > 0 && weight > 0 && ((float64(mixins)/float64(weight))*100 > 25) {
				valueF = 0
			}
			g.SetFitness(int(valueF))
		} else {
			g.SetFitness(0)
		}
	} else {
		g.SetFitness(0)
	}
}

func (bs *BrewSimulator) ExitFunc(g goga.Genome) bool {
	if g.GetFitness() > 130 {
		close(bs.ResultChannel)
	}
	return g.GetFitness() > 130
}

// calculateMixins аналогична calculate_mixins из Python
func calculateMixins(a, b, c, d, e, maxA, maxB, maxC, maxD, maxE int) int {
	mixins := 0
	if maxA == 0 {
		mixins += a
	}
	if maxB == 0 {
		mixins += b
	}
	if maxC == 0 {
		mixins += c
	}
	if maxD == 0 {
		mixins += d
	}
	if maxE == 0 {
		mixins += e
	}
	return mixins
}

// Вспомогательная функция для подсчёта единиц
func countOnes(arr []int) int {
	count := 0
	for _, v := range arr {
		if v == 1 {
			count++
		}
	}
	return count
}
