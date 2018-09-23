package tree

import (
	"fmt"
)

// Node 定义结构体
type Node struct {
	Value       int
	Left, Right *Node
}

// Print 输出
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// SetValue 赋值
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Value = value
}

// CreateNode Create a new node
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
