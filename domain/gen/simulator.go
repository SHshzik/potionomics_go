package gen

import (
	"context"
	"fmt"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/tomcraven/goga"
)

type BrewSimulator struct {
	IngredientsInInventory []domain.Ingredient
	Capacity               int
	ResultChannel          chan []domain.Ingredient
	MaxA                   int
	MaxB                   int
	MaxC                   int
	MaxD                   int
	MaxE                   int
	MinFitness             int
	Ctx                    context.Context
}

func (bs *BrewSimulator) OnBeginSimulation() {
}

func (bs *BrewSimulator) OnEndSimulation() {
}

func (bs *BrewSimulator) Simulate(g goga.Genome) {
	bits := g.GetBits().GetAll()
	a, b, c, d, e, weight, value := 0, 0, 0, 0, 0, 0, 0
	if countOnes(bits) <= bs.Capacity {
		for i, selected := range bits {
			if selected == 1 {
				item := bs.IngredientsInInventory[i]
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
		mixins := calculateMixins(a, b, c, d, e, bs.MaxA, bs.MaxB, bs.MaxC, bs.MaxD, bs.MaxE)
		if weight > 0 {
			valueF := float64(value) * (1 - (float64(mixins) / float64(weight)))

			if (bs.MaxA > 0 && a > bs.MaxA) ||
				(bs.MaxB > 0 && b > bs.MaxB) ||
				(bs.MaxC > 0 && c > bs.MaxC) ||
				(bs.MaxD > 0 && d > bs.MaxD) ||
				(bs.MaxE > 0 && e > bs.MaxE) {
				valueF = 0
			}
			if valueF > 0 && weight > (bs.MaxA+bs.MaxB+bs.MaxC+bs.MaxD+bs.MaxE) {
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
	select {
	case <-bs.Ctx.Done():
		fmt.Println("Работа отменена:", bs.Ctx.Err())
	default:
		if g.GetFitness() > bs.MinFitness {
			close(bs.ResultChannel)
		}
		return g.GetFitness() > bs.MinFitness
	}

	return false
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
