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

//用闭包函数 改版
func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node){
		n.Print() // 闭包自由发挥处
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)){
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node Node) Print() { // func print(node Node) {
	fmt.Print(node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node)TraverseWithChannel() chan *Node {
	outChan := make(chan *Node)

	go func() {
		node.TraverseFunc(func(n *Node) {
			outChan <- n
		})
		close(outChan)
	}()
	return outChan
}