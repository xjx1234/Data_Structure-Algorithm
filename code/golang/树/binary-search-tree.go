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
	"math"
)

type Tree_Node struct {
	Val   int
	Left  *Tree_Node
	Right *Tree_Node
}

/** 验证是否为二叉搜索树 */
func isValidBST(root *Tree_Node) bool {
	if root == nil {
		return true
	}
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *Tree_Node, min, max int) bool {
	if root == nil {
		return true
	}
	if min >= root.Val || max <= root.Val {
		return false
	}
	return isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max)
}

/** 递归搜索固定值的节点信息 */
func searchBST(root *Tree_Node, val int) *Tree_Node {
	if root == nil {
		return nil
	} else if root.Val == val {
		return root
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

/** 添加二叉搜索树节点信息 */
func insertIntoBST(root *Tree_Node, val int) *Tree_Node {
	if root == nil {
		return &Tree_Node{Val: val}
	}
	if root.Val > val {
		if root.Left != nil {
			insertIntoBST(root.Left, val)
		} else {
			root.Left = &Tree_Node{Val: val}
		}
	} else {
		if root.Right != nil {
			insertIntoBST(root.Right, val)
		} else {
			root.Right = &Tree_Node{Val: val}
		}
	}
	return root
}

/**
删除二叉搜索树节点信息
对于第一种情况来说：当前节点没有左子树，直接 return root.right 即可。
对于第二种情况来说：当前节点没有右子树，直接 return root.left 即可。
对于第三种情况来说：当前节点既有左子树又有右子树，
(1) 只需将root的左子树放到root的右子树的最下面的左叶子节点的左子树上即可
(2) 只需将root的右子树放到root的左子树的最下面的左叶子节点的右子树上即可
*/
func deleteNode(root *Tree_Node, key int) *Tree_Node {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		/** 前驱节点替代方案(1) */
		previousNode := root.Left
		for previousNode.Right != nil {
			previousNode = previousNode.Right
		}
		previousNode.Right = root.Right
		root = root.Left

		/** 后驱节点替代方案(2) */
		nextNode := root.Right
		for nextNode.Left != nil {
			nextNode = nextNode.Left
		}
		nextNode.Left = root.Left
		root = root.Right
	}
	return root
}

func main() {

}
