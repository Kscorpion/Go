package main

import "fmt"

type List struct {
	data int
	next *List
}

var (
	link *List
	p *List
	node *List
	newHead *List
	newNode *List
	recv *List
	n int
)

func InitNode ()(link *List){
	link = &List{}
	return
}

func BuildLink()(node *List){
	link = InitNode()
	for n=1;n<10;n++{
		node = InitNode()
		node.data = n
		node.next = link
		link = node
	}
	return
}

func LinkRecv(p *List)(newHead *List){
	newHead = nil
	newNode = InitNode()
	for  {
		if p!=nil {
			newNode = p
			p = p.next
			newNode.next = newHead
			newHead = newNode
		}else {
			break
		}
	}
	return
}

func PrintLink(node *List){
	for{
		if node!=nil {
			fmt.Printf("%d",node.data)
			node = node.next
		}else {
			break
		}
	}
	println()
}

func main(){
	p = BuildLink()
	PrintLink(p)
	recv = LinkRecv(p)
	PrintLink(recv)
}
