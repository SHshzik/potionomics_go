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

type Ingredient struct {
	Name string `json:"name"`
	A    int
	B    int
	C    int
	D    int
	E    int
}

type Potion struct {
	Name string `json:"name"`
}

type Cauldron struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Magmin   int    `json:"magmin"`
}

type GenerateRequest struct {
	Potion   Potion
	Cauldron Cauldron
}
