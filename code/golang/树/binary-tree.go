/**
* @Author: XJX
* @Description:
	二叉树相关操作
	涉及相关leetCode题目
	LeetCode 144. 二叉树的前序遍历
	LeetCode 94. 二叉树的中序遍历
	LeetCode 145. 二叉树的后序遍历
	LeetCode 102. 二叉树的层序遍历
	LeetCode 剑指 Offer 55 - I. 二叉树的深度
	LeetCode 104. 二叉树的最大深度
	LeetCode 226. 翻转二叉树
* @File: binary-tree.go
* @Date: 2020/11/13 15:44
*/

package main

import "fmt"

/** 二叉树节点结构体定义 */
type BiTNode struct {
	Value       int
	Left, Right *BiTNode
}

/** 递归方式前序遍历 */
func (nodes *BiTNode) PreOrderTraverse() []int {
	result := []int{}
	if nodes != nil {
		result = append(result, nodes.Value)
		result = append(result, nodes.Left.PreOrderTraverse()...)
		result = append(result, nodes.Right.PreOrderTraverse()...)
	}
	return result
}

/** 迭代方式前序遍历*/
func (nodes *BiTNode) PreOrderTraverseByIterative() []int {
	var stack []*BiTNode
	result := []int{}
	cur := nodes
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			result = append(result, cur.Value)
			stack = append(stack, cur.Right)
			cur = cur.Left
		}
		index := len(stack) - 1
		cur = stack[index]
		stack = stack[:index]
	}
	return result
}

/** 递归方式中序遍历 */
func (nodes *BiTNode) InOrderTraverse() []int {
	result := []int{}
	if nodes != nil {
		result = append(result, nodes.Left.InOrderTraverse()...)
		result = append(result, nodes.Value)
		result = append(result, nodes.Right.InOrderTraverse()...)
	}
	return result
}

/** 迭代方式中序遍历 */
func (nodes *BiTNode) InOrderTraverseByIterative() []int {
	var stack []*BiTNode
	result := []int{}
	cur := nodes
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Value)
		cur = cur.Right
	}
	return result
}

/** 递归方式后序遍历 */
func (nodes *BiTNode) PostOrderTraverse() []int {
	result := []int{}
	if nodes != nil {
		result = append(result, nodes.Left.PostOrderTraverse()...)
		result = append(result, nodes.Right.PostOrderTraverse()...)
		result = append(result, nodes.Value)
	}
	return result
}

/** 迭代方式后序遍历 */
func (nodes *BiTNode) PostOrderTraverseByIterative() []int {
	var stack []*BiTNode
	var prev *BiTNode
	result := []int{}
	cur := nodes
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Right == nil || cur.Right == prev {
			result = append(result, cur.Value)
			prev = cur
			cur = nil
		} else {
			stack = append(stack, cur)
			cur = cur.Right
		}
	}
	return result
}

/**
*	层级遍历
*   时间复杂度：O(n),空间复杂度：O(n)
 */
func (nodes *BiTNode) LevelOrder() []int {
	var stack []*BiTNode
	result := []int{}
	stack = append(stack, nodes)
	for j := 0; j < len(stack); j++ {
		node := stack[j]
		result = append(result, node.Value)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return result
}

/** 二叉树的深度 */
func (nodes *BiTNode) maxDepth() int {
	if nodes == nil {
		return 0
	}
	if nodes.Left == nil && nodes.Right == nil {
		return 1
	}
	leftLevel := nodes.Left.maxDepth()
	rightLevel := nodes.Right.maxDepth()
	if leftLevel > rightLevel {
		return 1 + leftLevel
	} else {
		return 1 + rightLevel
	}
}

/**
* 反转二叉树
 */
func invertTree(root *BiTNode) *BiTNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Right = left
	root.Left = right
	return root
}

func main() {

	/**
	创建二叉树
				0
	           / \
			  1	  2
	         / \ / \
	        3  4 5  6
	*/
	root := &BiTNode{0, nil, nil}          //创建节点
	rootLeftNode := &BiTNode{1, nil, nil}  //创建节点
	rootRightNode := &BiTNode{2, nil, nil} //创建节点
	root.Left = rootLeftNode               //赋值左节点
	root.Right = rootRightNode             //赋值右节点
	leftNode1 := &BiTNode{3, nil, nil}     //创建节点
	rightNode1 := &BiTNode{4, nil, nil}    //创建节点
	rootLeftNode.Left = leftNode1          //赋值左节点
	rootLeftNode.Right = rightNode1        //赋值右节点
	leftNode2 := &BiTNode{5, nil, nil}     //创建节点
	rightNode2 := &BiTNode{6, nil, nil}    //创建节点
	rootRightNode.Left = leftNode2         //赋值左节点
	rootRightNode.Right = rightNode2       //赋值右节点

	pre := root.PreOrderTraverse() //递归前序遍历
	fmt.Println("递归前序遍历顺序:", pre)
	in := root.InOrderTraverse() //递归中序遍历
	fmt.Println("递归中序遍历顺序:", in)
	post := root.PostOrderTraverse() //递归后序遍历
	fmt.Println("递归后序遍历顺序:", post)

	pre2 := root.PreOrderTraverseByIterative() //迭代前序遍历
	fmt.Println("迭代前序遍历顺序:", pre2)
	in2 := root.InOrderTraverseByIterative() //迭代中序遍历
	fmt.Println("迭代中序遍历顺序:", in2)
	post2 := root.PostOrderTraverseByIterative() //迭代后序遍历
	fmt.Println("迭代后序遍历顺序:", post2)

	lev := root.LevelOrder() //分层遍历
	fmt.Println("分层遍历:", lev)

	num := root.maxDepth()
	fmt.Println("最大深度:", num)

	tree := invertTree(root) //二叉树反转
	fmt.Println(tree)

}
