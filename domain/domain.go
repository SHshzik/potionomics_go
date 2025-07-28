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
	Name                        string `json:"name"`
	A                           int
	B                           int
	C                           int
	D                           int
	E                           int
	PerfectForHealtPotion       bool
	PerfectForManaPotion        bool
	PerfectForStaminaPotion     bool
	PerfectForSpeedPotion       bool
	PerfectForTolerancePotion   bool
	PerfectForFireTonic         bool
	PerfectForIceTonic          bool
	PerfectForThunderTonic      bool
	PerfectForShadowTonic       bool
	PerfectForRadiationTonic    bool
	PerfectForSightEnhancer     bool
	PerfectForAlertnessEnhancer bool
	PerfectForInsightEnhancer   bool
	PerfectForDowsingEnhancer   bool
	PerfectForSeekingEnhancer   bool
	PerfectForPoisonCure        bool
	PerfectForDrowsinessCure    bool
	PerfectForPetrificationCure bool
	PerfectForSilenceCure       bool
	PerfectForCurseCure         bool
	Translit                    string `json:"translit"`
}

func (i *Ingredient) PerfectForPotion(p Potion) bool {
	if p.Name == "Health Potion" {
		return i.PerfectForHealtPotion
	}
	if p.Name == "Mana Potion" {
		return i.PerfectForManaPotion
	}
	if p.Name == "Stamina Potion" {
		return i.PerfectForStaminaPotion
	}
	if p.Name == "Speed Potion" {
		return i.PerfectForSpeedPotion
	}
	if p.Name == "Tolerance Potion" {
		return i.PerfectForTolerancePotion
	}
	if p.Name == "Fire Tonic" {
		return i.PerfectForFireTonic
	}
	if p.Name == "Ice Tonic" {
		return i.PerfectForIceTonic
	}
	if p.Name == "Thunder Tonic" {
		return i.PerfectForThunderTonic
	}
	if p.Name == "Shadow Tonic" {
		return i.PerfectForShadowTonic
	}
	if p.Name == "Radiation Tonic" {
		return i.PerfectForRadiationTonic
	}
	if p.Name == "Sight Enhancer" {
		return i.PerfectForSightEnhancer
	}
	if p.Name == "Alertness Enhancer" {
		return i.PerfectForAlertnessEnhancer
	}
	if p.Name == "Insight Enhancer" {
		return i.PerfectForInsightEnhancer
	}
	if p.Name == "Dowsing Enhancer" {
		return i.PerfectForDowsingEnhancer
	}
	if p.Name == "Seeking Enhancer" {
		return i.PerfectForSeekingEnhancer
	}
	if p.Name == "Poison Cure" {
		return i.PerfectForPoisonCure
	}
	if p.Name == "Drowsiness Cure" {
		return i.PerfectForDrowsinessCure
	}
	if p.Name == "Petrification Cure" {
		return i.PerfectForPetrificationCure
	}
	if p.Name == "Silence Cure" {
		return i.PerfectForSilenceCure
	}
	if p.Name == "Curse Cure" {
		return i.PerfectForCurseCure
	}

	return false
}

type GenerateRequest struct {
	Potion   Potion
	Cauldron Cauldron
	WithShop bool
	IsStrict bool
}
