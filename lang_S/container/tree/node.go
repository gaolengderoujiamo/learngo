package tree

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func (node *Node) Print() {
	fmt.Println(node.Val)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Val = value
}

// 工厂函数创建node
func CreateNode(value int) *Node {
	return &Node{Val: value} // 返回局部变量的地址,在go语言中是允许的
}
