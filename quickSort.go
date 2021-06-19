package main

import "fmt"

func quickSort(arr []int, left, right int) []int {
	if left > right { //很重要
		return arr
	}
	i := left
	j := right
	base := arr[left]
	for j > i {
		for arr[j] >= base && j > i {
			j--
		}
		for arr[i] <= base && j > i {
			i++
		}
		if j > i {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left] = arr[i]
	arr[i] = base
	quickSort(arr, 0, i-1)
	quickSort(arr, i+1, right)
	return arr
}
func main() {
	a := []int{12, 34, 67, 2, 89, 14}
	fmt.Println(quickSort(a, 0, len(a)-1))
}
