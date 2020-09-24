/**
* @Author: XJX
* @Description: 归并排序实现
* @File: merge_sort
* @Date: 2020/9/22 13:55
 */

package main

import "fmt"

/** 并归核心算法 */
func merge_sort(data []int) []int {
	len := len(data)
	if len <= 1 {
		return data
	}
	middle := len / 2
	left := merge_sort(data[:middle])
	right := merge_sort(data[middle:])
	return merge(left, right)
}

/** 合并 */
func merge(left, right []int) []int {
	left_len := len(left)
	right_len := len(right)
	i, j := 0, 0
	temp := make([]int, 0)
	for i < left_len && j < right_len {
		if left[i] < right[j] {
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			j++
		}
	}

	if i < left_len {
		temp = append(temp, left[i:]...)
	} else if j < right_len {
		temp = append(temp, right[j:]...)
	}
	return temp
}

func main() {
	data := []int{8, 4, 5, 7, 1, 3, 6, 2}
	fmt.Println("并归排序前:", data)
	result := merge_sort(data)
	fmt.Println("并归排序后:", result)
}
