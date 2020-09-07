/**
* @Author: XJX
* @Description: 顺序队列实现
* @File: SeqQueue.go
* @Date: 2020/9/4 16:24
 */

package main

import (
	"fmt"
	"log"
)

//定义顺序队列结构
type SeqQueue struct {
	data []interface{} //存储数据数组
	top  int           //头队下标
	rear int           //尾队下标
	size int           //内容高度
}

//定义顺序栈的接口
type ArrayQueue interface {
	InitQueue(int)
	QueueEmpty() bool
	QueueFull() bool
	QueueLength() int
	enQueue(string) error
	deQueue() interface{}
	QueueTop() interface{}
	ClearQueue()
	ShowQueue()
}

//初始化队列
func (Q *SeqQueue) InitQueue(n int) {
	data := make([]interface{}, n)
	Q.data = data
	Q.top = -1
	Q.rear = -1
	Q.size = n
}

//检测队列是否为空
func (Q *SeqQueue) QueueEmpty() bool {
	return Q.QueueLength() == 0
}

//检测队列是否为满
func (Q *SeqQueue) QueueFull() bool {
	return Q.rear == (Q.size - 1)
}

//获取队列长度
func (Q *SeqQueue) QueueLength() int {
	len := Q.rear - Q.top
	return len
}

//入队列,支持动态扩容
func (Q *SeqQueue) enQueue(data string) error {
	if Q.QueueFull() {
		newArray := make([]interface{}, Q.size*2)
		j := 0
		for i := Q.top + 1; i <= Q.rear; i++ {
			newArray[j] = Q.data[i]
			j++
		}
		Q.top = -1
		Q.rear = j - 1
		Q.size = Q.size * 2
		Q.data = newArray
	}
	Q.rear++
	Q.data[Q.rear] = data
	return nil
}

//出队列
func (Q *SeqQueue) deQueue() interface{} {
	if Q.QueueEmpty() {
		log.Fatal("empty Queue!!")
	}
	Q.top++
	result := Q.data[Q.top]
	return result
}

func (Q *SeqQueue) QueueTop() interface{} {
	if Q.QueueEmpty() {
		log.Fatal("empty Queue!!")
	}
	result := Q.data[Q.top+1]
	return result
}

//清空队列
func (Q *SeqQueue) ClearQueue() {
	Q.top = -1
	Q.rear = -1
}

//打印队列
func (Q *SeqQueue) ShowQueue() {
	if Q.QueueEmpty() {
		log.Fatal("Queue empty!")
	}
	start := 0
	if Q.top > 0 {
		start = Q.top + 1
	}
	for i := start; i <= Q.rear; i++ {
		fmt.Printf("array[%d]=%d \t\n", i, Q.data[i])
	}
}

func main() {
	var ss ArrayQueue
	ss = &SeqQueue{}
	ss.InitQueue(5)
	ss.enQueue("one")
	ss.enQueue("two")
	ss.enQueue("333")
	ss.enQueue("444")
	ss.ShowQueue()
	fmt.Println("deQueue:", ss.deQueue())
	fmt.Println("deQueue:", ss.deQueue())
	ss.ShowQueue()
	fmt.Println("len:", ss.QueueLength())
	fmt.Println("top:", ss.QueueTop())
	ss.enQueue("5555")
	ss.enQueue("6666")
	ss.enQueue("7777")
	ss.enQueue("88888")
	ss.ShowQueue()
	fmt.Printf("%+v \n", ss)
	fmt.Println("deQueue:", ss.deQueue())
	fmt.Println("deQueue:", ss.deQueue())
	fmt.Println("deQueue:", ss.deQueue())
	fmt.Printf("%+v \n", ss)
	ss.ShowQueue()
	fmt.Println("top:", ss.QueueTop())
}
