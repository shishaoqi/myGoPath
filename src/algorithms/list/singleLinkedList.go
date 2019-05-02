package list

import (
	"fmt"
	)

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
}

// create
func Create(val int) *Node{
	return &Node{value: val}
}

// InsertFirst
func (list *LinkedList)InsertFirst(i int) {
	data := &Node{value: i}
	if list.head != nil {
		data.next = list.head
	}
	list.head = data
}

// append
// InsertLast
func (list *LinkedList)Append(val int) {
	newN := &Node{value: val}
	if list.head == nil {
		list.head = newN
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newN
}

// insert
func (list *LinkedList)Insert(val, insertVal int) bool {
	currenN := list.head
	for currenN != nil{
		if currenN.value == val {
			nextN := currenN.next
			insertN := Node{value: insertVal}
			if nextN != nil {
				insertN.next = nextN
			}
			currenN.next = &insertN
			return true
		}
		currenN = currenN.next
	}

	fmt.Printf("not found value of %d, append %d  in the last of list\n", val, insertVal)
	list.Append(insertVal)
	return false
}

// RemoveByValue
func (list *LinkedList)RemoveByValue(val int) bool {
	current := list.head
	for current != nil {
		if current.value == val {
			list.head = current.next
			return true
		}
		current = current.next
	}
	return false
}

// RemoveByIndex
func (list *LinkedList)RemoveByIndex(index int) bool {
	current := list.head
	if current == nil {
		fmt.Printf("false, list is empty!")
		return false
	}
	i := 0
	for current != nil {
		if i == index {
			list.head = current.next
			return true
		}
		current = current.next
		i++
	}

	fmt.Printf("false, not find!")
	return false
}

// SearchValue
func (list *LinkedList)SearchValue(val int) bool {
	current := list.head
	for current != nil {
		if current.value == val {
			list.head = current.next
			return true
		}
		current = current.next
	}
	return false
}

// GetFirst
func (list LinkedList) GetFirst() (int, bool) {
	current := list.head
	if current == nil {
		return 0, false
	}
	return current.value, true
}

// GetLast
func (list LinkedList) GetLast() (int, bool) {
	current := list.head
	if current == nil {
		return 0, false
	}
	last := current
	for current != nil {
		last = current
		current = current.next
	}

	return last.value, true
}

// iterate
func (list *LinkedList)Iterate() {
	tempN := list.head
	i := 0
	for {
		fmt.Printf("list node %d = %d\n", i, tempN.value)
		if tempN.next == nil {
			break
		}
		i++
		tempN = tempN.next
	}
}







