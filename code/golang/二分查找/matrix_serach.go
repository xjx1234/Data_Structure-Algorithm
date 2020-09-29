/**
* @Author: XJX
* @Description: 排序矩阵查找实现
* @题目: 给定M×N矩阵，每一行、每一列都按升序排列，请编写代码找出某元素。
* @File: matrix_serach.go
* @Date: 2020/9/28 17:11
 */

package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	flag := false
	for i := 0; i < len(matrix); i++ {
		mid := searchData(matrix[i], target)
		if mid != -1 {
			flag = true
			break
		}
	}
	return flag
}

func searchData(data []int, target int) int {
	mid := -1
	len := len(data)
	if len > 0 {
		if target < data[0] || target > data[len-1] {
			return -1
		} else if len == 1 && data[0] == target {
			return 0
		} else {
			low := 0
			high := len - 1
			for low <= high {
				mid = (low + high) >> 1
				if data[mid] == target {
					return mid
				} else if data[mid] > target {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
		}
	}
	return -1
}

func main() {
	data := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	flag := searchMatrix(data, 13)
	fmt.Println("查询结果:", flag)
}
