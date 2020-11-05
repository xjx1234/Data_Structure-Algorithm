/**
* @Author: XJX
* @Description:
	LeetCode 454
	给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
	为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。
* @File: 4sum-ii.go
* @Date: 2020/11/4 16:01
*/

package main

import "fmt"

func fourSumCount(A []int, B []int, C []int, D []int) int {
	num := 0
	m := make(map[int]int)
	for _, c := range C {
		for _, d := range D {
			m[c+d]++
		}
	}

	for _, a := range A {
		for _, b := range B {
			if v, ok := m[-a-b]; ok {
				num += v
			}
		}
	}
	return num
}

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	num := fourSumCount(A, B, C, D)
	fmt.Println(num)
}
