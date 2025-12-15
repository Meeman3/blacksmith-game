package main

type Material struct {
	BasicItem
	Rarity             int
	DamageMultiplier   float64
	WeightMultiplier   float64
	HandlingMultiplier float64
	HardnessMultiplier float64
}

// Material types
var Wood Material = Material{
	BasicItem: BasicItem{ID: "00001",
		Name:     "Wood",
		MaxStack: 64},
	Rarity:             0,
	DamageMultiplier:   1,
	WeightMultiplier:   1,
	HandlingMultiplier: 1,
	HardnessMultiplier: 1,
}

var Stone Material = Material{
	BasicItem: BasicItem{ID: "00002",
		Name:     "Stone",
		MaxStack: 64},
	Rarity:             1,
	DamageMultiplier:   2,
	WeightMultiplier:   2,
	HandlingMultiplier: 2,
	HardnessMultiplier: 2,
}

var Coal Material = Material{
	BasicItem: BasicItem{ID: "00003",
		Name:     "Coal",
		MaxStack: 64},
	Rarity:             1,
	DamageMultiplier:   0,
	WeightMultiplier:   0,
	HandlingMultiplier: 0,
	HardnessMultiplier: 0,
}

var Iron Material = Material{
	BasicItem: BasicItem{ID: "00004",
		Name:     "Iron",
		MaxStack: 64},
	Rarity:             1,
	DamageMultiplier:   3,
	WeightMultiplier:   3,
	HandlingMultiplier: 3,
	HardnessMultiplier: 3,
}

var Steel Material = Material{
	BasicItem: BasicItem{ID: "00005",
		Name:     "Steel",
		MaxStack: 64},
	Rarity:             2,
	DamageMultiplier:   4,
	WeightMultiplier:   4,
	HandlingMultiplier: 4,
	HardnessMultiplier: 4,
}

var Mithril Material = Material{
	BasicItem: BasicItem{ID: "00006",
		Name:     "Mithril",
		MaxStack: 16},
	Rarity:             4,
	DamageMultiplier:   5,
	WeightMultiplier:   5,
	HandlingMultiplier: 5,
	HardnessMultiplier: 5,
}

var DragonBones Material = Material{
	BasicItem: BasicItem{ID: "00007",
		Name:     "Dragon Bones",
		MaxStack: 2},
	Rarity:             5,
	DamageMultiplier:   6,
	WeightMultiplier:   6,
	HandlingMultiplier: 6,
	HardnessMultiplier: 6,
}
