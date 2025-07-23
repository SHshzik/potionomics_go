package domain

type Ingredient struct {
	Name string
	A    int
	B    int
	C    int
	D    int
	E    int
}

var mapSaveToDomainIngredientName = map[string]string{}
