/**
* @Author: XJX
* @Description:
* @File: balance-tree.go
* @Date: 2020/11/20 15:43
 */

package main

/** 定义平衡二叉树节点 */
type AVLNode struct {
	Left, Right *AVLNode
	Val         int
	Height      int
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
	if t.Right == nil{
		return t.Val
	}else{
		return t.Right.getMax()
	}
}

func main() {

}
