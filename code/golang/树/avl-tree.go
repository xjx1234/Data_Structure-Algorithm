/**
* @Author: XJX
* @Description: 二叉平衡搜索树的实现
* @File: avl-tree.go
* @Date: 2020/11/20 15:43
 */

package main

import "fmt"

/** 定义平衡二叉树节点 */
type AVLNode struct {
	Left, Right *AVLNode
	Val         int
	Height      int
}

/**
 * @Description:  遍历所有节点
 * @receiver t 节点
 * @return []int 返回节点数组
 */
func (t *AVLNode) getAll() []int {
	values := []int{}
	return addValues(values, t)

}

func addValues(values []int, t *AVLNode) []int {
	if t != nil {
		values = addValues(values, t.Left)
		values = append(values, t.Val)
		fmt.Println(t.Val, t.Height)
		values = addValues(values, t.Right)
	}
	return values
}

/**
 * @Description: 查询相关值节点
 * @receiver t avinode节点
 * @param val 查找值
 * @return *AVLNode 返回查询值
 */
func (t *AVLNode) Serach(val int) *AVLNode {
	if t == nil {
		return nil
	}
	node := &AVLNode{}
	if t.Val == val {
		return t
	} else if t.Val > val {
		node = t.Left.Serach(val)
	} else {
		node = t.Right.Serach(val)
	}
	return node
}

/**
 * @Description: 添加节点
 * @receiver t 节点
 * @param val 添加值
 * @return *AVLNode 返回添加好后的节点
 */
func (t *AVLNode) insertNode(val int) *AVLNode {
	if t == nil {
		newNode := &AVLNode{nil, nil, val, 1}
		return newNode
	}
	if val < t.Val {
		t.Left = t.Left.insertNode(val)
		t = t.TreeBalance()
	} else if val > t.Val {
		t.Right = t.Right.insertNode(val)
		t = t.TreeBalance()
	}
	t.Height = max(t.Left.getHeight(), t.Right.getHeight()) + 1
	return t
}

/**
 * @Description: 删除节点
 * @receiver t 操作节点信息
 * @param val 删除值
 * @return *AVLNode 返回处理后的节点信息
 */
func (t *AVLNode) deleteNode(val int) *AVLNode {
	if t == nil {
		return t
	}
	if val < t.Val {
		t.Left = t.Left.deleteNode(val)
	} else if val > t.Val {
		t.Right = t.Right.deleteNode(val)
	} else {
		if t.Left != nil && t.Right != nil {
			curNode := t.Left
			for curNode.Right != nil {
				curNode = curNode.Right
			}
			t.Val, curNode.Val = curNode.Val, t.Val
			t.Left.deleteNode(curNode.Val)
		} else if t.Left != nil {
			t = t.Left
		} else {
			t = t.Right
		}
	}
	if t != nil {
		t.Height = max(t.Left.getHeight(), t.Right.getHeight()) + 1
		t = t.TreeBalance()
	}
	return t
}

/**
 * @Description: 判断是否要自平衡旋转
 * @receiver t 节点信息
 * @return *AVLNode 返回旋转后节点信息
 */
func (t *AVLNode) TreeBalance() *AVLNode {
	if t.Right.getHeight()-t.Left.getHeight() == 2 { //右子树高度大于左子数高度2
		if t.Right.Right.getHeight() > t.Right.Left.getHeight() {
			t = t.leftRotate()
		} else {
			t = t.rightThenLeftRotate()
		}
	} else if t.Left.getHeight()-t.Right.getHeight() == 2 { //左子树高度大于右子数高度2
		if t.Left.Left.getHeight() > t.Left.Right.getHeight() {
			t = t.rightRotate()
		} else {
			t = t.LeftThenRightRotate()
		}
	}
	return t
}

/**
 * @Description:  左旋转
 * @receiver t avinode节点
 * @return *AVLNode 返回旋转后的节点
 */
func (t *AVLNode) leftRotate() *AVLNode {
	rootNode := t.Right
	t.Right = rootNode.Left
	rootNode.Left = t
	t.Height = max(t.Left.getHeight(), t.Right.getHeight()) + 1
	rootNode.Height = max(rootNode.Left.getHeight(), rootNode.Right.getHeight()) + 1
	return rootNode
}

/**
 * @Description: 右旋转
 * @receiver t  t avinode节点
 * @return *AVLNode 返回旋转后的节点
 */
func (t *AVLNode) rightRotate() *AVLNode {
	rootNode := t.Left
	t.Left = rootNode.Right
	rootNode.Right = t
	t.Height = max(t.Left.getHeight(), t.Right.getHeight()) + 1
	rootNode.Height = max(rootNode.Left.getHeight(), rootNode.Right.getHeight()) + 1
	return rootNode
}

/**
 * @Description:  先左旋再右旋
 * @receiver t avinode节点
 * @return *AVLNode 返回旋转后的节点
 */
func (t *AVLNode) LeftThenRightRotate() *AVLNode {
	sonRootNode := t.Left.leftRotate()
	t.Left = sonRootNode
	return t.rightRotate()
}

/**
 * @Description: 先右旋再左旋
 * @receiver t avinode节点
 * @return *AVLNode 返回旋转后的节点
 */
func (t *AVLNode) rightThenLeftRotate() *AVLNode {
	sonRootNode := t.Right.rightRotate()
	t.Right = sonRootNode
	return t.leftRotate()
}

/**
 * @Description:  比较大小
 * @param a 参数a值
 * @param b 参数b值
 * @return int 返回最大值
 */
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/**
 * @Description:  获取AVI节点高度
 * @receiver t 接收AVLNODE节点
 * @return int 返回高度
 */
func (t *AVLNode) getHeight() int {
	if t == nil {
		return 0
	}
	return t.Height
}

/**
 * @Description: 获取最小值
 * @receiver t 接收AVLNODE节点
 * @return int 返回最小值
 */
func (t *AVLNode) getMin() int {
	if t == nil {
		return -1
	}
	cur := t.Left
	for cur != nil {
		cur = cur.Left
	}
	return cur.Val
}

/**
 * @Description: 递归获取最大值
 * @receiver t 接收AVLNODE节点
 * @return int 返回最大值
 */
func (t *AVLNode) getMax() int {
	if t == nil {
		return -1
	}
	if t.Right == nil {
		return t.Val
	} else {
		return t.Right.getMax()
	}
}

func main() {
	root := &AVLNode{nil, nil, 10, 1}
	root.insertNode(7)
	root.insertNode(15)
	root.insertNode(20)
	root.insertNode(24)
	data := root.getAll()
	fmt.Println(data)
	root.deleteNode(24)
	data1 := root.getAll()
	fmt.Println(data1)
}
