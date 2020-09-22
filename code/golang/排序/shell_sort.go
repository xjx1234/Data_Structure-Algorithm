/**
* @Author: XJX
* @Description: 希尔排序实现
* @File: shell_sort.go
* @Date: 2020/9/22 11:15
 */

package main

import (
	"fmt"
	"math/rand"
)

/** 希尔排序实现*/
func shell_sort(data []int) {
	len := len(data)
	if len <= 1 {
		return
	}
	h := int(len / 3) + 1
	for h >= 1 {
		for i := h; i < len; i++ {
			for j := i; j >= h && data[j] < data[j-h]; j -= h {
				fmt.Println("index", i, j, j-h, h)
				fmt.Println("data", data[j], data[j-h])
				swap(data, j, j-h)
			}
		}
		fmt.Println("data"+string(h)+":", data)
		h = int(h / 3)
	}
}

func swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}

func main() {
	var data []int
	for i := 0; i < 11; i++ {
		data = append(data, rand.Intn(100))
	}
	fmt.Println("原始数据:", data)
	shell_sort(data)
	fmt.Println("希尔排序后数据:", data)
}
