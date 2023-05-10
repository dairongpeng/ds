package graph

type Graph[T int | int64 | float64 | string | *interface{}] struct {
	// 点的集合，编号为1的点是什么，用map
	nodes map[T]*Node[T]
	// 边的集合(用hash实现set)
	edges map[*Edge[T]]string
}

// NewGraph 初始化图结构
func NewGraph[T int | int64 | float64 | string | *interface{}]() *Graph[T] {
	return &Graph[T]{
		nodes: make(map[T]*Node[T], 0),
		edges: make(map[*Edge[T]]string, 0),
	}
}

// GetNodes 获取图的点集合
func (g *Graph[T]) GetNodes() map[T]*Node[T] {
	return g.nodes
}

// GetEdges 获取图的边集合
func (g *Graph[T]) GetEdges() map[*Edge[T]]string {
	return g.edges
}

// AddNode 往图中加入一个点
func (g *Graph[T]) AddNode(v T) *Node[T] {
	node := &Node[T]{
		value: v,
		in:    0,
		out:   0,
		nexts: nil,
		edges: nil,
	}
	g.nodes[v] = node

	return node
}

// AddEdge 往图中加入一个边
func (g *Graph[T]) AddEdge(from, to *Node[T], weight int) {
	edge := &Edge[T]{
		weight: weight,
		from:   from,
		to:     to,
	}
	// 两个点的入度出度维护
	from.out++
	to.in++

	// from的邻接点维护
	from.nexts = append(from.nexts, to)
	// from的直接下级边维护
	from.edges = append(from.edges, edge)

	g.edges[edge] = ""
}
