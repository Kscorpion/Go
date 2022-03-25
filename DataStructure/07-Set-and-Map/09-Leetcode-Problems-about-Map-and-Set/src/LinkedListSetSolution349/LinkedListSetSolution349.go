package main

import "fmt"

type Node struct {
	e    interface{}
	next *Node
}

type LinkedList struct {
	dummyHead *Node // 虚拟头结点，不计入size
	size      int
}

func GetLinkedList() *LinkedList {
	return &LinkedList{
		dummyHead: &Node{},
	}
}

// 获取链表中的元素个数
func (l *LinkedList) GetSize() int {
	return l.size
}

// 返回链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("Add failed. Illegal index.")
	}

	// 获得待插入节点的前一个节点
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = &Node{e, prev.next}
	l.size++
}

// 在链表头添加新的元素e
func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

// 在链表末尾添加新的元素e
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

// 获得链表的第index(0-based)个位置的元素
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("Add failed. Illegal index.")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

// 获得链表的第一个元素
func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

// 获得链表的最后一个元素
func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

// 修改链表的第index(0-based)个位置的元素为e
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= l.size {
		panic("Set failed. Illegal index.")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

// 查找链表是否存在元素e
func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next

	for cur != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}

	return false
}

// 从链表中删除index(0-based)位置的元素，返回删除的元素
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("Remove failed. Index is illegal.")
	}

	// prev 是待删除元素的前一个元素
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil
	l.size--

	return retNode.e
}

// 从链表中删除第一个元素，返回删除的元素
func (l *LinkedList) RemoveFirst() {
	l.Remove(0)
}

// 从链表中删除最后一个元素，返回删除的元素
func (l *LinkedList) RemoveLast() {
	l.Remove(l.size - 1)
}

func (l *LinkedList) RemoveElement(e interface{}) {
	prev := l.dummyHead

	for prev.next != nil {
		if prev.next.e == e {
			break
		}
		prev = prev.next
	}

	if prev.next != nil {
		delNode := prev.next
		prev.next = delNode.next
		delNode.next = nil
		l.size--
	}
}

/// Leetcode 349. Intersection of Two Arrays
/// https://leetcode.com/problems/intersection-of-two-arrays/description/
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	linkedListSet := GetLinkedList()

	for _, num := range nums1 {
		linkedListSet.AddFirst(num)
	}

	for _, num := range nums2 {
		if linkedListSet.Contains(num) {
			res = append(res, num)
			linkedListSet.RemoveElement(num)
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(nums1, nums2))
}
