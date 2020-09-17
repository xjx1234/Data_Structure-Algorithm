/**
* @Author: XJX
* @Description: 冒泡排序实现
* @File: bubble_sort.go
* @Date: 2020/9/17 17:17
 */

package main

import "fmt"

//冒泡排序算法
func bubbleSort(data []int) {
	arrLen := len(data)
	if arrLen <= 1 {
		return
	}
	for i := 0; i < arrLen; i++ {
		isChange := false //是否有数据交换
		for j := 0; j < arrLen-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				isChange = true
			}
		}
		//无数据交换 提交结束
		if isChange == false {
			break
		}
	}
}

func main() {
	data := []int{9, 6, 7, 5, 8, 2, 1}
	fmt.Println("原始：", data)
	bubbleSort(data)
	fmt.Println("排序后：", data)
	data2 := []int{10, 11, 12, 9, 6, 7, 5, 8, 2, 1}
	bubbleSort(data2)
	fmt.Println("排序后：", data2)
}
