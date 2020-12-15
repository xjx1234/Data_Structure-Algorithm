/**
* @Author: XJX
* @Description: 红黑树的实现
* @File: rb-tree.go
* @Date: 2020/12/15 15:49
 */

package main

/**
 * @Description:  定义红黑树节点
 */
type RBNode struct {
	Color               bool
	Left, Right, Father *RBNode
	Val                 int
}

func main() {

}
