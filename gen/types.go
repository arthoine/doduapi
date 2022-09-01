package gen

type SearchIndexedItem struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SuperType   string `json:"super_type"`
	TypeName    string `json:"type_name"`
}

type SearchIndexedMount struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	FamilyName string `json:"family_name"`
}

type SearchIndexedSet struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"highest_equipment_level"`
}

type EffectConditionDbEntry struct {
	Id   int
	Name string
}

type MappedMultiangCondition struct {
	Element   string            `json:"element"`
	ElementId int               `json:"element_id"`
	Operator  string            `json:"operator"`
	Value     int               `json:"value"`
	Templated map[string]string `json:"templated"`
}

type MappedMultilangRecipe struct {
	ResultId int                          `json:"result_id"`
	Entries  []MappedMultilangRecipeEntry `json:"entries"`
}

type MappedMultilangRecipeEntry struct {
	ItemId   int `json:"item_id"`
	Quantity int `json:"quantity"`
	//ItemType map[string]string `json:"item_type"`
}

type MappedMultilangSet struct {
	AnkamaId int                       `json:"ankama_id"`
	Name     map[string]string         `json:"name"`
	ItemIds  []int                     `json:"items"`
	Effects  [][]MappedMultilangEffect `json:"effects"`
	Level    int                       `json:"level"`
}

type MappedMultilangMount struct {
	AnkamaId   int                     `json:"ankama_id"`
	Name       map[string]string       `json:"name"`
	FamilyId   int                     `json:"family_id"`
	FamilyName map[string]string       `json:"family_name"`
	Effects    []MappedMultilangEffect `json:"effects"`
}

type MappedMultilangCharacteristic struct {
	Value map[string]string `json:"value"`
	Name  map[string]string `json:"name"`
}

type MappedMultilangEffect struct {
	Min              int               `json:"min"`
	Max              int               `json:"max"`
	Type             map[string]string `json:"type"`
	MinMaxIrrelevant int               `json:"min_max_irrelevant"`
	Templated        map[string]string `json:"templated"`
	ElementId        int               `json:"element_id"`
}

type MappedMultilangItemType struct {
	Id          int               `json:"id"`
	Name        map[string]string `json:"name"`
	SuperTypeId int               `json:"superTypeId"`
	CategoryId  int               `json:"categoryId"`
}

type MappedMultilangItem struct {
	AnkamaId               int                             `json:"ankama_id"`
	Type                   MappedMultilangItemType         `json:"type"`
	Description            map[string]string               `json:"description"`
	Name                   map[string]string               `json:"name"`
	Image                  string                          `json:"image"`
	Conditions             []MappedMultiangCondition       `json:"conditions"`
	Level                  int                             `json:"level"`
	UsedInRecipes          []int                           `json:"used_in_recipes"`
	Characteristics        []MappedMultilangCharacteristic `json:"characteristics"`
	Effects                []MappedMultilangEffect         `json:"effects"`
	DropMonsterIds         []int                           `json:"dropMonsterIds"`
	CriticalHitBonus       int                             `json:"criticalHitBonus"`
	TwoHanded              bool                            `json:"twoHanded"`
	MaxCastPerTurn         int                             `json:"maxCastPerTurn"`
	ApCost                 int                             `json:"apCost"`
	Range                  int                             `json:"range"`
	CriticalHitProbability int                             `json:"criticalHitProbability"`
	Pods                   int                             `json:"pods"`
	IconId                 int                             `json:"iconId"`
}

type JSONGameSpellType struct {
	Id          int `json:"id"`
	LongNameId  int `json:"longNameId"`
	ShortNameId int `json:"shortNameId"`
}

type JSONGameSpell struct {
	Id            int   `json:"id"`
	NameId        int   `json:"nameId"`
	DescriptionId int   `json:"descriptionId"`
	TypeId        int   `json:"typeId"`
	Order         int   `json:"order"`
	IconId        int   `json:"iconId"`
	SpellLevels   []int `json:"spellLevels"`
}

type JSONLangDict struct {
	Texts    map[string]string `json:"texts"`    // "1": "Account- oder Abohandel",
	IdText   map[string]int    `json:"idText"`   // "790745": 27679,
	NameText map[string]int    `json:"nameText"` // "ui.chat.check0": 65984
}

type JSONGameRecipe struct {
	Id            int   `json:"resultId"`
	NameId        int   `json:"resultNameId"`
	TypeId        int   `json:"resultTypeId"`
	Level         int   `json:"resultLevel"`
	IngredientIds []int `json:"ingredientIds"`
	Quantities    []int `json:"quantities"`
	JobId         int   `json:"jobId"`
	SkillId       int   `json:"skillId"`
}

type LangDict struct {
	Texts    map[int]string
	IdText   map[int]int
	NameText map[string]int
}

type JSONGameBonus struct {
	Amount        int   `json:"amount"`
	Id            int   `json:"id"`
	CriterionsIds []int `json:"criterionsIds"`
	Type          int   `json:"type"`
}

type JSONGameAreaBounds struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type JSONGameArea struct {
	Id              int                `json:"id"`
	NameId          int                `json:"nameId"`
	SuperAreaId     int                `json:"superAreaId"`
	ContainHouses   bool               `json:"containHouses"`
	ContainPaddocks bool               `json:"containPaddocks"`
	Bounds          JSONGameAreaBounds `json:"bounds"`
	WorldmapId      int                `json:"worldmapId"`
	HasWorldMap     bool               `json:"hasWorldMap"`
}

type JSONGameItemPossibleEffect struct {
	EffectId     int `json:"effectId"`
	MinimumValue int `json:"diceNum"`
	MaximumValue int `json:"diceSide"`
	Value        int `json:"value"`

	BaseEffectId  int `json:"baseEffectId"`
	EffectElement int `json:"effectElement"`
	Dispellable   int `json:"dispellable"`
	SpellId       int `json:"spellId"`
	Duration      int `json:"duration"`
}

type JSONGameSet struct {
	Id      int                            `json:"id"`
	ItemIds []int                          `json:"items"`
	NameId  int                            `json:"nameId"`
	Effects [][]JSONGameItemPossibleEffect `json:"effects"`
}

type JSONGameItemType struct {
	Id          int `json:"id"`
	NameId      int `json:"nameId"`
	SuperTypeId int `json:"superTypeId"`
	CategoryId  int `json:"categoryId"`
}

type JSONGameEffect struct {
	Id                       int  `json:"id"`
	DescriptionId            int  `json:"descriptionId"`
	IconId                   int  `json:"iconId"`
	Characteristic           int  `json:"characteristic"`
	Category                 int  `json:"category"`
	UseDice                  bool `json:"useDice"`
	Active                   bool `json:"active"`
	TheoreticalDescriptionId int  `json:"theoreticalDescriptionId"`
	BonusType                int  `json:"bonusType"` // -1,0,+1
	ElementId                int  `json:"elementId"`
}

type JSONGameItem struct {
	Id            int `json:"id"`
	TypeId        int `json:"typeId"`
	DescriptionId int `json:"descriptionId"`
	IconId        int `json:"iconId"`
	NameId        int `json:"nameId"`
	Level         int `json:"level"`

	PossibleEffects        []JSONGameItemPossibleEffect `json:"possibleEffects"`
	RecipeIds              []int                        `json:"recipeIds"`
	Pods                   int                          `json:"realWeight"`
	ParseEffects           bool                         `json:"useDice"`
	EvolutiveEffectIds     []int                        `json:"evolutiveEffectIds"`
	DropMonsterIds         []int                        `json:"dropMonsterIds"`
	ItemSetId              int                          `json:"itemSetId"`
	Criteria               string                       `json:"criteria"`
	CriticalHitBonus       int                          `json:"criticalHitBonus"`
	TwoHanded              bool                         `json:"twoHanded"`
	MaxCastPerTurn         int                          `json:"maxCastPerTurn"`
	ApCost                 int                          `json:"apCost"`
	Range                  int                          `json:"range"`
	CriticalHitProbability int                          `json:"criticalHitProbability"`
}

type JSONGameBreed struct {
	Id            int `json:"id"`
	ShortNameId   int `json:"shortNameId"`
	LongNameId    int `json:"longNameId"`
	DescriptionId int `json:"descriptionId"`
}

type JSONGameMount struct {
	Id       int                          `json:"id"`
	FamilyId int                          `json:"familyId"`
	NameId   int                          `json:"nameId"`
	Effects  []JSONGameItemPossibleEffect `json:"effects"`
}

type JSONGameMountFamily struct {
	Id      int    `json:"id"`
	NameId  int    `json:"nameId"`
	HeadUri string `json:"headUri"`
}

type JSONGameNPC struct {
	Id             int     `json:"id"`
	NameId         int     `json:"nameId"`
	DialogMessages [][]int `json:"dialogMessages"`
	DialogReplies  [][]int `json:"dialogReplies"`
	Actions        []int   `json:"actions"`
}

type JSONGameData struct {
	Items         map[int]JSONGameItem
	Sets          map[int]JSONGameSet
	ItemTypes     map[int]JSONGameItemType
	effects       map[int]JSONGameEffect
	bonuses       map[int]JSONGameBonus
	Recipes       map[int]JSONGameRecipe
	spells        map[int]JSONGameSpell
	spellTypes    map[int]JSONGameSpellType
	areas         map[int]JSONGameArea
	Mounts        map[int]JSONGameMount
	classes       map[int]JSONGameBreed
	Mount_familys map[int]JSONGameMountFamily
	npcs          map[int]JSONGameNPC
}
