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

//打印链表信息
func (L *SingleLinkList) Print() {
	if L.Size == 0 {
		fmt.Println("单链表为空")
	}
	format := ""
	p := L.Head.Next
	for p != nil {
		format += fmt.Sprintf("%+v", p.Data)
		if p.Next != nil {
			format += "->"
		}
		p = p.Next
	}
	fmt.Println(format)
}

/*
	指定节点序号插入单链表新节点
	虽然插入操作时间复杂度为O(1),但由于需要遍历index位置,所以插入最终操作的时间为复杂度为O(n).
*/
func (L *SingleLinkList) InsertNode(value interface{}, index int) {
	if index > L.Size || index < 0 {
		panic("index out of range")
	}
	findNode := new(SingleLinkNode)
	if index == 0 && L.Size == 0 {
		findNode = L.Head
	} else {
		findNode = L.FindByIndex(index)
		fmt.Printf("%+v", findNode)
	}
	curNode := &SingleLinkNode{
		Data: value,
		Next: findNode.Next,
	}
	findNode.Next = curNode
	L.Size++
}

/*
	根据序号移除指定节点
	虽然删除操作时间复杂度为O(1),但由于需要遍历index位置,所以删除最终操作的时间为复杂度为O(n)
*/
func (L *SingleLinkList) RemoveNode(index int) {
	if index > L.Size || index < 0 || (L.Size == 0 && index == 0) {
		panic("index out of range")
	}
	if index == 0 {
		if L.Size == 1 {
			L.Head.Next = nil
		} else {
			zeroNode := L.Head.Next
			L.Head.Next = zeroNode.Next
		}
	} else {
		findNode := L.FindByIndex(index)
		nextNode := findNode.Next
		findNode.Next = nextNode.Next
	}
	L.Size--
}

func (L *SingleLinkList) UpdateNode(index int, data interface{}) {
	if index > L.Size || index < 0 || (L.Size == 0 && index == 0) {
		panic("index out of range")
	}
	updateNode := L.FindByIndex(index)
	updateNode.Data = data
}

/*
	根据序号查询指定节点信息
*/
func (L *SingleLinkList) FindByIndex(index int) *SingleLinkNode {
	if index > L.Size || index < 0 || (L.Size == 0 && index == 0) {
		panic("index out of range")
	}
	curNode := L.Head.Next
	for i := 0; i < index; i++ {
		curNode = curNode.Next

	}
	return curNode
}

/*
	根据值查询节点信息
*/
func (L *SingleLinkList) FindByData(data interface{}) (int, *SingleLinkNode) {
	curNode := L.Head.Next
	index := -1
	findNode := new(SingleLinkNode)
	for i := 0; i < L.Size; i++ {
		if curNode.Data == data {
			index = i
			findNode = curNode
			break
		}
		curNode = curNode.Next
	}
	return index, findNode
}

func main() {
	linkList := new(SingleLinkList)
	linkList.InitLink()
	linkList.InsertNode("11111", 0)
	linkList.InsertNode("2222", 1)
	//linkList.InsertNode("3333", 2)
	//linkList.InsertNode("xxxx", 3)
	//linkList.InsertNode("zzzzzz", 4)
	linkList.Print()
	/*linkList.RemoveNode(0)
	linkList.Print()
	linkList.RemoveNode(1)
	linkList.Print()
	index, _ := linkList.FindByData("zzzzzz")
	fmt.Println(index)
	//fmt.Printf("%+v", node)
	linkList.UpdateNode(0, "hello")
	linkList.Print()*/

}
