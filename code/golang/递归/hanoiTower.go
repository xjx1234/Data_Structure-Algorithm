/**
* @Author: XJX
* @Description: 汉诺塔递归方案
* @File: hanoiTower.go
* @Date: 2020/9/10 18:45
 */

package main

import "fmt"

func hanota(A []int, B []int, C []int) []int {
	n := len(A)
	hanio(n, &A, &B, &C)
	return C
}

func hanio(n int, a *[]int, b *[]int, c *[]int) {
	if n < 1 || n > 14 {
		return
	} else if n == 1 {
		*c = append(*c, (*a)[len(*a)-1])
		*a = (*a)[:len(*a)-1]
	} else {
		hanio(n-1, a, c, b)
		*c = append(*c, (*a)[len(*a)-1])
		*a = (*a)[:len(*a)-1]
		hanio(n-1, b, a, c)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{}
	c := []int{}
	c = hanota(a, b, c)
	fmt.Printf("%+v", c)
}
