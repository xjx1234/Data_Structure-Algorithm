/**
* @Author: XJX
* @Description: 单链表的实现
* @File: SingleListNode.go
* @Date: 2020/8/28 15:02
 */
package main

import (
	"fmt"
)

//定义一个单链表节点
type SingleLinkNode struct {
	Data interface{}
	Next *SingleLinkNode
}

//定义一个单链表
type SingleLinkList struct {
	Head *SingleLinkNode
	Size int
}

//初始化单链表
func (L *SingleLinkList) InitLink() *SingleLinkList {
	L.Size = 0
	L.Head = new(SingleLinkNode)
	return L
}

//获取单链表长度
func (L *SingleLinkList) GetSize() int {
	return L.Size
}

//返回单链表是否为空
func (L *SingleLinkList) IsEmpty() bool {
	return L.Size == 0
}

func (L *SingleLinkList) Print() {
	if L.Size == 0 {
		fmt.Println("单链表为空")
	}
}

/*
	指定节点序号插入单链表新节点
	虽然插入操作时间复杂度为O(1),但由于需要遍历index位置,所以插入最终操作的时间为复杂度为O(n).
*/
func (L *SingleLinkList) InsertNode(value interface{}, index int) {
	if index > L.Size || index < 0 {
		panic("index out of range")
	}
	preNode := L.Head
	for i := 0; i < index; i++ {
		preNode = preNode.Next
	}
	curNode := &SingleLinkNode{
		Data: value,
		Next: preNode.Next,
	}
	preNode.Next = curNode
	L.Size++
}

func main() {

	linkList := new(SingleLinkList)
	linkList.InitLink()

}
