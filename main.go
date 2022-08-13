package main

import "fmt"



//哈希表 数组+链表
//关键：实现hash函数、解决哈希冲突（实现链表）、链表的增删改除、以及获取方式。
//时间复杂度O(1)

func main() {
	PutData("1","a")
	PutData("2","b")
	PutData("3","c")
	PutData("4","d")
	PutData("5","e")
	PutData("6","f")
	PutData("7","g")
	PutData("8","h")
	PutData("9","i")
	PutData("10","j")
	PutData("11","k")
	PutData("12","l")
	PutData("13","m")
	PutData("14","n")
	PutData("15","o")
	PutData("16","p")
	v,_:=GetData("1")
	v1,_:=GetData("16")
	fmt.Println(v)
	fmt.Println(v1)
	PrintNode(hashArr[11])
	fmt.Println("")
	GetAll()
}
