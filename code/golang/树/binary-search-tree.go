/**
* @Author: XJX
* @Description:
	LeetCode 98  : 验证二叉搜索树
	LeetCode 700 : 二叉搜索树中的搜索
	LeetCode 701 : 二叉搜索树中的插入操作
	LeetCode 450 : 二叉搜索树中的节点删除操作
* @File: binary-search-tree.go
* @Date: 2020/11/18 16:36
*/

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * @Description:  验证是否为查找二叉树
 * @receiver root 根节点
 * @return bool 返回是否二叉树
 */
func (root *TreeNode) isValidBST() bool {
	if root == nil {
		return true
	}
	return isBST(root, math.MinInt64, math.MaxInt64)
}

/**
 * @Description: 检测是否符合二叉树函数
 * @param root 根节点
 * @param min 最小值
 * @param max 最大值
 * @return bool 返回是否符合条件
 */
func isBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if min >= root.Val || max <= root.Val {
		return false
	}
	return isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max)
}

/**
 * @Description: 搜索固定值的节点
 * @receiver root 根节点
 * @param val 搜索值
 * @return *TreeNode 返回指定节点或者子树
 */
func (root *TreeNode) searchBST(val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	node := &TreeNode{}
	if root.Val > val {
		node = root.Left.searchBST(val)
	} else {
		node = root.Right.searchBST(val)
	}
	return node
}

/**
 * @Description: 删除二叉搜索树节点
 * @param root 根节点
 * @param key 删除值
 * @return *TreeNode 返回删除后的树
 */
func (root *TreeNode) deleteNode(key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = root.Left.deleteNode(key)
	} else if root.Val < key {
		root.Right = root.Right.deleteNode(key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			nextNode := root.Left
			for nextNode.Right != nil {
				nextNode = nextNode.Right
			}
			nextNode.Right = root.Right
			return root.Right
			//root = root.Left
			/**
			nextNode := root.Right
			for nextNode.Left != nil {
				nextNode = nextNode.Left
			}
			nextNode.Left = root.Left
			root = root.Right
			*/
		}
	}
	return root
}

/**
 * @Description: 增加节点
 * @param root 根节点
 * @param val 值
 * @return *TreeNode 返回增加后的节点
 */
func (root *TreeNode) insertIntoBST(val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{val, nil, nil}
		} else {
			root.Left = root.Left.insertIntoBST(val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{val, nil, nil}
		} else {
			root.Right = root.Right.insertIntoBST(val)
		}
	}
	return root
}

/**
 * @Description:  前序遍历二叉树
 * @receiver nodes 根节点
 * @return []int 返回前序遍历结果
 */
func (nodes *TreeNode) PreOrderTraverse() []int {
	result := []int{}
	if nodes != nil {
		result = append(result, nodes.Val)
		result = append(result, nodes.Left.PreOrderTraverse()...)
		result = append(result, nodes.Right.PreOrderTraverse()...)
	}
	return result
}

func main() {

	root := &TreeNode{33, nil, nil}          //创建节点
	rootLeftNode := &TreeNode{18, nil, nil}  //创建节点
	rootRightNode := &TreeNode{35, nil, nil} //创建节点
	root.Left = rootLeftNode
	root.Right = rootRightNode
	rootLeftNode_1 := &TreeNode{12, nil, nil} //创建节点
	rootLeftNode_2 := &TreeNode{23, nil, nil} //创建节点
	rootLeftNode.Left = rootLeftNode_1
	rootLeftNode.Right = rootLeftNode_2
	rootLeftNode_2_1 := &TreeNode{6, nil, nil}  //创建节点
	rootLeftNode_2_2 := &TreeNode{14, nil, nil} //创建节点
	rootLeftNode_2_3 := &TreeNode{21, nil, nil} //创建节点
	rootLeftNode_2_4 := &TreeNode{25, nil, nil} //创建节点
	rootLeftNode_1.Left = rootLeftNode_2_1
	rootLeftNode_1.Right = rootLeftNode_2_2
	rootLeftNode_2.Left = rootLeftNode_2_3
	rootLeftNode_2.Right = rootLeftNode_2_4
	rootRightNode_1 := &TreeNode{34, nil, nil} //创建节点
	rootRightNode_2 := &TreeNode{38, nil, nil} //创建节点
	rootRightNode.Left = rootRightNode_1
	rootRightNode.Right = rootRightNode_2

	find := root.searchBST(35) //根据值搜索二叉查找树
	fmt.Println(find, find.Val, find.Left.Val, find.Right.Val)

	isOkBST := root.isValidBST() //是否二叉搜索树
	fmt.Println(isOkBST)

	root = root.insertIntoBST(40) //新增值为40的节点
	d := root.PreOrderTraverse()
	fmt.Println(d) //输出结果

	root1 := root.deleteNode(23) //删除值为23的节点
	d1 := root1.PreOrderTraverse()
	fmt.Println(d1) //输出结果
}
