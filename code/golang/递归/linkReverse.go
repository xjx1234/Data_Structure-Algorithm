/**
* @Author: XJX
* @Description: 单链表递归反转实现
* @File: linkReverse.go
* @Date: 2020/9/10 18:44
 */

package main

import "fmt"

//构造链表节点结构体
type Node struct {
	next *Node
	data interface{}
}

//反转核心函数
func reverse(node *Node) *Node {
	if node.next == nil {
		return node
	} else {
		newHead := reverse(node.next)
		node.next.next = node
		node.next = nil
		return newHead
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	head := &Node{data: "head", next: nil}
	current := head

	// 构造链表
	for i := 0; i < len(numbers); i++ {
		next := &Node{
			data: numbers[i],
			next: nil,
		}
		current.next = next
		current = current.next
		current.next = nil
	}

	fmt.Printf("正常链表：")
	p := head
	for true {
		if p == nil {
			break
		} else {
			fmt.Printf("%d->", p.data)
			p = p.next
		}
	}
	fmt.Printf("nil\n")
	fmt.Printf("反转链表：")
	head = reverse(head)
	p = head
	for true {
		if p == nil {
			break
		} else {
			fmt.Printf("%d->", p.data)
			p = p.next
		}
	}
	fmt.Printf("nil\n")
}
