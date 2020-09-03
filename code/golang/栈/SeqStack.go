/**
* @Author: XJX
* @Description: 顺序栈实现
* @File: SeqStack.go
* @Date: 2020/9/3 15:05
 */

package main

import (
	"fmt"
	"log"
)

//定义顺序栈的结构体
type SeqStack struct {
	data []string
	top  int
	n    int
}

//定义顺序栈的接口
type StackFun interface {
	InitStack(int)
	StackEmpty() bool
	StackFull() bool
	StackLength() int
	Push(string) error
	Pop() string
	StackTop() string
	ClearStack()
}

//初始化栈
func (S *SeqStack) InitStack(n int) {
	S.data = make([]string, n)
	fmt.Println(len(S.data))
	S.n = n
	S.top = -1
}

//判断栈是否空
func (S *SeqStack) StackEmpty() bool {
	return S.top == -1
}

//判断栈是否为满
func (S *SeqStack) StackFull() bool {
	fmt.Println("top", S.top, "n", S.n-1)
	return S.top >= (S.n - 1)
}

//栈的长度
func (S *SeqStack) StackLength() int {
	return S.top + 1
}

//入栈 支持动态扩容
func (S *SeqStack) Push(data string) error {
	if S.StackFull() {
		fmt.Println(data, "run here")
		newLen := S.n * 2
		newData := make([]string, newLen)
		for i := 0; i < S.top; i++ {
			newData[i] = S.data[i]
		}
		S1 := SeqStack{
			data: newData,
			n:    newLen,
			top:  S.top,
		}
		*S = S1
	}
	S.top++
	S.data[S.top] = data
	return nil
}

//出栈
func (S *SeqStack) Pop() string {
	tmp := ""
	if S.StackEmpty() {
		log.Fatal("空栈...")
	} else {
		tmp = S.data[S.top]
		S.top--
	}
	return tmp
}

//获取栈顶
func (S *SeqStack) StackTop() string {
	if S.StackEmpty() {
		log.Fatal("空栈...")
	}
	return S.data[S.top]
}

//清空栈
func (S *SeqStack) ClearStack() {
	S.top = -1
}

func main() {
	var s StackFun
	s = &SeqStack{}
	s.InitStack(5)
	s.Push("1")
	s.Push("2")
	s.Push("3")
	s.Push("4")
	s.Push("5")
	s.Push("6")
	s.Push("7")
	s.Push("8")
	d := s.Pop()
	fmt.Println("pop", d)
	fmt.Println("len", s.StackLength())
	fmt.Println("top",s.StackTop())
}
