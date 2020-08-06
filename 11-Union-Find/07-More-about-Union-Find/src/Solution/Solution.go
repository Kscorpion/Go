package main

import (
	"Play-with-Data-Structures/11-Union-Find/02-Quick-Find/src/LinkedListSet"
	"fmt"
)

/// Leetcode 547. Friend Circles
/// https://leetcode.com/problems/friend-circles/description/
///
/// 课程中在这里暂时没有介绍这个问题
/// 该代码主要用于使用Leetcode上的问题测试我们的UF类
// 我们的第一版Union-Find
type UnionFind6 struct {
	parent []int // parent[i]表示第一个元素所指向的父节点
	rank   []int // rank[i]表示以i为根的集合中元素个数
}

// 构造函数
func Constructor(size int) *UnionFind6 {
	rank := make([]int, size)
	parent := make([]int, size)
	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i := 0; i < len(parent); i++ {
		rank[i] = 1
		parent[i] = i
	}

	return &UnionFind6{
		rank:   rank,
		parent: parent,
	}
}

func (this *UnionFind6) GetSize() int {
	return len(this.parent)
}

// 查找过程, 查找元素p所对应的集合编号
// O(h)复杂度, h为树的高度
func (this *UnionFind6) find(p int) int {
	if p < 0 || p > len(this.parent) {
		panic("p is out of range.")
	}

	if p != this.parent[p] {
		this.parent[p] = this.find(this.parent[p])
	}

	return this.parent[p]
}

// 查看元素p和元素q是否所属一个集合
// O(h)复杂度, h为树的高度
func (this *UnionFind6) IsConnected(p int, q int) bool {
	return this.find(p) == this.find(q)
}

// 合并元素p和元素q所属的集合
// O(h)复杂度, h为树的高度
func (this *UnionFind6) UnionElements(p int, q int) {
	pRoot := this.find(p)
	qRoot := this.find(q)

	if pRoot == qRoot {
		return
	}

	// 根据两个元素所在树的rank不同判断合并方向
	// 将rank低的集合合并到rank高的集合上
	if this.rank[pRoot] < this.rank[qRoot] {
		this.parent[pRoot] = this.parent[qRoot]
	} else if this.rank[pRoot] > this.rank[qRoot] {
		this.parent[qRoot] = this.parent[pRoot]
	} else { // rank[pRoot] == rank[qRoot]
		this.parent[pRoot] = qRoot
		this.rank[qRoot] += 1 // 此时, 我维护rank的值
	}
}

func findCircleNum(M [][]int) int {
	length := len(M)
	uf := Constructor(length)

	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				uf.UnionElements(i, j)
			}
		}
	}

	linkedListSet := LinkedListSet.Constructor()
	for i := 0; i < length; i++ {
		linkedListSet.Add(uf.find(i))
	}

	return linkedListSet.GetSize()
}

func main() {
	M := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}

	fmt.Println(findCircleNum(M))
}
