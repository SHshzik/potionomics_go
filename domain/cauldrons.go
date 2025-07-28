package domain

type BDCauldrons map[ID]Cauldron

type Cauldron struct {
	ID       ID     `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Magmin   int    `json:"magmin"`
	Translit string `json:"translit"`
}
