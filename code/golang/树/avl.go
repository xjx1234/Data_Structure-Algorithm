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

func main() {

}
