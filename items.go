package main

type Item struct {
	ID       string
	Name     string
	MaxStack int
	Rarity   int
	ItemType string
}

var Wood Item = Item{
	ID:       "00001",
	Name:     "Wood",
	MaxStack: 64,
	Rarity:   0,
	ItemType: "material",
}

var Stone Item = Item{
	ID:       "00002",
	Name:     "Stone",
	MaxStack: 64,
	Rarity:   1,
	ItemType: "material",
}

var Coal Item = Item{
	ID:       "00003",
	Name:     "Coal",
	MaxStack: 64,
	Rarity:   1,
	ItemType: "material",
}

var Iron Item = Item{
	ID:       "00004",
	Name:     "Iron",
	MaxStack: 64,
	Rarity:   1,
	ItemType: "material",
}

var Steel Item = Item{
	ID:       "00005",
	Name:     "Steel",
	MaxStack: 64,
	Rarity:   2,
	ItemType: "material",
}

var Mithril Item = Item{
	ID:       "00006",
	Name:     "Mithril",
	MaxStack: 16,
	Rarity:   4,
	ItemType: "material",
}

var DragonBones Item = Item{
	ID:       "00007",
	Name:     "Dragon Bones",
	MaxStack: 2,
	Rarity:   5,
	ItemType: "material",
}
