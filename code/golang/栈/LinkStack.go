/**
* @Author: XJX
* @Description: 链栈的实现
* @File: LinkStack.go
* @Date: 2020/9/3 15:03
 */

package main

import (
	"fmt"
	"log"
)

//定义链节点的结构体
type LinkStackNode struct {
	Data interface{}
	Next *LinkStackNode
}

//定义链栈的结构体
type LinkStackList struct {
	Head *LinkStackNode
	Size int
}

//初始化链栈
func (LS *LinkStackList) InitStack() {
	LS.Head = new(LinkStackNode)
	LS.Size = 0
}

func (LS *LinkStackList) Print() {
	if LS.Size == 0 {
		fmt.Println("单链表为空")
	}
	format := ""
	p := LS.Head.Next
	for p != nil {
		format += fmt.Sprintf("%+v", p.Data)
		if p.Next != nil {
			format += "->"
		}
		p = p.Next
	}
	fmt.Println(format)
}

//检查是否空栈
func (LS *LinkStackList) StackEmpty() bool {
	return LS.Size == 0
}

//返回链栈长度
func (LS *LinkStackList) StackLength() int {
	return LS.Size
}

//入栈
func (LS *LinkStackList) Push(data string) error {
	newNode := &LinkStackNode{
		Data: data,
	}
	if LS.Size == 0 {
		newNode.Next = nil
	} else {
		newNode.Next = LS.Head.Next
	}
	LS.Head.Next = newNode
	LS.Size++
	return nil
}

//出栈
func (LS *LinkStackList) Pop() interface{} {
	if LS.StackEmpty() {
		log.Fatal("空栈")
	}
	firstNode := LS.Head.Next
	if LS.Size == 1 {
		LS.Head.Next = nil
	} else {
		LS.Head.Next = firstNode.Next
	}
	data := firstNode.Data
	LS.Size--
	return data
}

//获取顶栈
func (LS *LinkStackList) StackTop() interface{} {
	if LS.StackEmpty() {
		log.Fatal("空栈")
	}
	data := LS.Head.Next.Data
	return data
}

//清空栈
func (LS *LinkStackList) ClearStack() {
	LS.Head.Next = nil
	LS.Size = 0
}

//定义链栈的接口
type LinkStackFun interface {
	InitStack()
	StackEmpty() bool
	StackLength() int
	Push(string) error
	Pop() interface{}
	StackTop() interface{}
	ClearStack()
	Print()
}

func main() {
	var s LinkStackFun
	s = &LinkStackList{}
	s.InitStack()
	s.Push("one")
	s.Push("two")
	s.Print()
	fmt.Println("len", s.StackLength())
	fmt.Println("pop", s.Pop())
	s.Print()
}
