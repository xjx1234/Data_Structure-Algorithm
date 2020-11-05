/**
* @Author: XJX
* @Description:
	LeetCode:1160
	给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。
	假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。
	注意：每次拼写（指拼写词汇表中的一个单词）时，chars 中的每个字母都只能用一次。
	返回词汇表 words 中你掌握的所有单词的 长度之和。
* @File: find-words.go
* @Date: 2020/10/26 16:54
*/

package main

import (
	"fmt"
	"strings"
)

/**
根据hash统计算法
*/
func countCharacters(words []string, chars string) int {
	count := 0
	stat := make([]int, 128)
	newStat := make([]int, 128)
	for i := range chars {
		stat[chars[i]]++
	}
	for _, v := range words {
		copy(newStat, stat)
		isStat := true
		for _, v1 := range v {
			if newStat[byte(v1)] > 0 {
				newStat[byte(v1)]--
			} else {
				isStat = false
				break
			}
		}
		if isStat {
			count += len(v)
			fmt.Println(v)
		}
	}
	return count
}

/**
穷举办法解题
*/
func countCharacters1(words []string, chars string) int {
	count := 0
here:
	for _, v := range words {
		for _, v1 := range v {
			if strings.Count(string(v), string(v1)) > strings.Count(string(chars), string(v1)) {
				continue here
			}
		}
		count += len(v)
	}
	return count
}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	count := countCharacters1(words, chars)
	count1 := countCharacters(words, chars)
	fmt.Println("len", count, count1)
}
