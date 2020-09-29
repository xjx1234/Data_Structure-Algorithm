/**
* @Author: XJX
* @Description: 简单的二分查找实现
* @题目: 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1
* @File: binary-search.go
* @Date: 2020/9/28 11:31
 */

package main

import "fmt"

/**
非递归解决方案
1. high = mid - 1 与 low = mid + 1 是为了防止产生死循环，当low与high一致会产生死循环
2. mid = (low + high) >> 1 这样性能会更好，也有教程推荐  mid =  low + ((high - low) >> 1) 方案
*/
func serach(nums []int, target int) int {
	len := len(nums)
	if len >= 1 {
		mid := 0
		low := 0
		high := len - 1
		for low <= high {
			mid = (low + high) >> 1
			fmt.Println(low, high, mid)
			if nums[mid] == target {
				return mid
			} else if nums[mid] > target {
				high = mid - 1
			} else if nums[mid] < target {
				low = mid + 1
			}
		}
	}
	return -1
}

//递归解决方案
func serachByRecursion(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	return bsearch(nums, low, high, target)
}

func bsearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) >> 1
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return bsearch(nums, low, mid-1, target)
	} else if nums[mid] < target {
		return bsearch(nums, mid+1, high, target)
	}
	return -1
}

func main() {
	data := []int{}
	for i := 0; i < 51; i++ {
		data = append(data, i)
	}

	key := serach(data, 5)
	fmt.Println("查询结果:", key)
	//key2 := serach(data, 9)
	//fmt.Println("查询结果:", key2)
	//key3 := serachByRecursion(data, 5)
	//fmt.Println("查询结果:", key3)
}
