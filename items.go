package main

type Item struct {
	ID       string
	Name     string
	MaxStack int
	Rarity   int
}

var Wood Item = Item{
	ID:       "00001",
	Name:     "Wood",
	MaxStack: 64,
	Rarity:   0,
}

var Stone Item = Item{
	ID:       "00002",
	Name:     "Stone",
	MaxStack: 64,
	Rarity:   1,
}

var Coal Item = Item{
	ID:       "00003",
	Name:     "Coal",
	MaxStack: 64,
	Rarity:   1,
}

var Iron Item = Item{
	ID:       "00004",
	Name:     "Iron",
	MaxStack: 64,
	Rarity:   1,
}

var Steel Item = Item{
	ID:       "00005",
	Name:     "Steel",
	MaxStack: 64,
	Rarity:   2,
}

var Mithril Item = Item{
	ID:       "00006",
	Name:     "Mithril",
	MaxStack: 16,
	Rarity:   1,
}

var DragonBones Item = Item{
	ID:       "00007",
	Name:     "Dragon Bones",
	MaxStack: 2,
	Rarity:   5,
}
