package tree

import "fmt"

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