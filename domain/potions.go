package domain

type BDPotions map[ID]Potion

type Potion struct {
	ID        ID     `json:"id"`
	Name      string `json:"name"`
	Tier      int    `json:"tier"`
	BaseValue int    `json:"base_value"`
	A         int    `json:"a"`
	B         int    `json:"b"`
	C         int    `json:"c"`
	D         int    `json:"d"`
	E         int    `json:"e"`
	Balance   []int  `json:"balance"`
	Translit  string `json:"translit"`
}
