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
	for !isBlack(n.Father) { //父节点如果是黑色节点，则不需要调整
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

/**
 * @Description: 删除节点
 * @receiver t 红黑树
 * @param v 需要删除的值
 */
func (t *RBTree) deleteNode(v int) {
	findNode := t.Serach(v)
	if findNode == nil {
		return
	}
	fixColor := findNode.Color
	fixNode := &RBNode{Color: Black, Father: findNode.Father}
	if findNode.Left == nil { //如果该节点为单子树，则该节点左节点为空，则用右节点替换
		t.replaceNode(findNode, findNode.Right)
		if findNode.Right != nil {
			fixNode = findNode.Right
		}
	} else if findNode.Right == nil { //如果该节点为单子树，则该节点右节点为空，则用左节点替换
		t.replaceNode(findNode, findNode.Left)
		if findNode.Left != nil {
			fixNode = findNode.Left
		}
	} else {
		succNode := t.findSucceedNode(findNode.Right) //查找后继节点
		fixColor = succNode.Color

		if succNode.Right == nil {
			if succNode.Father != findNode {
				fixNode = &RBNode{Father: succNode.Father, Color: Black}
			} else {
				fixNode = &RBNode{Father: succNode, Color: Black}
			}
		} else {
			fixNode = succNode.Right
		}
		if succNode.Father != findNode {
			t.replaceNode(succNode, succNode.Right)
			succNode.Right = findNode.Right
			succNode.Right.Father = succNode
		}
		t.replaceNode(findNode, succNode)
		succNode.Left = findNode.Left
		succNode.Left.Father = succNode
		succNode.Color = findNode.Color
	}
	if fixColor == Black {
		t.deleteBalance(fixNode)
	}
}

/**
 * @Description: 删除后再平衡
 * @receiver t 红黑树
 * @param n 开始平衡的节点
 */
func (t *RBTree) deleteBalance(n *RBNode) {

}

/**
 * @Description: 替换节点
 * @receiver t 红黑树
 * @param findNode 当前节点
 * @param replaceNode 需要替换的节点
 */
func (t *RBTree) replaceNode(findNode, replaceNode *RBNode) {
	if findNode.Father == nil { //如果当前节点父节点为空
		t.root = replaceNode //设置替代节点为根节点
		if replaceNode != nil {
			replaceNode.Father = nil
		}
	} else if findNode == findNode.Father.Left {
		findNode.Father.Left = replaceNode //替换当前节点父节点的左节点为替代节点
	} else {
		findNode.Father.Right = replaceNode //替换当前节点父节点的右节点为替代节点
	}
	if replaceNode != nil {
		replaceNode.Father = findNode.Father //将替代节点的父节点设置为当前节点的父节点
	}
}

/**
 * @Description: 查找后继节点
 * @receiver t 红黑树
 * @param n 开始搜索节点
 * @return *RBNode 返回后继节点
 */
func (t *RBTree) findSucceedNode(n *RBNode) *RBNode {
	for n.Left != nil {
		n = n.Left
	}
	return n
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
