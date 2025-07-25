package gen

import (
	"context"
	"fmt"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/tomcraven/goga"
)

type BrewSimulator struct {
	IngredientsInInventory []domain.InventoryCell
	Capacity               int
	ResultChannel          chan []domain.Ingredient
	MaxA                   int
	MaxB                   int
	MaxC                   int
	MaxD                   int
	MaxE                   int
	MinFitness             int
	Ctx                    context.Context
	Proportions            []int
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
				item := bs.IngredientsInInventory[i].Ingredient
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
		maxA, maxB, maxC, maxD, maxE := CalculateMaxValues(weight, bs.Proportions)
		if weight > 0 {
			valueF := float64(value) * (1 - (float64(mixins) / float64(weight)))
			fmt.Println(valueF)

			if (bs.MaxA > 0 && a > maxA) ||
				(bs.MaxB > 0 && b > maxB) ||
				(bs.MaxC > 0 && c > maxC) ||
				(bs.MaxD > 0 && d > maxD) ||
				(bs.MaxE > 0 && e > maxE) {
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
		close(bs.ResultChannel)
		return true
	default:
		if g.GetFitness() > bs.MinFitness {
			close(bs.ResultChannel)
		}
		return g.GetFitness() > bs.MinFitness
	}
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

func CalculateMaxValues(maxVolume int, proportions []int) (int, int, int, int, int) {
	sum := 0
	for _, p := range proportions {
		sum += p
	}

	result := make([]int, 5)
	for i, p := range proportions {
		result[i] = maxVolume * p / sum
	}

	return result[0], result[1], result[2], result[3], result[4]
}
