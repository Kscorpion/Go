package main

import "fmt"

func main() {
	//递归
	fmt.Println(fibonacci1(10))
	//递归
	fmt.Println(fibonacci2(10))
	//双指针
	fmt.Println(fibonacci3(10))
}

//斐波那契数列-普通地柜
//时间复杂度O(2^n)
func fibonacci1(num int) int {
	if num == 1 {
		return 0
	}
	if num == 2 {
		return 1
	}
	return fibonacci1(num-1) + fibonacci1(num-2)
}

//斐波那契数列-递归优化
//时间复杂度 O(N) 空间复杂度 O(N)
func fibonacci2(num int) int {
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
	list[num] = fibonacci2(num-1) + fibonacci2(num-2)
	return list[num]
}

//斐波那契数列-双指针
//时间复杂度O(N)
func fibonacci3(num int) int {
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
