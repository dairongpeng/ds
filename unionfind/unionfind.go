package unionfind

import "github.com/dairongpeng/ds/stack/arraystack"

// UnionFind 并查集结构
type UnionFind[T int | int64 | float64 | string | *interface{}] struct {
	// 并查集中的点和该点的代表节点的映射
	flag map[T]T
	// 当前点，是代表点，会在sizeMap中记录该代表点的连通个数
	size map[T]int
}

// NewUnionFind 构建一个并查集结构
func NewUnionFind[T int | int64 | float64 | string | *interface{}](values []T) *UnionFind[T] {
	f := make(map[T]T, 0)
	s := make(map[T]int, 0)

	// 一开始时：一个v的代表点是自身； 每个v都是代表点，且代表点只代表自己size为1
	for _, v := range values {
		f[v] = v
		s[v] = 1
	}

	return &UnionFind[T]{
		flag: f,
		size: s,
	}
}

// FindFlag 在并查集结构中找一个节点的代表节点
// 从点cur开始，一直往上找，找到不能再往上的代表点，返回
// 通过把路径上所有节点指向最上方的代表节点，目的是把findFlag优化成O(1)的
func (set *UnionFind[T]) FindFlag(cur T) T {
	// 在找flag的过程中，沿途所有节点加入当前容器，便于后面扁平化处理
	path := arraystack.New[T]()
	// 当前节点的父亲不是指向自己，进行循环
	for cur != set.flag[cur] {
		// 在向某个点代表节点查找的过程中，沿途所有点入栈
		path.Push(cur)
		// 向上移动
		cur = set.flag[cur]
	}
	// 循环结束，cur此时是最上的代表节点。且沿途所有点已经在栈中
	// 把沿途所有节点拍平，都指向当前最上方的代表节点
	// a -> b -> c -> d 拍平后=> a -> d; b -> d; c -> d
	for !path.IsEmpty() {
		pathV, _ := path.Pop()
		set.flag[pathV] = cur
	}
	return cur
}

// Find 判断两个元素是否在同一个并查集中
func (set *UnionFind[T]) Find(a, b T) bool {
	// 比较a的最上的代表点和b最上的代表点, 如果相同，ab在一个集合中
	return set.FindFlag(a) == set.FindFlag(b)
}

func (set *UnionFind[T]) Union(a, b T) {
	flagA := set.FindFlag(a)
	flagB := set.FindFlag(b)

	if flagA != flagB {
		flagASize := set.size[flagA]
		flagBSize := set.size[flagB]
		if flagASize >= flagBSize { // a所在的集合比b所在的集合大， b集合往a集合合并
			set.flag[flagB] = flagA
			set.size[flagA] = flagASize + flagBSize
			// flagB已经不是代表点了
			delete(set.size, flagB)
		} else { // b所在集合比a所在集合大， a集合往b集合合并
			set.flag[flagA] = flagB
			set.size[flagB] = flagASize + flagBSize
			// flagA已经不是代表点了
			delete(set.size, flagA)
		}
	}
}
