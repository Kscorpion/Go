package main

import (
	"fmt"
	"math"
)

type win struct {
	winlen int
}

func (win *win) max(a []int, count int) int {
	var (
		max    int
		newmax int
	)
	if win.winlen > count {
		return -1
	}
	for x, y := range a {
		if x > win.winlen {
			break
		}
		max += y
	}
	fmt.Println(max)
	for i := win.winlen; i < count; i++ {
		newmax += a[i] - a[i-win.winlen]
		max = int(math.Max(float64(newmax), float64(max)))
	}
	return max
}

func main() {
	w := &win{
		winlen: 4,
	}

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 30, 44, 54, 32, 567, 1, 2, 3}
	fmt.Println("定制长度:", w.winlen)
	max := w.max(arr, len(arr))
	fmt.Println(max)
}
