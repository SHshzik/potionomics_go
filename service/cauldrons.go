package service

import (
	"strconv"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/google/uuid"
)

func GetBDCauldrons(cauldronsRecords [][]string) domain.BDCauldrons {
	allCauldrons := make(domain.BDCauldrons, len(cauldronsRecords))

	for i, cauldronsRecord := range cauldronsRecords[1:] {
		id := uuid.New().String()
		capacity, _ := strconv.Atoi(cauldronsRecord[1])
		magmin, _ := strconv.Atoi(cauldronsRecord[2])
		allCauldrons[id] = domain.Cauldron{
			ID:       id,
			Name:     cauldronsRecord[0],
			Capacity: capacity,
			Magmin:   magmin,
			Translit: cauldronsRecord[10],
			Order:    i,
		}
	}

	return allCauldrons
}
