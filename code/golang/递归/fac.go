/**
* @Author: XJX
* @Description: 斐波那契数列的实现
* 题： 写一个函数，输入 `n` ，求斐波那契（Fibonacci）数列的第 `n` 项
* @File: fac.go
* @Date: 2020/9/9 19:57
 */

package main

import "fmt"

/**
	斐波那契数列核心递归核心解决方案
	算法的时间复杂度是O(2^n), 递归方案较为费时，可以通过其他辅助方案解决
 */
func fac(n int) int {
	var ret int
	if n > 2 {
		ret = fac(n-1) + fac(n-2)
	} else if n == 1 || n == 2 {
		ret = 1
	}
	//fmt.Println(ret)
	return ret
}

func main() {
	fmt.Println(fac(43))
}
