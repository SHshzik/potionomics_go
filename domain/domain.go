package domain

type ID int

type BDIngredients map[string]Ingredient

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

type GenerateRequest struct {
	Potion   Potion
	Cauldron Cauldron
	WithShop bool
	WithOpt  bool
}
