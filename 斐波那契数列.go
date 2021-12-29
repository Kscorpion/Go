package main

import "fmt"

func main() {
	//递归
	fmt.Println(fbnq1(10))
	//递归
	fmt.Println(fbnq2(10))
	//双指针
	fmt.Println(fbnq2(10))
}

//斐波那契数列-普通地柜
//时间复杂度O(2^n)
func fbnq1(num int) int {
	if num == 1 {
		return 0
	}
	if num == 2 {
		return 1
	}
	return fbnq1(num-1) + fbnq1(num-2)
}

//斐波那契数列-递归优化
//时间复杂度 O(N) 空间复杂度 O(N)
func fbnq2(num int) int {
	var (
		list = make([]int, num+1)
	)
	if num == 1 {
		return 0
	}
	if num == 2 {
		return 1
	}
	if list[num] != 0 {
		return list[num]
	}
	list[num] = fbnq2(num-1) + fbnq2(num-2)
	return list[num]
}

//斐波那契数列-双指针
//时间复杂度O(N)
func fbnq3(num int) int {
	if num == 1 {
		return 0
	}
	if num == 2 {
		return 1
	}
	low := 0
	high := 1
	for i := 2; i < num; i++ {
		sum := low + high
		low = high
		high = sum
	}
	return high
}
