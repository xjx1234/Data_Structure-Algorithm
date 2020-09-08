/**
* @Author: XJX
* @Description: 环形队列实现 - 数组结构  解决了数组结构频繁搬迁数据的问题
* @File: CircleQueue.go
* @Date: 2020/9/4 16:26
 */

package main

import (
	"fmt"
	"log"
)

//定义环形队列结构体
type CircleQueue struct {
	size int
	data []interface{}
	head int
	tail int
}

//定义环形队列的接口
type ArrayCircleQueueFun interface {
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
func (CQ *CircleQueue) InitQueue(n int) {
	data := make([]interface{}, n)
	CQ.data = data
	CQ.head = 0
	CQ.tail = 0
	CQ.size = n
}

//检测队列是否为空
func (CQ *CircleQueue) QueueEmpty() bool {
	return CQ.tail == CQ.head
}

/*
	检测队列是否为满
	可以用公式： (tail+1)%size=head 方式检测
*/
func (CQ *CircleQueue) QueueFull() bool {
	return (CQ.tail+1)%CQ.size == CQ.head
}

/*
	返回队列长度
	可用公式:  (tail + size - head) % size
*/
func (CQ *CircleQueue) QueueLength() int {
	return (CQ.tail + CQ.size - CQ.head) % CQ.size
}

//入队列
func (CQ *CircleQueue) enQueue(data string) error {
	if CQ.QueueFull() {
		log.Fatal("Queue full!!!")
	}
	CQ.data[CQ.tail] = data
	CQ.tail = (CQ.tail + 1) % CQ.size
	return nil
}

//出队列
func (CQ *CircleQueue) deQueue() interface{} {
	if CQ.QueueEmpty() {
		log.Fatal("Queue empty!!!")
	}
	val := CQ.data[CQ.head]
	CQ.head = (CQ.head + 1) % CQ.size
	return val
}

//获取队列顶部
func (CQ *CircleQueue) QueueTop() interface{} {
	if CQ.QueueEmpty() {
		log.Fatal("Queue empty!!!")
	}
	val := CQ.data[CQ.head]
	return val
}

//清空队列
func (CQ *CircleQueue) ClearQueue() {
	CQ.head = 0
	CQ.tail = 0
}

//打印队列
func (CQ *CircleQueue) ShowQueue() {
	size := CQ.QueueLength()
	if size == 0 {
		fmt.Println("queue empty")
	}
	index := CQ.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\n", index, CQ.data[index])
		index = (index + 1) % CQ.size // 计算出下一队首索引
	}
	fmt.Println()
}

func main() {
	var s ArrayCircleQueueFun
	s = &CircleQueue{}
	s.InitQueue(8)
	s.enQueue("one")
	s.enQueue("two")
	s.enQueue("three")
	s.enQueue("four")
	s.ShowQueue()
	fmt.Println("deQueue", s.deQueue())
	fmt.Println("deQueue", s.deQueue())
	fmt.Println("deQueue", s.deQueue())
	s.ShowQueue()

}
