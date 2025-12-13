package main

import "fmt"

type WeaponPart struct {
	BasicItem
	PartType string
}

type Weapon struct {
	BasicItem
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

// Weapons can't stack, so max stack is 1
func (w Weapon) GetMaxStack() int {
	return 1
}

// variable to ensure unique ID (only in 1 session)
var HighestSwordID int

// placeholder damage, speed, part and sword
var defaultSwordDamage int = 1
var defaultSwordSpeed int = 1

var defaultPart WeaponPart = WeaponPart{
	BasicItem: BasicItem{
		ID:       "0000010",
		Name:     "BasicPart",
		MaxStack: 32},
	PartType: "Basic",
}
var defaultSword Item = CreateWeapon("Sword", defaultPart, defaultPart, defaultPart)

func newSwordID() string {
	HighestSwordID++
	return fmt.Sprintf("S%.6d", HighestSwordID)
}

// basic weapon creation system, just mushes parts together
func CreateWeapon(WeaponType string, Part1, Part2, Part3 WeaponPart) Item {
	if WeaponType == "Sword" {
		newSword := Sword{
			Weapon: Weapon{
				BasicItem: BasicItem{ID: newSwordID(),
					Name:     fmt.Sprintf("%v + %v + %v Sword", Part1.Name, Part2.Name, Part3.Name),
					MaxStack: 1},
				Damage: float64(defaultSwordDamage),
				Speed:  float64(defaultSwordSpeed),
			},
			Blade:  Part1,
			Guard:  Part2,
			Handle: Part3,
		}
		return newSword
	}

	return nil
}
