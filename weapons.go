package main

type WeaponPart struct {
	Item
	PartType string
}

type Weapon struct {
	Damage float64
	Speed  float64
}

type Sword struct {
	Weapon
	Blade         WeaponPart
	Guard         WeaponPart
	Handle        WeaponPart
	SheathAugment Item
}
