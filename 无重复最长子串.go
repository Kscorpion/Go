package main

import "fmt"

func main() {
	s := "abba"
	fmt.Println(lengthOfLongestSubString(s))
}

func lengthOfLongestSubString(s string) int {
	length := len(s)
	ans := 0
	cmap := make([]int, 128)
	left := 0
	for right := 0; right < length; right++ {
		index := s[right]
		//若有重复字符 左指针移动到该位置
		//两种情况 1存在过 【字符下标】 2没存在过 【0】
		left = max(left, cmap[index])
                //始终保存最大长度
		ans = max(ans, right-left+1)
                //
		cmap[index] = right + 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
