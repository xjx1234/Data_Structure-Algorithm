/**
* @Author: XJX
* @Description: 设计跳表
* @File: skiplist.go
* @Date: 2020/9/29 14:43
 */

package main

import (
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

type SkipList struct {
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
func Constructor() SkipList {
	return SkipList{head: newNode(0, MaxLevel), level: 1}
}

/** 添加跳表数据 */
func (slist *SkipList) Add(num int) {
	current := slist.head
	for i := slist.level - 1; i >= 0; i-- {

	}
}

func main() {

}
