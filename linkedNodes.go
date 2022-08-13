package main

import "fmt"


//Map 定义map的数据结构
type  Map struct {
	Key 	string
	Value 	string
}

// Node 定义Node节点
type Node struct{
	Data Map
	Next *Node
}

// CreateHeadNode 创建头节点
func CreateHeadNode(k,v string) *Node  {
	node:= new(Node)
	node.Data.Key = k
	node.Data.Value = v
	node.Next = nil
	return node
}

// AddNode 新增Node
func AddNode(k,v string,curNode *Node) {
	node := &Node{}
	node.Data.Key = k
	node.Data.Value = v
	node.Next = nil
	curNode.Next = node
}

// PrintNode 打印链表
func PrintNode(head *Node) {
	p:=head
	for {
		if p.Next == nil {
			fmt.Printf("%s:%s",p.Data.Key,p.Data.Value)
			return
		}
		fmt.Printf("%s:%s->",p.Data.Key,p.Data.Value)
		p =p.Next
	}
}


func CountNodeNums(head *Node) int {
	p:=head
	nums:=0
	for {
		if p.Next == nil {
			nums ++
			return nums
		}
		nums ++
		p =p.Next
	}
}

// GetTail 获取链表尾节点
func GetTail(head *Node) *Node  {
	p :=head
	for {
		if p.Next == nil {
			return p
		}
		p=p.Next
	}
}






