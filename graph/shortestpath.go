package graph

import "math"

/** 图的最短路径算法 **/

// Dijkstra 图的最短路径算法的一种实现
// 给定一个图的某个节点，返回这个节点到图的其他点的最短距离
// 某个点不在map中记录，则from点到该点位正无穷（不可达）
// 算法思路：
//
//	Dijkstra算法用于求解**非负权有向图**（或仅含正权有向图）的**单源**最短路径，算法基于广度优先搜索算法 bfs，利用贪心策略实现。
//	它的思路是，从起点开始，不断扩展距离最小的顶点，依次得到所有顶点的最短路径。
func (g *Graph[T]) Dijkstra(from *Node[T]) map[*Node[T]]int {
	// 从from出发到所有点的最小距离表（DP表）
	distanceMap := make(map[*Node[T]]int, 0)
	// from到from距离为0
	distanceMap[from] = 0
	// 已经求过距离的节点，存在selectedNodes中，不会再被选中记录
	selectedNodesSet := make(map[*Node[T]]string)
	// 得到没选择过的点的最小距离。该函数为方法，目的是为了把from的类型传递过去。
	// 贪心找到最近的点，再通过该点（桥接点）荡开去，继续贪心。这里此时那倒的桥节点就是刚初始化塞进去的from
	minNode := from.getMinDistanceAndUnselectedNode(distanceMap, selectedNodesSet)

	// 得到minNode之后
	for minNode != nil {
		// 把minNode对应的距离取出,此时minNode就是桥连点
		distance := distanceMap[minNode]
		// 把minNode上所有的邻边拿出来
		// 这里就是要拿到例如A到C和A到桥连点B再到C哪个距离小的距离
		for _, edge := range minNode.edges {
			// 某条边对应的下一跳节点toNode
			toNode := edge.to
			// 如果关于from的distencMap中没有去toNode的记录，表示正无穷，直接添加该条
			if _, ok := distanceMap[toNode]; !ok {
				// from到minNode的距离加上个minNode到当前to节点的边距离
				distanceMap[toNode] = distance + edge.weight
			} else { // 如果有，看该距离是否更小，更小就更新（贪心到一条更优的路径）
				minDistance := int(math.Min(float64(distanceMap[toNode]), float64(distance+edge.weight)))
				distanceMap[toNode] = minDistance
			}
		}
		// 锁上minNode，表示from通过minNode到其他节点的最小值已经找到并且维护到了dp表
		// minNode将不再使用
		selectedNodesSet[minNode] = ""
		// 再在没有选择的节点中挑选MinNode当成from的桥接点，直到所有点都选择过，会退出查找（dp表已经全部维护）。
		minNode = from.getMinDistanceAndUnselectedNode(distanceMap, selectedNodesSet)
	}
	// 最终distanceMap全部更新，dp表返回
	return distanceMap
}

// getMinDistanceAndUnselectedNode 得到没选择过的点的最小距离
func (from *Node[T]) getMinDistanceAndUnselectedNode(
	distanceMap map[*Node[T]]int, selectedNodesSet map[*Node[T]]string) *Node[T] {
	var minNode *Node[T] = nil
	minDistance := math.MaxInt
	for node, distance := range distanceMap {
		// 没有被选择过，且距离最小
		if _, ok := selectedNodesSet[node]; !ok &&
			distance < minDistance {
			minNode = node
			minDistance = distance
		}
	}
	return minNode
}

// Floyd 算法用来求图的最短路径，可以处理权值为负的场景。其基本思想是利用中间点的集合逐步逼近最终解，不断更新每两点之间的距离。
//
// 算法思路:
//  1. 对于图中任意两个顶点，如果它们之间有一条边，则这两个顶点之间的距离为这条边的权值，否则它们之间的距离为无穷大。
//  2. 对于每个中间顶点 k，将顶点按照 k 的增序进行遍历，即先访问 k = 0，再访问 k = 1，以此类推。
//  3. 对于每一对顶点 i 和 j，若通过 k 顶点可以使 i 到 j 的距离缩短，则更新距离。（处理边的权值为负的问题）
//  4. 最终，得到的距离矩阵即为所有顶点之间的最短距离。
//
// Floyd 算法的**时间复杂度为O(N^3)**，其中 N 表示图中的节点数。Floyd 算法可以处理带有负权边的图，在一定程度上弥补了 Dijkstra 算法的不足。
//  1. 可以兼容边权值为负的情况
//  2. 可以处理多源问题
//
// 但是，Floyd **算法的空间复杂度为O(N^2)**，对于节点数较大的图可能会占用过多的内存。
func (g *Graph[T]) Floyd() map[*Node[T]]map[*Node[T]]int {
	// 初始化距离矩阵
	distanceMap := make(map[*Node[T]]map[*Node[T]]int)
	for from := range g.nodes {
		// 为每个点都生成一个距离map，键为目标节点，值为源点到目标点的距离
		distanceMap[g.nodes[from]] = make(map[*Node[T]]int)
		for to := range g.nodes {
			if from == to {
				distanceMap[g.nodes[from]][g.nodes[to]] = 0
				continue
			}
			// 默认距离为正无穷
			distanceMap[g.nodes[from]][g.nodes[to]] = math.MaxInt
		}
	}

	// 根据实际边权初始化距离矩阵
	for e := range g.edges {
		from := e.from
		to := e.to
		distanceMap[from][to] = e.weight
	}

	// Floyd 算法核心代码
	for k := range g.nodes {
		for i := range g.nodes {
			for j := range g.nodes {
				// 如果经过中间点k，从 i 到达 j 有更短路径，则替换原来的距离i到j的距离
				if distanceMap[g.nodes[i]][g.nodes[k]] != math.MaxInt &&
					distanceMap[g.nodes[k]][g.nodes[j]] != math.MaxInt &&
					distanceMap[g.nodes[i]][g.nodes[k]]+distanceMap[g.nodes[k]][g.nodes[j]] < distanceMap[g.nodes[i]][g.nodes[j]] {
					distanceMap[g.nodes[i]][g.nodes[j]] = distanceMap[g.nodes[i]][g.nodes[k]] + distanceMap[g.nodes[k]][g.nodes[j]]
				}
			}
		}
	}

	return distanceMap
}
