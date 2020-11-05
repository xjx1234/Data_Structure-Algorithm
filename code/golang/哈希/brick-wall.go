/**
* @Author: XJX
* @Description:
	LeetCode:554
	你的面前有一堵矩形的、由多行砖块组成的砖墙。 这些砖块高度相同但是宽度不同。你现在要画一条自顶向下的、穿过最少砖块的垂线。
	砖墙由行的列表表示。 每一行都是一个代表从左至右每块砖的宽度的整数列表。
	如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。你需要找出怎样画才能使这条线穿过的砖块数量最少，并且返回穿过的砖块数量。
	你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。
* @File: brick-wall
* @Date: 2020/11/4 17:22
*/

package main

import "fmt"

func leastBricks(wall [][]int) int {
	brickNum := make(map[int]int)
	maxLevel := 0
	for i := 0; i < len(wall); i++ {
		sum := 0
		for j := 0; j < len(wall[i])-1; j++ {
			fmt.Println("i:", i, "j:", j, "v:", wall[i][j])
			sum += wall[i][j]
			brickNum[sum] += 1
			fmt.Println(brickNum)
			if brickNum[sum] > maxLevel {
				maxLevel = brickNum[sum]
			}
			fmt.Println(maxLevel)
		}
	}
	return len(wall) - maxLevel
}

func main() {
	data := [][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}
	fmt.Println(leastBricks(data))
}
