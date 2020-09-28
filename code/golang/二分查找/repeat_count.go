/**
* @Author: XJX
* @Description: 在有序数组中，查找某个元素出现的次数
* @题目: 统计一个数字在排序数组中出现的次数。
* @File: repeat_count.go
* @Date: 2020/9/28 16:27
 */

package main

import "fmt"

func search(nums []int, target int) int {
	count := 0
	low := 0
	mid := 0
	high := len(nums) - 1
	for low <= high {
		mid = (low + high) >> 1
		if nums[mid] == target {
			count++
			break
		} else if nums[mid] > target {
			high = mid - 1
		} else  {
			low = mid + 1
		}
	}

	if count > 0 {
		for i := mid + 1; i <= len(nums)-1; i++ {
			if nums[i] == target {
				count++
			} else {
				break
			}
		}

		for j := mid - 1; j >= 0; j-- {
			if nums[j] == target {
				count++
			} else {
				break
			}
		}
	}
	return count
}

func main() {
	data := []int{5, 7, 7,  8, 8, 10}
	num := search(data, 6)
	fmt.Println("出现次数:", num)
}
