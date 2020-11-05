/**
* @Author: XJX
* @Description:
	LeetCode:500
	给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词
* @File: keyboard-row
* @Date: 2020/11/4 17:05
*/

package main

func findWords(words []string) []string {
	var result []string
	dist := map[byte]byte{
		'q': 1, 'Q': 1,
		'w': 1, 'W': 1,
		'e': 1, 'E': 1,
		'r': 1, 'R': 1,
		't': 1, 'T': 1,
		'y': 1, 'Y': 1,
		'u': 1, 'U': 1,
		'i': 1, 'I': 1,
		'o': 1, 'O': 1,
		'p': 1, 'P': 1,
		'a': 2, 'A': 2,
		's': 2, 'S': 2,
		'd': 2, 'D': 2,
		'f': 2, 'F': 2,
		'g': 2, 'G': 2,
		'h': 2, 'H': 2,
		'j': 2, 'J': 2,
		'k': 2, 'K': 2,
		'l': 2, 'L': 2,
		'z': 3, 'Z': 3,
		'x': 3, 'X': 3,
		'c': 3, 'C': 3,
		'v': 3, 'V': 3,
		'b': 3, 'B': 3,
		'n': 3, 'N': 3,
		'm': 3, 'M': 3,
	}
	for _, v := range words {
		flag := true
		for i := len(v) - 1; i > 0; i-- {
			if dist[v[i]] != dist[v[i-1]] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, v)
		}
	}
	return result
}

func main() {

}
