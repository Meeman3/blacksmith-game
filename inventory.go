package main

import "fmt"

// types for Items, InventorySpaces and Inventories.
type Item struct {
	ID       string
	Name     string
	MaxStack int
}

type Space struct {
	Item       Item
	StackCount int
	Locked     bool
}

type Inventory struct {
	Spaces []Space
}

var EmptyItem = Item{
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
		if i.Spaces[j].Item.ID == ID {
			num += i.Spaces[j].StackCount
		}
	}

	return num
}

func (i *Inventory) AddItem(item Item, amount int) {
	// check or if it stacks before putting in new slot
	stacked := false

	// iterable for checking if open stacks of item exist
	for j := 0; j < len(i.Spaces); j++ {
		if i.Spaces[j].Item == item && !stacked {
			fmt.Print("adding to stack\n")
			if i.Spaces[j].StackCount < i.Spaces[j].Item.MaxStack {
				countDifference := i.Spaces[j].Item.MaxStack - i.Spaces[j].StackCount
				// if the amount to add is more than available space, add to max then continue searching
				if amount > countDifference {
					amount -= countDifference
					i.Spaces[j].StackCount += countDifference
				} else {
					i.Spaces[j].StackCount += amount
					// if all wanting to add is stacked, doesn't continue
					stacked = true
					break
				}
			}
		}
	}

	// if can't stack then adds to new  space
	if !stacked {
		for j := 0; j < len(i.Spaces); j++ {
			if i.Spaces[j].Item.ID == "" && !i.Spaces[j].Locked {
				fmt.Print("adding to new space\n")
				i.Spaces[j].Item = item
				// if amount to add is greater than max stack, makes a full stack and carries the rest on
				if item.MaxStack < amount {
					i.Spaces[j].StackCount = item.MaxStack
					amount -= item.MaxStack
				} else {
					i.Spaces[j].StackCount = amount
					stacked = true
					break
				}
			}
		}
	}

	if !stacked {
		fmt.Print("No space for items \n")
	}
}

func (i *Inventory) RemoveItem(item Item, amount int) {
	// if there isn't enough in in, cancel
	if i.TotalCount(item.ID) < amount {
		fmt.Printf("Not enough items\n")
		return
	}

	// iterable ovar all spaces
	for j := 0; j < len(i.Spaces); j++ {
		// used to go from the end of spaces instead of start (so we won't have multiple non-full stacks)
		k := len(i.Spaces) - j - 1
		if i.Spaces[k].Item == item {
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
