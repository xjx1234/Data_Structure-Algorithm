/**
* @Author: XJX
* @Description: 双向链表实现
* @File: DLinkListNode.go
* @Date: 2020/9/2 11:31
 */

package main

import "fmt"

//定义双向链表节点结构体
type DLinkNode struct {
	Data interface{}
	Next *DLinkNode
	Prev *DLinkNode
}

//定义双向链表结构体
type DLinkList struct {
	Head *DLinkNode
	Size int
}

//初始化双链表
func (D *DLinkList) InitDlink() {
	D.Size = 0
	D.Head = new(DLinkNode)
}

//获取双链表长度
func (D *DLinkList) GetSize() int {
	return D.Size
}

//返回双链表是否为空
func (D *DLinkList) IsEmpty() bool {
	return D.Size == 0
}

//打印双链表
func (D *DLinkList) Print() {
	str := ""
	p := D.Head
	for p != nil {
		str += fmt.Sprint(p.Data)
		p = p.Next
		if p != nil {
			str += fmt.Sprint("=>")
		}
	}
	fmt.Println(str)
}

//根据序号查询节点
func (D *DLinkList) FindByIndex(index int) *DLinkNode {
	if index < 0 || index > D.Size {
		panic("index out of range")
	}
	if D.Size == 0 {
		panic("DlinkList is empty")
	}
	findNode := D.Head.Next
	for i := 0; i < index; i++ {
		findNode = findNode.Next
	}
	return findNode
}

//根据内容查找双链表节点
func (D *DLinkList) FindByData(data interface{}) (int, *DLinkNode) {
	curNode := D.Head.Next
	index := -1
	findNode := new(DLinkNode)
	for i := 0; i < D.Size; i++ {
		if curNode.Data == data {
			index = i
			findNode = curNode
			break
		}
		curNode = curNode.Next
	}
	return index, findNode
}

// 插入双链表节点数据
func (D *DLinkList) InsertDNode(index int, data interface{}) {
	if index < 0 || index > D.Size {
		panic("index out of range")
	}
	prevNode := new(DLinkNode)
	findNode := new(DLinkNode)
	if index == 0 {
		prevNode = D.Head
		if D.Size != 0 {
			findNode = D.FindByIndex(0)
		}
	} else {
		prevNode = D.FindByIndex(index - 1)
		findNode = prevNode.Next
	}
	newNode := &DLinkNode{
		Data: data,
		Next: findNode,
		Prev: prevNode,
	}
	if index == D.Size {
		newNode.Next = nil
	}
	prevNode.Next = newNode
	D.Size++
}

//修改双链表内容
func (D *DLinkList) updateNodeData(index int, data interface{}) {
	if index < 0 || index > D.Size {
		panic("index out of range")
	}
	findNode := D.FindByIndex(index)
	findNode.Data = data
}

//移除双链表节点
func (D *DLinkList) RemoveDNode(index int) {
	if index < 0 || index > D.Size {
		panic("index out of range")
	}
	findNode := D.FindByIndex(index)
	prevNode := findNode.Prev
	nextNode := findNode.Next
	prevNode.Next = nextNode
	nextNode.Prev = prevNode
	findNode.Next = nil
	findNode.Prev = nil
	D.Size--
}

func main() {
	dlink := new(DLinkList)
	dlink.InitDlink()
	dlink.InsertDNode(0, "one")
	dlink.Print()
	dlink.InsertDNode(1, "two")
	dlink.InsertDNode(1, "two-new")
	dlink.Print()
	dlink.RemoveDNode(1)
	dlink.Print()
	index, node := dlink.FindByData("two")
	fmt.Println(index)
	fmt.Printf("%+v", node)
	dlink.updateNodeData(1, "two-two")
	dlink.Print()
}
