package main

import "fmt"

func BinarySearch(arr []int, tag int) int {
	length := len(arr)
	l := 0
	r := length - 1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == tag {
			return mid
		} else if tag < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(list, 7))
}
