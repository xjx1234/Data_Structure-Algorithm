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
	Color               int8    //颜色
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
 * @Description: 判断节点是否为黑色节点
 * @param n 红黑树
 * @return bool 是否为黑节点
 */
func isBlack(n *RBNode) bool {
	if n == nil || n.Color == Black {
		return true
	} else {
		return false
	}
}

/**
 * @Description: 设置颜色
 * @param n 红黑树
 * @param color 所需要设置的颜色
 */
func setColor(n *RBNode, color int8) {
	if n == nil {
		return
	}
	n.Color = color
}

/**
 * @Description: 搜索值
 * @receiver t 红黑树
 * @param val 搜索的值
 * @return *RBNode 返回查询的节点
 */
func (t *RBTree) Serach(val int) *RBNode {
	n := t.root
	for n != nil {
		if n.Val == val {
			return n
		} else if n.Val > val {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return nil
}

/**
 * @Description: 查找兄弟节点
 * @receiver t 红黑树
 * @param n 需要查找的节点
 * @return bro 返回兄弟节点
 */
func (t *RBTree) SerachBroNode(n *RBNode) (bro *RBNode) {
	if n.Father == nil {
		return nil
	}
	if n.Father.Left == n {
		bro = n.Father.Right
	} else if n.Father.Right == n {
		bro = n.Father.Left
	} else {
		if n.Father.Left == nil {
			bro = n.Father.Right
		} else {
			bro = n.Father.Left
		}
	}
	return bro
}

/**
 * @Description: 红黑树插入操作
 * @receiver t 红黑树
 * @param val 指定插入值
 */
func (t *RBTree) insert(val int) {
	/** 加入根节点为空，则将新添加的值设置为根节点*/
	if t.root == nil {
		t.root = &RBNode{Left: nil, Right: nil, Val: val, Color: Black, Father: nil}
		return
	}
	root := t.root
	inserNode := &RBNode{Left: nil, Right: nil, Val: val, Color: Red, Father: nil}
	for root != nil {
		if val < root.Val {
			root = root.Left
		} else if val > root.Val {
			root = root.Right
		} else {
			return
		}
	}
	inserNode.Father = root
	if val > root.Val {
		root.Right = inserNode
	} else {
		root.Left = inserNode
	}
	t.insertBalance(inserNode)
}

/**
 * @Description: 调整插入的红黑树平衡
	1. 如果父节点为黑色节点，则不需要调整
	2. 如果父节点为红色节点，叔节点也为红色节点，则需要将新增节点父节点与叔节点变为黑色，父节点的父节点变为红色
	3. 如果新增节点在左子树上，分2种情况:
	(1) 如果父节点和新增节点在一条直线上，则父节点变黑色，父节点的父节点变红色，并且以祖父节点右旋
	(2) 如果父节点和新增节点不在一条直线上，先左旋，然后符合3-(1)标准，再按3-(1)处理
	4. 如果新增节点再右子树上，分2种情况:
	(1) 如果父节点和新增节点在一条直线上，则父节点变黑色，父节点的父节点变红色,并且以祖父节点左旋
	(2) 如果父节点和新增节点不在一条直线上，先右旋，然后符合4-(1)标准，再按4-(1)处理
 * @receiver t 红黑树
 * @param n 最新插入节点
*/
func (t *RBTree) insertBalance(n *RBNode) {
	if !isBlack(n.Father) { //父节点如果是黑色节点，则不需要调整
		uncleNode := t.SerachBroNode(n)
		if !isBlack(uncleNode) {
			n.Father.Color = Black
			uncleNode.Color = Black
			uncleNode.Father.Color = Red
			n = n.Father.Father
		} else if n.Father == n.Father.Father.Left {
			if n == n.Father.Left {
				n.Father.Color = Black
				n.Father.Father.Color = Red
				n = n.Father.Father
				t.rightRotate(n)
			} else {
				n = n.Father
				t.leftRotate(n)
			}
		} else {
			if n.Father == n.Father.Father.Right {
				n.Father.Color = Black
				n.Father.Father.Color = Red
				n = n.Father.Father
				t.leftRotate(n)
			} else {
				n = n.Father
				t.rightRotate(n)
			}
		}
		t.root.Color = Black
	}
}

func (t *RBTree) deleteNode(n *RBNode) {

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
