package BST

import (
	"Play-with-Data-Structures/Utils/Interfaces"
	"bytes"
	"fmt"
)

type Node struct {
	e     interface{}
	left  *Node
	right *Node
}

func generateNode(e interface{}) *Node {
	return &Node{e: e}
}

type BST struct {
	root *Node
	size int
}

func Constructor() *BST {
	return &BST{}
}

func (this *BST) GetSize() int {
	return this.size
}

func (this *BST) IsEmpty() bool {
	return this.size == 0
}

// 向二分搜索树中添加新的元素 e
func (this *BST) Add(e interface{}) {
	this.root = this.add(this.root, e)
}

// 向以 Node 为跟的二分搜索树中插入元素 e，递归算法
// 返回插入新节点后二分搜索树的根
func (this *BST) add(n *Node, e interface{}) *Node {
	if n == nil {
		this.size++
		return generateNode(e)
	}

	// 递归调用
	if Interfaces.Compare(e, n.e) < 0 {
		n.left = this.add(n.left, e)
	} else if Interfaces.Compare(e, n.e) > 0 {
		n.right = this.add(n.right, e)
	}
	return n
}

// 看二分搜索树中是否包含元素 e
func (this *BST) Contains(e interface{}) bool {
	return contains(this.root, e)
}

// 看以 Node 为根的二分搜索树是否包含元素 e，递归算法
func contains(n *Node, e interface{}) bool {
	if n == nil {
		return false
	}

	if Interfaces.Compare(e, n.e) == 0 {
		return true
	} else if Interfaces.Compare(e, n.e) < 0 {
		return contains(n.left, e)
	} else {
		return contains(n.right, e)
	}
}

// 二分搜索树的前序遍历
func (this *BST) PreOrder() {
	preOrder(this.root)
}

// 前序遍历以 Node 为根的二分搜索树，递归算法
func preOrder(n *Node) {
	if n == nil {
		return
	}

	fmt.Println(n.e)
	preOrder(n.left)
	preOrder(n.right)
}

// 二分搜索树的非递归前序遍历
//func (this *BST) PreOrderNR() {
//	// 使用之前我们自己实现的数组栈
//	stack := ArrayStack.Constructor(20)
//	stack.Push(this.root)
//
//	for !stack.IsEmpty() {
//		cur := stack.Pop().(*Node)
//		fmt.Println(cur.e)
//
//		if cur.right != nil {
//			stack.Push(cur.right)
//		}
//		if cur.left != nil {
//			stack.Push(cur.left)
//		}
//	}
//}

// 二分搜索树的中序遍历
func (this *BST) InOrder() {
	inOrder(this.root)
}

// 中序遍历以 Node 为根的二分搜索树，递归算法
func inOrder(n *Node) {
	if n == nil {
		return
	}

	inOrder(n.left)
	fmt.Println(n.e)
	inOrder(n.right)
}

// 二分搜索树的后序遍历
func (this *BST) PostOrder() {
	postOrder(this.root)
}

// 后序遍历以 Node 为根的二分搜索树，递归算法
func postOrder(n *Node) {
	if n == nil {
		return
	}

	postOrder(n.left)
	postOrder(n.right)
	fmt.Println(n.e)
}

// 二分搜索树的层序遍历
//func (this *BST) LevelOrder() {
//	// 使用我们之前实现的循环队列
//	queue := LoopQueue.Constructor(20)
//	queue.Enqueue(this.root)
//	for !queue.IsEmpty() {
//		cur := queue.Dequeue().(*Node)
//		fmt.Println(cur.e)
//
//		if cur.left != nil {
//			queue.Enqueue(cur.left)
//		}
//		if cur.right != nil {
//			queue.Enqueue(cur.right)
//		}
//	}
//}

// 寻找二分搜索树的最小元素
func (this *BST) Minimum() interface{} {
	if this.size == 0 {
		panic("BST is empty!")
	}
	return minimum(this.root).e
}

// 返回以 Node 为根的二分搜索树的最小值所在的节点
func minimum(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return minimum(n.left)
}

// 寻找二分搜索树的最大元素
func (this *BST) Maximum() interface{} {
	if this.size == 0 {
		panic("BST is empty!")
	}
	return maximum(this.root).e
}

// 返回以 Node 为根的二分搜索树的最大值所在的节点
func maximum(n *Node) *Node {
	if n.right == nil {
		return n
	}
	return maximum(n.right)
}

// 从二分搜索树中删除最小值所在的节点，返回最小值
func (this *BST) RemoveMin() interface{} {
	// 获得最小值
	ret := this.Minimum()
	this.root = this.removeMin(this.root)
	return ret
}

// 删除以 Node 为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *BST) removeMin(n *Node) *Node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		this.size--
		return rightNode
	}
	n.left = this.removeMin(n.left)
	return n
}

// 从二分搜索树中删除最小值所在的节点，返回最小值
func (this *BST) RemoveMax() interface{} {
	// 获得最小值
	ret := this.Maximum()
	this.root = this.removeMax(this.root)
	return ret
}

// 删除以 Node 为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *BST) removeMax(n *Node) *Node {
	if n.right == nil {
		leftNode := n.left
		n.left = nil
		this.size--
		return leftNode
	}
	n.right = this.removeMax(n.right)
	return n
}

func (this *BST) String() string {
	var buffer bytes.Buffer
	generateBSTSting(this.root, 0, &buffer)
	return buffer.String()
}

// 生成以 Node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(n *Node, depth int, buffer *bytes.Buffer) {
	if n == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(n.e) + "\n")
	generateBSTSting(n.left, depth+1, buffer)
	generateBSTSting(n.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
