/**
* @Author: XJX
* @Description: 红黑树的实现
* @File: rb-tree.go
* @Date: 2020/12/15 15:49
 */

package main

/**
 * @Description: 定义红黑树的颜色
 * @return nc
 */
const (
	Red   = 1 //红色
	Black = 2 //黑色
)

/**
 * @Description:  定义红黑树节点
 */
type RBNode struct {
	Color               bool    //颜色
	Left, Right, Father *RBNode //左右和父节点
	Val                 int     //值
}

/**
 * @Description: 定义红黑树结构
 */
type RBTree struct {
	root *RBNode
}

/**
 * @Description: 红黑树左旋
 *
 *          node              tr(r)
 *		   /  \              /  \
 *        tl   tr(r)  =>    node rr
 *            /  \         /  \
 *           rl   rr      tl   rl
 *
 * @receiver t 红黑树
 * @param node 红黑树节点
 * @return *RBNode 返回红黑树
 */
func (t *RBTree) leftRotate(node *RBNode) {
	if node != nil {
		rootNode := node.Right
		rootNode.Father = node.Father
		if node.Father != nil {
			if node.Father.Left == node {
				node.Father.Left = rootNode
			} else {
				node.Father.Right = rootNode
			}
		} else {
			t.root = rootNode
		}
		node.Father = rootNode
		node.Right = rootNode.Left
		if node.Right != nil {
			node.Right.Father = node
		}
		rootNode.Left = node
	}
}

/**
 * @Description: 红黑树右旋
 *
 *          node             tl(r)
 *		   /  \              /  \
 *      (r)tl  tr  =>       ll   node
 *       /  \                   / \
 *      ll   lr               lr   tr
 *
 * @receiver t 红黑树
 * @param node 红黑树节点
 * @return *RBNode 返回旋转后的树
 */
func (t *RBTree) rightRotate(node *RBNode) {
	if node != nil {
		rootNode := node.Left
		rootNode.Father = node.Father
		if node.Father != nil {
			if node.Father.Left == node {
				node.Father.Left = rootNode
			} else {
				node.Father.Right = rootNode
			}
		} else {
			t.root = rootNode
		}
		node.Father = rootNode
		node.Left = rootNode.Right
		if node.Left != nil {
			node.Left.Father = node
		}
		rootNode.Right = node
	}
}

func main() {

}
