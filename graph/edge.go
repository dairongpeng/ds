package graph

// Edge 图中的边元素表示
type Edge[T int | int64 | float64 | string | *interface{}] struct {
	// 边的权重信息
	weight int
	// 出发的节点
	from *Node[T]
	// 指向的节点
	to *Node[T]
}
