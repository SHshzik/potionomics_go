package domain

type Ingredient struct {
	Name string
	A    int
	B    int
	C    int
	D    int
	E    int
}

type Potion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Cauldron struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GenerateRequest struct {
	Potion   Potion
	Cauldron Cauldron
}
