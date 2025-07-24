package service

import "github.com/SHshzik/potionomics_go/domain"

func GetBDPotions(potionsRecords [][]string) domain.BDPotions {
	allPotions := make(domain.BDPotions, len(potionsRecords))

	for _, potionRecord := range potionsRecords[1:] {
		baseName := potionRecord[0]
		name := toLower(baseName)
		allPotions[name] = domain.Potion{
			Name: baseName,
		}
	}

	return allPotions
}
