package list

type DbNode struct {
	data interface{}
	pre *DbNode
	next *DbNode
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