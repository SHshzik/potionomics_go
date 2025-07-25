package domain

type (
	BDPotions     map[string]Potion
	BDIngredients map[string]Ingredient
	BDCauldrons   map[string]Cauldron
)

type BrewResult struct {
	ID      string `json:"id"`
	Receipt []Ingredient
}

type InventoryCell struct {
	Ingredient Ingredient `json:"ingredient"`
	CellNumber int        `json:"cell_number"`
}

type Ingredient struct {
	Name     string `json:"name"`
	A        int
	B        int
	C        int
	D        int
	E        int
	Translit string `json:"translit"`
}

type Potion struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Proportions []int  `json:"proportions"`
	Translit    string `json:"translit"`
	Order       int    `json:"order"`
}

type Cauldron struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Magmin   int    `json:"magmin"`
	Translit string `json:"translit"`
	Order    int    `json:"order"`
}

type GenerateRequest struct {
	Potion   Potion
	Cauldron Cauldron
}
