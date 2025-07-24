package service

import (
	"strconv"

	"github.com/SHshzik/potionomics_go/domain"
)

func GetBDCauldrons(cauldronsRecords [][]string) domain.BDCauldrons {
	allCauldrons := make(domain.BDCauldrons, len(cauldronsRecords))

	for _, cauldronsRecord := range cauldronsRecords[1:] {
		baseName := cauldronsRecord[0]
		name := toLower(baseName)
		capacity, _ := strconv.Atoi(cauldronsRecord[1])
		magmin, _ := strconv.Atoi(cauldronsRecord[2])
		allCauldrons[name] = domain.Cauldron{
			Name:     baseName,
			Capacity: capacity,
			Magmin:   magmin,
		}
	}

	return allCauldrons
}
