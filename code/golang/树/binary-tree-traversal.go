/**
* @Author: XJX
* @Description:
	LeetCode 144: 二叉树的前序遍历
	LeetCode 94: 二叉树的中序遍历
	LeetCode 145: 二叉树的后序遍历
* @File: binary-tree-traversal.go
* @Date: 2020/11/13 15:44
*/

package main

/** 二叉树节点结构体定义 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
*	递归前序遍历二叉树    本节点 -> 左子树 -> 右子树
*	时间复杂度：O(n)O(n)，其中 nn 是二叉树的节点数。每一个节点恰好被遍历一次
*	空间复杂度：O(n)O(n)，为递归过程中栈的开销，平均情况下为 O(\log n)O(logn)，最坏情况下树呈现链状，为 O(n)O(n)。
 */
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		res = append(res, preorderTraversal(root.Left)...)
		res = append(res, preorderTraversal(root.Right)...)
	}
	return res
}

/**
*	递归中序遍历二叉树  左子树 -> 本节点 -> 右子树
*	时间复杂度：O(n)O(n)，其中 nn 是二叉树的节点数。每一个节点恰好被遍历一次
*	空间复杂度：O(n)O(n)，为递归过程中栈的开销，平均情况下为 O(\log n)O(logn)，最坏情况下树呈现链状，为 O(n)O(n)。
 */
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, inorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

/**
*	递归后序遍历二叉树  左子树  -> 右子树 -> 本节点
*	时间复杂度：O(n)O(n)，其中 nn 是二叉树的节点数。每一个节点恰好被遍历一次
*	空间复杂度：O(n)O(n)，为递归过程中栈的开销，平均情况下为 O(\log n)O(logn)，最坏情况下树呈现链状，为 O(n)O(n)。
 */
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, postorderTraversal(root.Left)...)
		res = append(res, postorderTraversal(root.Right)...)
		res = append(res, root.Val)
	}
	return res
}
