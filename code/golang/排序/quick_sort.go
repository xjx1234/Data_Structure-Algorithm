/**
* @Author: XJX
* @Description: 快速排序实现
* @File: quick_sort.go
* @Date: 2020/9/22 13:55
 */

package main

import "fmt"

/** 快排实现算法 */
func quick_sort(data []int, begin, end int) {
	if begin >= end {
		return
	}
	pivot := data[begin] //选取第一个为对比值
	pivot_key := begin
	for i := begin + 1; i <= end; i++ {
		if data[i] < pivot {
			data[pivot_key] = data[i]
			data[i] = data[pivot_key+1]
			pivot_key++
		}
	}
	data[pivot_key] = pivot
	quick_sort(data, begin, pivot_key-1)
	quick_sort(data, pivot_key+1, end)
}

func main() {
	data := []int{39, 28, 55, 87, 66, 3, 17, 39}
	fmt.Println("排序前:", data)
	quick_sort(data, 0, len(data)-1)
	fmt.Println("排序后:", data)
}
