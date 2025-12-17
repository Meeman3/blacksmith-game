package main

import (
	"fmt"
)

func main() {
	var option string
	fmt.Printf("Hi, welcome to the game, what would you like to do? Select one (only forge available for now)\n")
	fmt.Printf("Mine\n")
	fmt.Printf("Inventory\n")
	fmt.Printf("Forge\n")
	fmt.Scanln(&option)

	if option == "Forge" {
		forgingChoice()
	} else {
		fmt.Printf("This is either not a function or is not implemented yet. Please choose again\n")
		main()
	}
}

func forgingChoice() {
	var choice string
	fmt.Printf("What would you like to Forge?\n")
	fmt.Printf("Weapon Part\n")
	fmt.Printf("Weapon\n")
	fmt.Scanln(&choice)

	if choice == "Weapon" {
		weaponChoice()
	}
	if choice == "WeaponPart" {
		partCrafting()
	}
}

func weaponChoice() {
	var weapon, Part1, Part2, Part3 string
	fmt.Printf("Please choose your weapon type\n")
	fmt.Printf("Sword (More coming later)\n")
	fmt.Scanln(&weapon)

	if weapon == "Sword" {
		fmt.Printf("What Blade would you like to use\n")
		fmt.Scanln(&Part1)
		fmt.Printf("What Guard would you like to use\n")
		fmt.Scanln(&Part2)
		fmt.Printf("What Handle would you like to use\n")
		fmt.Scanln(&Part3)

	} else {
		fmt.Printf("Invalid choice, please try again\n")
	}
}

func partCrafting() {
	var Material, PartType string
	fmt.Printf("Please choose what weapon type you would like to make:\n")
	fmt.Printf("Blade\n")
	fmt.Printf("Guard\n")
	fmt.Printf("Handle\n")
	fmt.Scanln(&PartType)
	fmt.Scanln("Please choose what material to make it out of\n")
	fmt.Scanln(&Material)
}
