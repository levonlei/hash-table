package main

import "fmt"

//实现哈希表


// 定义一个node的数组
var hashArr [16]*Node

// PutData 添加数据
func PutData(k,v string) {
	var pos = HashCode(k)
	if hashArr[pos] == nil {
		//在这个位置先创建一个头节点
		hashArr[pos] = CreateHeadNode(k,v)
		return
	}
	node :=GetTail(hashArr[pos])
	AddNode(k,v,node)
	return
}

// GetData
func GetData(k string)(string,bool) {
	pos := HashCode(k)
	p := hashArr[pos]
	for {
		if p == nil {
			return "",false
		}
		if p.Data.Key == k {
			return p.Data.Value,true
		}
		p=p.Next
	}
}


// GetAll 获取所有数据
func GetAll() {
	for i:=0;i<16;i++ {
		if hashArr[i] == nil {
			fmt.Printf("%d:\n",i)
			continue
		}
		fmt.Printf("%d:",i)
		PrintNode(hashArr[i])
		fmt.Printf("\n")
	}
}

//典型的hash算法
//HashCode 将任何长度的字符串，通过运算，散列成0-15整数
func HashCode(key string) int {
	var index int = 0
	index = int(key[0])
	for k := 0; k < len(key); k++ {
		//1103515245是个好数字，使通过hashCode散列出的0-15的数字的概率是相等的
		index *= (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16 - 1
	return index
}


