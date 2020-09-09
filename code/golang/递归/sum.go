/**
* @Author: XJX
* @Description: 利用递归方案求和
* 题目: SUM(n) = 1 + 2 + 3 + ... + n
* @File: sum.go
* @Date: 2020/9/9 19:22
 */

package main

import "fmt"

/*
	递归核心函数 sum(n) 实现
	函数调用时间复杂度为 O(1) ，因此递归调用的n次产生的时间复杂度为 O(n)
*/
func sum(n int) int {
	s := 0
	if n == 1 {
		s = s + n
	} else if n > 1 {
		s = n + sum(n-1)
	}
	return s
}

func main() {
	num := sum(10)
	fmt.Println(num)
}
