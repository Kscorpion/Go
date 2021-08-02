package main

import "fmt"

type L1 struct {
	data int
	next *L1
}

//头插法
func createListHead(n int) *L1 {
	temp := &L1{}
	temp = nil
	for i := n; i >= 0; i-- {
		p := &L1{}
		p.data = i
		p.next = temp
		temp = p
	}
	return temp
}

//尾插法
func createListTail(n int) *L1 {
	H := &L1{}
	temp := H
	for i := n; i >= 0; i-- {
		p := &L1{}
		p.data = i
		temp.next = p
		temp = p
	}
	temp.next = nil
	return H.next
}

func LinkRec(p *L1) (newHead *L1) {
	newNode := &L1{}
	for p != nil {
		newNode = p
		p = p.next
		newNode.next = newHead
		newHead = newNode
	}
	return
}
//合并有序链表（仅限从小到大有序）
func mergeTwoLists(l1 *L1, l2 *L1) *L1 {
	newHead := &L1{}
	n := newHead
	//等长度比较完成
	for l1 != nil && l2 != nil {
		//如需从大到小把下面的' < '改为' > '
		if l1.data < l2.data {
			n.next = l1
			l1 = l1.next
		} else {
			n.next = l2
			l2 = l2.next
		}
		n = n.next
	}
	if l1 != nil {
		n.next = l1
	}
	if l2 != nil {
		n.next = l2
	}
	return newHead.next
}

func displayL1(l1 *L1) {
	for l1 != nil {
		fmt.Println(l1.data)
		l1 = l1.next
	}
}
func main() {
	fmt.Println("头插结果")
	displayL1(createListHead(10))
	fmt.Println("尾插结果")
	displayL1(createListTail(10))
	fmt.Println("反转结果")
	displayL1(LinkRec(createListTail(10)))
}
