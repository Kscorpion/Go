package main

import "fmt"

type heap struct{}

var (
	n int
	c1 int
	c2 int
	max int
	temp int
	last_node int
	parent int
	h *heap
	tree []int
)

func(heap *heap)swap(arr[]int,swap_a int,swap_b int){
	temp = arr[swap_a]
	arr[swap_a] = arr[swap_b]
	arr[swap_b] = temp
	return
}

func (heap *heap)heapify(arr[]int,count int,i int){
	if i>=count {
		return
	}
	c1 = 2 * i +1
	c2 = 2 * i +2
	max = i
	if(c1<count && arr[c1]>arr[max]){
		max = c1
	}
	if(c2<count && arr[c2]>arr[max]){
		max = c2
	}
	if( i != max ){
		heap.swap(arr,max,i)
		heap.heapify(arr,count,max)
	}
	return
}

func (heap *heap)builheap(arr[]int,count int){
	last_node = count-1
	parent = (last_node-1)/2
	for n=parent;n>=0;n-- {
		heap.heapify(arr,count,n)
	}
	return
}

func (heap *heap)heapsort(arr[]int,count int){
	heap.builheap(arr,count)
	for n= count-1;n>=0;n--{
		heap.swap(arr,n,0)
		heap.heapify(arr,n,0)
	}
	return
}

func main(){
	h = &heap{}
	tree =[]int{153,43,123,54,213,5,29993,65,21}
	h.heapsort(tree,len(tree))
	for _,val :=range tree{
		fmt.Println(val)
	}
	return
}
