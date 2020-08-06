package UnionFind6

// rank[i]表示以i为根的集合所表示的树的层数
// 在后续的代码中, 我们并不会维护rank的语意, 也就是rank的值在路径压缩的过程中, 有可能不在是树的层数值
// 这也是我们的rank不叫height或者depth的原因, 他只是作为比较的一个标准
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
