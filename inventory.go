package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// types for InventorySpaces and Inventories. Item defined in materials.go

type Space struct {
	Item       Item
	StackCount int
	Locked     bool
}

type Inventory struct {
	Spaces []Space
}

// "Empty" items and spaces, defaults
var EmptyItem = BasicItem{
	ID:       "",
	Name:     "",
	MaxStack: 0,
}

var EmptySpace = Space{
	Item:       EmptyItem,
	StackCount: 0,
	Locked:     false,
}

// counts amount of specific item in inventory
func (i *Inventory) TotalCount(ID string) int {
	num := 0

	for j := 0; j < len(i.Spaces); j++ {
		if i.Spaces[j].Item.GetID() == ID {
			num += i.Spaces[j].StackCount
		}
	}

	return num
}

func (i *Inventory) TotalCountName(Name string) int {
	num := 0

	for j := 0; j < len(i.Spaces); j++ {
		if i.Spaces[j].Item.GetName() == Name {
			num += i.Spaces[j].StackCount
		}
	}

	return num
}

func (i *Inventory) AddItem(item Item, amount int) {
	// check if stackable using amount
	for j := 0; j < len(i.Spaces) && amount > 0; j++ {
		if i.Spaces[j].Item.GetID() == item.GetID() {
			fmt.Print("adding to stack\n")
			if i.Spaces[j].StackCount < i.Spaces[j].Item.GetMaxStack() {
				countDifference := i.Spaces[j].Item.GetMaxStack() - i.Spaces[j].StackCount
				// if the amount to add is more than available space, add to max then continue searching
				if amount > countDifference {
					amount -= countDifference
					i.Spaces[j].StackCount += countDifference
				} else {
					i.Spaces[j].StackCount += amount
					amount = 0
					break
				}
			}
		}
	}

	// if can't stack then adds to new  space
	for j := 0; j < len(i.Spaces) && amount > 0; j++ {
		if i.Spaces[j].Item.GetID() == "" && !i.Spaces[j].Locked {
			fmt.Print("adding to new space\n")
			i.Spaces[j].Item = item
			// if amount to add is greater than max stack, makes a full stack and carries the rest on
			if item.GetMaxStack() < amount {
				i.Spaces[j].StackCount = item.GetMaxStack()
				amount -= item.GetMaxStack()
			} else {
				i.Spaces[j].StackCount = amount
				amount = 0
				break
			}
		}
	}

	if amount > 0 {
		fmt.Print("No space for items \n")
	}
}

func (i *Inventory) RemoveItem(item Item, amount int) {
	// if there isn't enough in, cancel
	if i.TotalCount(item.GetID()) < amount {
		fmt.Printf("Not enough items\n")
		return
	}

	// use amount to keep track of when to end
	for j := 0; j < len(i.Spaces) && amount > 0; j++ {
		// used to go from the end of spaces instead of start (so we won't have multiple non-full stacks)
		k := len(i.Spaces) - j - 1
		if i.Spaces[k].Item.GetID() == item.GetID() {
			fmt.Printf("Removing item(s)\n")
			// compares amount to remove to amount in space and remoess accordingly
			if amount > i.Spaces[k].StackCount {
				amount -= i.Spaces[k].StackCount
				i.Spaces[k] = EmptySpace
			}
			if amount < i.Spaces[k].StackCount {
				i.Spaces[k].StackCount -= amount
				amount = 0
				return
			}
			if amount == i.Spaces[k].StackCount {
				i.Spaces[k] = EmptySpace
				amount = 0
				return
			}
		}
	}

}

// functions for saving and loading the inventory, using encoding/json
func (i *Inventory) SaveInventory() error {
	save, _ := json.Marshal(i)

	return os.WriteFile("inventory.json", save, 0644)
}

func (i *Inventory) LoadInventory() error {
	loadedInv, _ := os.ReadFile("inventory.json")
	err := json.Unmarshal(loadedInv, i)

	return err
}
