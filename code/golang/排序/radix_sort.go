/**
* @Author: XJX
* @Description: 基数排序实现
* @File: radix_sort
* @Date: 2020/9/27 15:09
 */

package main

import "fmt"

func radix_sort(data []int) []int {
	len := len(data)
	if len <= 1 {
		return data
	}

	//获取最大值
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}

	//获取最大数位数
	maxDigit := 0
	for max > 0 {
		max = max / 10
		maxDigit++
	}

	var digit int //余数
	divisor := 1 //根据位数每次除的数 1 10 100依次类推
	bucket := [10][20]int{} //桶方案存放数据
	count := [10]int{}
	for i := 1; i <= maxDigit; i++ {
		for j := 0; j < len; j++ {
			tmp := data[j]
			digit = (tmp / divisor) % 10
			bucket[digit][count[digit]] = tmp //根据余数存放数据
			count[digit]++
		}

		k := 0
		for x := 0; x < 10; x++ {
			if count[x] == 0 {
				continue
			}
			for y := 0; y < count[x]; y++ {
				data[k] = bucket[x][y]
				k++
			}
			count[x] = 0
		}
		divisor = divisor * 10

	}
	return data
}

func main() {
	data := []int{49, 38, 65, 1, 0, 100, 101, 22, 43, 54, 76, 31}
	fmt.Println("排序前:", data)
	d := radix_sort(data)
	fmt.Println("排序后:", d)
}
