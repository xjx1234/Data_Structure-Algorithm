/**
* @Author: XJX
* @Description: 阶乘的递归解决
* 题: 阶乘即一个数从 n * (n-1) * (n-2) * ... * 1;用数学表示为 n!
* @File: factorial.go
* @Date: 2020/9/10 16:18
 */

package main

import "fmt"

func getFactorial(n int) int {
	var ret int
	if n == 1 {
		ret = 1
	} else if n > 1 {
		ret = n * getFactorial(n-1)
	}
	return ret
}

func main() {
	fmt.Println(getFactorial(6))
}
