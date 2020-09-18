/**
* @Author: XJX
* @Description: 插入排序算法实现
* @File: insertion_sort.go
* @Date: 2020/9/18 16:43
 */

package main

import "fmt"

//插入排序算法
func insertionSort(data []int) {
	len := len(data)
	if len <= 1 {
		return
	}
	for i := 1; i < len; i++ {
		value := data[i]
		j := i - 1
		for ; j >= 0; j-- {
			if data[j] > value {
				data[j+1] = data[j] //移动数据
			} else {
				break
			}
		}
		data[j+1] = value //插入数据
	}
}

func main() {
	data := []int{9, 6, 5, 8, 2, 1}
	fmt.Println("原始:", data)
	insertionSort(data)
	fmt.Println("排序:", data)
}
