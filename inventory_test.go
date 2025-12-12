package main

import (
	"testing"
)

var Item1 = Item{
	ID:       "1",
	Name:     "item1",
	MaxStack: 64,
}

var Item2 = Item{
	ID:       "2",
	Name:     "item2",
	MaxStack: 16,
}

var Item3 = Item{
	ID:       "3",
	Name:     "item3",
	MaxStack: 2,
}

var Item4 = Item{
	ID:       "4",
	Name:     "item4",
	MaxStack: 1,
}

func makeTestInventory() Inventory {
	defaultSpace := Space{
		Item: Item{
			ID:       "",
			Name:     "",
			MaxStack: 0,
		},
		StackCount: 0,
		Locked:     false,
	}
	testInv := Inventory{
		Spaces: make([]Space, 6),
	}

	for j := 0; j < len(testInv.Spaces); j++ {
		testInv.Spaces[j] = defaultSpace
	}
	return testInv
}

func TestAddItem(t *testing.T) {
	inv := makeTestInventory()
	inv.Spaces[5].Locked = true

	inv.AddItem(Item1, 129)
	inv.AddItem(Item3, 1)
	inv.AddItem(Item4, 2)
	inv.AddItem(Item1, 5)

	invShouldBe := makeTestInventory()
	invShouldBe.Spaces[0] = Space{
		Item:       Item1,
		StackCount: 64,
		Locked:     false,
	}
	invShouldBe.Spaces[1] = Space{
		Item:       Item1,
		StackCount: 64,
		Locked:     false,
	}
	invShouldBe.Spaces[2] = Space{
		Item:       Item1,
		StackCount: 6,
		Locked:     false,
	}
	invShouldBe.Spaces[3] = Space{
		Item:       Item3,
		StackCount: 1,
		Locked:     false,
	}
	invShouldBe.Spaces[4] = Space{
		Item:       Item4,
		StackCount: 1,
		Locked:     false,
	}
	invShouldBe.Spaces[5] = Space{
		Item: Item{
			ID:       "",
			Name:     "",
			MaxStack: 0,
		},
		StackCount: 0,
		Locked:     true,
	}

	for j := 0; j < len(invShouldBe.Spaces); j++ {
		if inv.Spaces[j] != invShouldBe.Spaces[j] {
			t.Errorf("space %.5d = %v, should be %v", j, inv.Spaces[j], invShouldBe.Spaces[j])
		}
	}

}

func TestRemoveItem(t *testing.T) {
	inv := makeTestInventory()
	inv.Spaces[5].Locked = true

	inv.AddItem(Item1, 129)
	inv.AddItem(Item3, 1)
	inv.AddItem(Item4, 2)
	inv.AddItem(Item1, 5)

	inv.RemoveItem(Item3, 4)
	inv.RemoveItem(Item4, 1)
	inv.RemoveItem(Item1, 1)
	inv.RemoveItem(Item1, 67)

	invShouldBe := makeTestInventory()

	invShouldBe.Spaces[0] = Space{
		Item:       Item1,
		StackCount: 64,
		Locked:     false,
	}

	invShouldBe.Spaces[1] = Space{
		Item:       Item1,
		StackCount: 2,
		Locked:     false,
	}

	invShouldBe.Spaces[2] = EmptySpace

	invShouldBe.Spaces[3] = Space{
		Item:       Item3,
		StackCount: 1,
		Locked:     false,
	}

	invShouldBe.Spaces[4] = Space{
		Item:       EmptyItem,
		StackCount: 0,
		Locked:     false,
	}

	invShouldBe.Spaces[5] = Space{
		Item: Item{
			ID:       "",
			Name:     "",
			MaxStack: 0,
		},
		StackCount: 0,
		Locked:     true,
	}

	for j := 0; j < len(invShouldBe.Spaces); j++ {
		if inv.Spaces[j] != invShouldBe.Spaces[j] {
			t.Errorf("space %.5d = %v, should be %v", j, inv.Spaces[j], invShouldBe.Spaces[j])
		}
	}
}
