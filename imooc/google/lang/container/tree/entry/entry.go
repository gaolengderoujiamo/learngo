package main

import (
	"fmt"
	"learngo/imooc/google/lang/container/tree"
)

// 通过组合的方式, 扩充Node类型, 添加了后续遍历的func
// 1."新结构"包含一个"旧结构"类型的变量
// 2.为"新结构"添加func
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

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	root.Print()
	root.SetValue(100)

	fmt.Println("pRoot...")
	var pRoot *tree.Node
	pRoot.SetValue(300)
	pRoot = &root
	pRoot.SetValue(200)
	pRoot.Print()
	fmt.Println()

	fmt.Println("Tranverse......")
	root.Tranverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Printf("nodeCount = %d\n", nodeCount)
	//nodes := []treeNode {
	//	{value: 3},
	//	{},
	//	{6,nil,&root},
	//}
	//for _,v := range nodes {
	//	fmt.Println(v)
	//}
	fmt.Println("postOrder......")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	fmt.Println("tranverse with channel......")
	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
