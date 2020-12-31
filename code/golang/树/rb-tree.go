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
	if findNode.Left != nil && findNode.Right != nil { //删除的节点有2个子节点,则寻找后继节点，将后继节点得值赋予删除节点后转换为删除叶子节点问题
		succNode := t.findSucceedNode(findNode.Right)
		findNode.Val = succNode.Val
		findNode = succNode
	}

	if findNode.Father == nil { //删除节点父节点为空，则该节点为根节点，设置根节点为空
		t.root = nil
	} else if findNode.Left == nil && findNode.Right == nil { //删除节点为叶子节点
		/** 如果删除节点为黑色节点，则先调整节点*/
		if findNode.Color == Black {
			t.deleteBalance(findNode)
		}
		if findNode.Father != nil {
			if findNode.Father.Left == findNode { //删除节点为左节点，则直接删除左节点
				findNode.Father.Left = nil
			} else if findNode.Father.Right == findNode {
				findNode.Father.Right = nil //删除节点为右节点，则直接删除右节点
			}
			findNode.Father = nil //将父节点只为空
		}
	} else { //左右节点有不为空得情况
		/** 替代节点选左右节点其中一个，优先左节点 */
		replaceNode := findNode.Left
		if replaceNode == nil {
			replaceNode = findNode.Right
		}
		replaceNode.Father = findNode.Father //设置替换节点父亲为被替换节点的父亲
		if findNode.Father == nil {
			t.root = replaceNode
		} else if findNode == findNode.Father.Left { //删除节点为左节点,则将删除节点的父节点的左节点替换为替换节点
			findNode.Father.Left = replaceNode
		} else {
			findNode.Father.Right = replaceNode //删除节点为右节点,则将删除节点的父节点的右节点替换为替换节点
		}
		findNode.Left = nil
		findNode.Right = nil
		findNode.Father = nil
		if findNode.Color == Black {
			t.deleteBalance(replaceNode)
		}
	}
}

/**
 * @Description: 删除后再平衡
 * @receiver t 红黑树
 * @param n 开始平衡的节点
 */
func (t *RBTree) deleteBalance(n *RBNode) {
	if t.root == nil {
		return
	}
	for n != t.root && isBlack(n) {
		if n == n.Father.Left {
			rNode := n.Father.Right
			if !isBlack(rNode) {
				setColor(rNode, Black)
				setColor(rNode.Father, Red)
				t.leftRotate(rNode.Father)
				rNode = n.Father.Right
			}
			if isBlack(rNode.Left) && isBlack(rNode.Right) { //无节点替换，则递归向上寻找平衡
				setColor(rNode, Red)
				n = n.Father
			} else {
				if isBlack(rNode.Right) {
					setColor(rNode.Left, Black)
					setColor(rNode, Red)
					t.rightRotate(rNode)
					rNode = n.Father.Right
				}
				setColor(rNode, n.Father.Color)
				setColor(n.Father, Black)
				setColor(rNode.Right, Black)
				t.leftRotate(n.Father)
				n = t.root
			}
		} else {
			rNode := n.Father.Left
			if !isBlack(rNode) {
				setColor(rNode, Black)
				setColor(rNode.Father, Red)
				t.rightRotate(rNode.Father)
				rNode = n.Father.Left
			}
			if isBlack(rNode.Left) || isBlack(rNode.Right) {
				setColor(rNode, Red)
				n = n.Father
			} else {
				if isBlack(rNode.Left) {
					setColor(rNode.Right, Black)
					setColor(rNode, Red)
					t.leftRotate(rNode)
					rNode = n.Father.Left
				}
				setColor(rNode, n.Father.Color)
				setColor(n.Father, Black)
				setColor(rNode.Left, Black)
				t.rightRotate(n.Father)
				n = t.root
			}
		}
	}
	n.Color = Black //替代节点是红色，则变色，变黑
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
