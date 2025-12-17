package main

import "fmt"

type WeaponPart struct {
	BasicItem
	PartType           string
	DamageMultiplier   float64
	HandlingMultiplier float64
	WeightMultiplier   float64
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
var HighestPartID int
var HighestSwordID int

// placeholder damage, speed, part and sword
var defaultSwordDamage int = 1
var defaultSwordSpeed int = 1

var defaultPart WeaponPart = WeaponPart{
	BasicItem: BasicItem{
		ID:       "P000001",
		Name:     "BasicPart",
		MaxStack: 32},
	PartType:           "Basic",
	DamageMultiplier:   1.0,
	HandlingMultiplier: 1.0,
	WeightMultiplier:   1.0,
}

var defaultSword Item = CreateWeapon("Sword", defaultPart, defaultPart, defaultPart).(Sword)

func newPartID() string {
	HighestPartID++
	return fmt.Sprintf("P%.6d", HighestPartID)
}

func CreateWeaponPart(PartType string, Material Material) WeaponPart {
	newPart := defaultPart
	// these are all the same rn but later will have different attributes depending on type
	if PartType == "Handle" {
		newPart := WeaponPart{
			BasicItem: BasicItem{
				ID:       newPartID(),
				Name:     Material.Name + " " + PartType,
				MaxStack: 1,
			},
			PartType:           PartType,
			DamageMultiplier:   Material.DamageMultiplier,
			HandlingMultiplier: Material.HandlingMultiplier,
			WeightMultiplier:   Material.WeightMultiplier,
		}
		return newPart

	}
	if PartType == "Guard" {
		newPart := WeaponPart{
			BasicItem: BasicItem{
				ID:       newPartID(),
				Name:     Material.Name + " " + PartType,
				MaxStack: 1,
			},
			PartType:           PartType,
			DamageMultiplier:   Material.DamageMultiplier,
			HandlingMultiplier: Material.HandlingMultiplier,
			WeightMultiplier:   Material.WeightMultiplier,
		}
		return newPart
	}
	if PartType == "Blade" {
		newPart := WeaponPart{
			BasicItem: BasicItem{
				ID:       newPartID(),
				Name:     Material.Name + " " + PartType,
				MaxStack: 1,
			},
			PartType:           PartType,
			DamageMultiplier:   Material.DamageMultiplier,
			HandlingMultiplier: Material.HandlingMultiplier,
			WeightMultiplier:   Material.WeightMultiplier,
		}
		return newPart
	}
	return newPart
}

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
				Damage: float64(defaultSwordDamage) * Part1.DamageMultiplier * Part2.DamageMultiplier * Part3.DamageMultiplier,
				Speed:  (float64(defaultSwordSpeed) * Part1.HandlingMultiplier * Part2.HandlingMultiplier * Part3.HandlingMultiplier) / (Part1.WeightMultiplier + Part2.WeightMultiplier + Part2.WeightMultiplier),
			},
			Blade:  Part1,
			Guard:  Part2,
			Handle: Part3,
		}
		return newSword
	}

	return nil
}
