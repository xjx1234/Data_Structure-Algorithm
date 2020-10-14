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

/** 添加跳表元素 */
func (this *Skiplist) Add(num int) {
	current := this.head
	updateNode := make([]*SkipNode, MaxLevel)

	for i := this.level - 1; i >= 0; i-- {
		if current.nextNode[i] == nil || current.nextNode[i].data > num {
			updateNode[i] = current
		} else {
			for current.nextNode[i] != nil && current.nextNode[i].data < num {
				current = current.nextNode[i]
			}
			updateNode[i] = current
		}
	}

	level := randLevel()
	if level > this.level {
		for i := this.level; i < level; i++ {
			updateNode[i] = this.head
		}
		this.level = level
	}

	node := newNode(num, level)
	for i := 0; i < level; i++ {
		node.nextNode[i] = updateNode[i].nextNode[i]
		updateNode[i].nextNode[i] = node
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
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(5)
	s.Add(6)
	s.Add(7)
	s.Add(8)
	fmt.Printf("%v", s.head.nextNode)
	fmt.Println(s.Search(5))
}
