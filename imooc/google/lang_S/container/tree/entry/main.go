package main

import (
	"fmt"
	"learngo/imooc/google/lang_S/container/tree"
)

// 1.组合,扩充已有类型的功能
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Val: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	var pRoot *tree.Node
	pRoot.SetValue(2)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()

	fmt.Println("Traverse...")
	pRoot.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxVal := 0
	for node := range c {
		fmt.Printf("%d ", node.Val)
		if node.Val > maxVal {
			maxVal = node.Val
		}
	}
	fmt.Println("\nmax value:", maxVal)
	//fmt.Println()
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()
}
