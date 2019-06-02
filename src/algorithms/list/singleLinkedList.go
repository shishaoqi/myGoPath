package list

import (
	"fmt"
	)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

// Create
func CreateNode(data interface{}) *Node {
	return &Node{data: data}
}

// InsertFirst
func (list *LinkedList)AddNodeHead(data interface{}) *Node {
	node := &Node{data: data}
	if list.head != nil {
		node.next = list.head
	}
	list.head = node
	return node
}

// Append
// InsertLast
func (list *LinkedList)AddNodeTail(val interface{}) *Node {
	node := &Node{data: val}

	if list.head == nil {
		list.head = node
	}
	if list.tail == nil {
		list.tail = node
		return node
	}

	list.tail.next = node
	list.tail = node
	return node
}

func (list *LinkedList)InsertNode(oldNode *Node, data interface{}) bool {
	insertN := Node{data: data}

	node := list.head
	for node != nil {
		if node == oldNode {
			nextN := node.next
			if nextN != nil {
				insertN.next = nextN
			} else {
				list.tail = &insertN
			}
			node.next = &insertN
			return true
		}
		node = node.next
	}
	fmt.Printf("not found node of %d, append %v  in the last of list\n", oldNode, data)
	list.AddNodeTail(data)
	return false
}

//
func (list *LinkedList)InsertNodeByValue(oldVal, insertVal interface{}) bool {
	node := list.head
	for node != nil{
		if node.data == oldVal {
			nextN := node.next
			insertN := Node{data: insertVal}
			if nextN != nil {
				insertN.next = nextN
			} else {
				list.tail = &insertN
			}
			node.next = &insertN
			return true
		}
		node = node.next
	}

	fmt.Printf("not found value of %d, append %d  in the last of list\n", oldVal, insertVal)
	list.AddNodeTail(insertVal)
	return false
}

func (list *LinkedList)DelNode(delNode *Node) bool {
	node := list.head
	head := list.head
	preNode := list.head

	for node != nil {
		if node == delNode {
			if node == head {
				list.head = node.next
			}
			preNode.next = node.next
			return true
		}
		preNode = node
		node = node.next
	}
	return false
}

// RemoveByValue
func (list *LinkedList)RemoveByValue(val int) bool {
	current := list.head
	for current != nil {
		if current.data == val {
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
		if current.data == val {
			list.head = current.next
			return true
		}
		current = current.next
	}
	return false
}

// GetFirst
func (list LinkedList) GetFirst() (*Node, bool) {
	current := list.head
	if current == nil {
		return nil, false
	}
	return current, true
}

// GetLast
func (list LinkedList) GetLast() (*Node, bool) {
	tail := list.tail
	if tail == nil {
		return nil, false
	}

	return tail, true
}

// Iterate
func (list *LinkedList)Iterate() {
	tempN := list.head
	i := 0
	for {
		fmt.Printf("list node %d = %d\n", i, tempN.data)
		if tempN.next == nil {
			break
		}
		i++
		tempN = tempN.next
	}
}







