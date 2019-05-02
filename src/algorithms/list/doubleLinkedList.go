package list

import "algorithms/interface"

type DbNode struct {
	data _interface.StuctType
	pre DbNode
	next DbNode
}

type DblList struct {
	head *DblList
}

// InsertFirst

// InsertLast

// RemoveByValue

// RemoveByIndex

// SearchValue

// GetFirst

// GetLast

// Foreach