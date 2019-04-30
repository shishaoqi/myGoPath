package list

import "fmt"

type Node struct {
	value int
	next *Node
}

type Head struct {
	length int
	start *Node
}

// create
func Create(val int) *Node{
	return &Node{value: val}
}

// append
func (head *Node)Append(val int) *Node {
	tempN := head
	for {
		if tempN.next == nil {
			break
		}
		tempN = tempN.next
	}
	nextN := Node{value: val}

	tempN.next = &nextN
	return head
}

// insert
func (head *Node)Insert(val, insertVal int) *Node {
	isFind := false
	currenN := head
	for {
		if currenN.value == val {
			nextN := currenN.next
			insertN := Node{value: insertVal}
			if nextN != nil {
				insertN.next = nextN
			}
			currenN.next = &insertN
			isFind = true
		}
		if currenN.next != nil {
			currenN = currenN.next
		} else {
			if !isFind {
				fmt.Printf("not found value of %d, append %d  in the last of list\n", val, insertVal)
			}
			break
		}
	}
	return head
}

// del
func (head *Node)Del(val int) *Node {
	tempN := head
	for {
		if tempN.value == val {
			head = tempN.next
			break
		}
		tempN = tempN.next
	}
	return head
}

// iterate
func (head *Node)Iterate() {
	tempN := head
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







