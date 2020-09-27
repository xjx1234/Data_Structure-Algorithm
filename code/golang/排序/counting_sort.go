/**
* @Author: XJX
* @Description: 计数排序实现
* @File: counting_sort.go
* @Date: 2020/9/25 15:27
 */

package main

import "fmt"

func counting_sort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	//取最大值
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}

	countArr := make([]int, max+1)
	sortedArr := make([]int, len(data))

	//统计每个数字出现的次数
	for _, v := range data {
		countArr[v]++
	}

	//统计数字累加和
	for i := 1; i <= max; i++ {
		countArr[i] += countArr[i-1]
	}

	//每个元素找到其位置
	for _, v := range data {
		sortedArr[countArr[v]-1] = v
		countArr[v]--
	}
	return sortedArr
}

func main() {
	data := []int{5, 1, 2, 3, 4, 5, 6, 7, 8, 8, 11}
	d := counting_sort(data)
	fmt.Println(d)
}
