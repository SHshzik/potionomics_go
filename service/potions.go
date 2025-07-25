package service

import (
	"strconv"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/google/uuid"
)

func GetBDPotions(potionsRecords [][]string) domain.BDPotions {
	allPotions := make(domain.BDPotions, len(potionsRecords))

	for i, potionRecord := range potionsRecords[1:] {
		id := uuid.New().String()
		a, _ := strconv.Atoi(potionRecord[3])
		b, _ := strconv.Atoi(potionRecord[4])
		c, _ := strconv.Atoi(potionRecord[5])
		d, _ := strconv.Atoi(potionRecord[6])
		e, _ := strconv.Atoi(potionRecord[7])
		allPotions[id] = domain.Potion{
			ID:          id,
			Name:        potionRecord[0],
			Proportions: []int{a, b, c, d, e},
			Translit:    potionRecord[9],
			Order:       i,
		}
	}

	return allPotions
}
