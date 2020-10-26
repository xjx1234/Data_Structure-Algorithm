/**
* @Author: XJX
* @Description: 设计跳表
* @File: skiplist.go
* @Date: 2020/9/29 14:43
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxLevel    = 32
	Probability = 0.25
)

type SkipNode struct {
	data     int
	nextNode []*SkipNode
}

type Skiplist struct {
	level int
	head  *SkipNode
}

/** 产生随机层级 */
func randLevel() (level int) {
	rand.Seed(time.Now().UnixNano())
	for level = 1; rand.Float32() < Probability && level < MaxLevel; level++ {
	}
	return
}

/** 创建跳表节点 */
func newNode(data int, level int) *SkipNode {
	return &SkipNode{data: data, nextNode: make([]*SkipNode, level)}
}

/** 初始哈跳表 */
func Constructor() Skiplist {
	return Skiplist{head: newNode(0, MaxLevel), level: 1}
}

/** 搜索跳表元素 */
func (this *Skiplist) Search(target int) bool {
	current := this.head
	for i := this.level - 1; i >= 0; i-- {
		for current.nextNode[i] != nil {
			if current.nextNode[i].data == target {
				return true
			} else if current.nextNode[i].data > target {
				break
			} else {
				current = current.nextNode[i]
			}
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	currentNode := this.head
	updateNode := make([]*SkipNode, MaxLevel)

	for i := this.level - 1; i >= 0; i-- {
		if currentNode.nextNode[i] == nil || currentNode.nextNode[i].data > num {
			updateNode[i] = currentNode
		} else {
			for currentNode.nextNode[i] != nil && currentNode.nextNode[i].data < num {
				currentNode = currentNode.nextNode[i]
			}
			updateNode[i] = currentNode
		}
	}

	level := randLevel()
	if level > this.level {
		for l := this.level; l < level; l++ {
			updateNode[l] = this.head
		}
		this.level = level
	}
	newNode := newNode(num, MaxLevel)

	for j := 0; j < this.level; j++ {
		newNode.nextNode[j] = updateNode[j].nextNode[j]
		updateNode[j].nextNode[j] = newNode
	}

}

/** 删除跳表元素 */
func (this *Skiplist) Erase(num int) bool {
	current := this.head
	flag := false
	for i := this.level - 1; i >= 0; i-- {
		for current.nextNode[i] != nil {
			if current.nextNode[i].data == num {
				tmp := current.nextNode[i]
				current.nextNode[i] = tmp.nextNode[i]
				tmp.nextNode[i] = nil
				flag = true
				break
			} else if current.nextNode[i].data > num {
				break
			} else {
				current = current.nextNode[i]
			}
		}
	}
	return flag
}

func main() {
	s := Constructor()
	s.Add(10)
	s.Add(3)
	fmt.Printf("%v", s.head.nextNode)
	fmt.Println(s.Search(5))
}
