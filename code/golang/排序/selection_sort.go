/**
* @Author: XJX
* @Description: 选择排序实现
* @File: selection_sort.go
* @Date: 2020/9/21 16:39
 */

package main

import "fmt"

/** 按升序排序 */
func selectionDesSort(data []int) {
	len := len(data)
	if len <= 1 {
		return
	}
	minIdx := 0
	for i := 0; i < len; i++ {
		minIdx = i
		k := i + 1
		for ; k < len; k++ {
			if data[minIdx] > data[k] {
				minIdx = k
			}
		}
		if minIdx != i {
			data[i], data[minIdx] = data[minIdx], data[i]
		}
	}
}

/** 按降序排序 */
func selectionAscSort(data []int) {
	len := len(data)
	if len <= 1 {
		return
	}
	minIdx := 0
	for i := 0; i < len; i++ {
		minIdx = i
		k := i + 1
		for ; k < len; k++ {
			if data[minIdx] < data[k] {
				minIdx = k
			}
		}
		if minIdx != i {
			data[i], data[minIdx] = data[minIdx], data[i]
		}
	}
}

func main() {
	data := []int{5, 1, 4, 6, 4, 9}
	selectionDesSort(data)
	fmt.Println(data)
	data2 := []int{5, 1, 4, 6, 4, 9}
	selectionAscSort(data2)
	fmt.Println(data2)
}
