/**
* @Author: XJX
* @Description: 循环单链表实现
* @File: CircleLinkNode.go
* @Date: 2020/9/2 17:15
 */

package main

import "fmt"

//定义一个循环单链表节点
type CircleLinkNode struct {
	Data interface{}
	Next *CircleLinkNode
}

//定义一个循环单链表
type CircleLinkList struct {
	Head *CircleLinkNode
	Size int
}

//初始化循环单链表
func (C *CircleLinkList) InitLink() *CircleLinkList {
	C.Size = 0
	C.Head = new(CircleLinkNode)
	return C
}

//获取循环单链表长度
func (C *CircleLinkList) GetSize() int {
	return C.Size
}

//返回循环单链表是否为空
func (C *CircleLinkList) IsEmpty() bool {
	return C.Size == 0
}

//打印循环链表信息
func (C *CircleLinkList) Print() {
	if C.Size == 0 {
		fmt.Println("循环单链表为空")
	}
	format := ""
	p := C.Head.Next
	fmt.Printf("%+v %p \n", C.Head, C.Head)
	for p != nil {
		fmt.Printf("%+v %p \n", p, p)
		if p.Next == C.Head {
			format += fmt.Sprintf("%+v", p.Data)
			format += fmt.Sprintf("%+v", "->HEAD")
		} else {
			format += fmt.Sprintf("%+v", p.Data)
		}
		if p.Next != C.Head {
			format += "->"
		} else {
			break
		}
		p = p.Next
	}
	fmt.Println(format)
}

/*
	指定节点序号插入循环单链表新节点
	虽然插入操作时间复杂度为O(1),但由于需要遍历index位置,所以插入最终操作的时间为复杂度为O(n).
*/
func (C *CircleLinkList) InsertNode(value interface{}, index int) {
	if index > C.Size || index < 0 {
		panic("index out of range")
	}
	preNode := new(CircleLinkNode)
	findNode := new(CircleLinkNode)
	if index == 0 {
		preNode = C.Head
		if C.Size != 0 {
			findNode = C.FindByIndex(0)
		}
	} else {
		preNode = C.FindByIndex(index - 1)
		findNode = preNode.Next
	}
	curNode := &CircleLinkNode{
		Data: value,
		Next: findNode,
	}
	if index == C.Size {
		curNode.Next = C.Head
	}
	preNode.Next = curNode
	C.Size++
}

/*
	根据序号移除指定节点
	虽然删除操作时间复杂度为O(1),但由于需要遍历index位置,所以删除最终操作的时间为复杂度为O(n)
*/
func (C *CircleLinkList) RemoveNode(index int) {
	if index > C.Size || index < 0 || (C.Size == 0 && index == 0) {
		panic("index out of range")
	}
	if index == 0 {
		if C.Size == 1 {
			C.Head.Next = nil
		} else {
			zeroNode := C.Head.Next
			C.Head.Next = zeroNode.Next
		}
	} else {
		findNode := C.FindByIndex(index - 1)
		if index == (C.Size - 1) {
			findNode.Next = C.Head
		} else {
			nextNode := findNode.Next
			findNode.Next = nextNode.Next
		}
	}
	C.Size--
}

/*
	根据序号查询指定节点信息
*/
func (C *CircleLinkList) FindByIndex(index int) *CircleLinkNode {
	if index > C.Size || index < 0 || (C.Size == 0 && index == 0) {
		panic("index out of range")
	}
	curNode := C.Head.Next
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}
	return curNode
}

func main() {
	circleLinkList := new(CircleLinkList)
	circleLinkList.InitLink()
	circleLinkList.InsertNode("1", 0)
	circleLinkList.Print()
	circleLinkList.InsertNode("2", 1)
	circleLinkList.Print()
	circleLinkList.InsertNode("2*1", 1)
	circleLinkList.Print()
	circleLinkList.InsertNode("3", 2)
	circleLinkList.Print()
	circleLinkList.RemoveNode(2)
	circleLinkList.Print()
	circleLinkList.RemoveNode(2)
	circleLinkList.Print()
}
