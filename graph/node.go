package graph

// Node 图中的点元素表示
type Node[T int | int64 | float64 | string | *interface{}] struct {
	// 点的身份标识
	value T
	// 入度，表示有多少个点连向该点
	in int
	// 出度，表示从该点出发连向别的节点多少
	out int
	// 直接邻居：表示由自己出发，直接指向哪些节点。指向节点的总数等于out
	nexts []*Node[T]
	// 直接下级边：表示由自己出发的边有多少
	edges []*Edge[T]
}
