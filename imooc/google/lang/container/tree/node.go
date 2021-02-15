package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 工厂函数,取代了构造函数
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// 为结构定义方法
// (node Node) 接受者retriever
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 使用指针作为方法的接受者
// 只有使用指针才能够改变结构内容
// nil指针也可以调用方法!

/*
			值接受者 vs 指针接受者

	要改变内容必须使用指针接受者
	结构过大也考虑使用指针接受者
	一致性: 如有存在指针接受者, 最好都是指针接受者
*/
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node." +
			"Ignored...")
		return
	}
	node.Value = value
}
