package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

func CreateNode(value int) *Node{
	return &Node{Value: value}
}

//原始
/*func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}*/

func (node Node) Print() { // func print(node Node) {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	node.Value = value
}