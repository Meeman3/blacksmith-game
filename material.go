package main

type Material struct {
	BasicItem
	Rarity int
}

// Material types
var Wood Material = Material{
	BasicItem: BasicItem{ID: "00001",
		Name:     "Wood",
		MaxStack: 64},
	Rarity: 0,
}

var Stone Material = Material{
	BasicItem: BasicItem{ID: "00002",
		Name:     "Stone",
		MaxStack: 64},
	Rarity: 1,
}

var Coal Material = Material{
	BasicItem: BasicItem{ID: "00003",
		Name:     "Coal",
		MaxStack: 64},
	Rarity: 1,
}

var Iron Material = Material{
	BasicItem: BasicItem{ID: "00004",
		Name:     "Iron",
		MaxStack: 64},
	Rarity: 1,
}

var Steel Material = Material{
	BasicItem: BasicItem{ID: "00005",
		Name:     "Steel",
		MaxStack: 64},
	Rarity: 2,
}

var Mithril Material = Material{
	BasicItem: BasicItem{ID: "00006",
		Name:     "Mithril",
		MaxStack: 16},
	Rarity: 4,
}

var DragonBones Material = Material{
	BasicItem: BasicItem{ID: "00007",
		Name:     "Dragon Bones",
		MaxStack: 2},
	Rarity: 5,
}
