/**
* @Author: XJX
* @Description: 链式队列实现
* @File: LinkQueue.go
* @Date: 2020/9/4 16:25
 */

package main

import (
	"fmt"
	"log"
)

//定义链表队列节点的结构体
type LinkQueueNode struct {
	Data interface{}
	Next *LinkQueueNode
}

//定义链表队列的结构体
type LinkQueueList struct {
	Head *LinkQueueNode
	Rear *LinkQueueNode
	Size int
}

//定义链表队列接口
type LinkQueueFun interface {
	InitQueue()
	QueueEmpty() bool
	QueueLength() int
	enQueue(string) error
	deQueue() interface{}
	QueueTop() interface{}
	ClearQueue()
	ShowQueue() []interface{}
}

//初始化队列
func (LQ *LinkQueueList) InitQueue() {
	LQ = &LinkQueueList{
		Size: 0,
		Head: nil,
		Rear: nil,
	}
}

//判断队列是否为空
func (LQ *LinkQueueList) QueueEmpty() bool {
	return LQ.Size == 0
}

//返回队列长度
func (LQ *LinkQueueList) QueueLength() int {
	return LQ.Size
}

//入队列
func (LQ *LinkQueueList) enQueue(data string) error {
	node := &LinkQueueNode{
		Data: data,
		Next: nil,
	}
	if LQ.Size == 0 {
		LQ.Head = node
		LQ.Rear = node
	} else {
		LQ.Rear.Next = node
		LQ.Rear = LQ.Rear.Next
	}
	LQ.Size++
	return nil
}

//出队列
func (LQ *LinkQueueList) deQueue() interface{} {
	if LQ.QueueEmpty() {
		log.Fatal("Queue empty")
	}
	item := LQ.Head.Data
	LQ.Head = LQ.Head.Next
	if LQ.Size == 1 {
		LQ.Rear = nil
	}
	LQ.Size--
	return item
}

//获取队列最前数据
func (LQ *LinkQueueList) QueueTop() interface{} {
	if LQ.QueueEmpty() {
		log.Fatal("Queue empty")
	}
	item := LQ.Head.Data
	return item
}

//打印队列
func (LQ *LinkQueueList) ShowQueue() (resp []interface{}) {
	buf := LQ.Head
	for i := 0; i < LQ.Size; i++ {
		resp = append(resp, buf.Data, "<--")
		buf = buf.Next
	}
	return
}

//清空队列
func (LQ *LinkQueueList) ClearQueue() {
	LQ.Size = 0
	LQ.Rear = nil
	LQ.Head = nil
}

func main() {
	var s LinkQueueFun
	s = &LinkQueueList{}
	s.InitQueue()
	fmt.Println(s.QueueLength())
	s.enQueue("1")
	s.enQueue("2")
	s.enQueue("3")
	s.enQueue("4")
	fmt.Println("len:", s.QueueLength())
	fmt.Println(s.ShowQueue())
	fmt.Println("deQueue:", s.deQueue())
	fmt.Println("deQueue:", s.deQueue())
	fmt.Println(s.ShowQueue())
	fmt.Println("top：", s.QueueTop())
}
