package main

// item interface + basic item struct (used in all item structs)
type Item interface {
	GetID() string
	GetName() string
	GetMaxStack() int
}

type BasicItem struct {
	ID       string
	Name     string
	MaxStack int
}

func (b BasicItem) GetID() string {
	return b.ID
}

func (b BasicItem) GetName() string {
	return b.Name
}

func (b BasicItem) GetMaxStack() int {
	return b.MaxStack
}
