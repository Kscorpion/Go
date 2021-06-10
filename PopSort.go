package main

import "fmt"

func PopSort(list []int) []int {
	var (
		i, j, temp, issort int
	)
	length := len(list)
	for i = 0; i < length-1; i++ {
		issort = 0
		for j = 0; j < length-i-1; j++ {
			if list[j] > list[j+1] {
				issort = 1
				temp = list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
		if issort == 0 {
			break
		}
	}
	return list
}

func main() {
	s := []int{3, 5, 4, 9, 7, 2, 8, 1}
	fmt.Println(PopSort(s))
}
