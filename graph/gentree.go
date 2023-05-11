package graph

/** 图的最小生成树算法 **/

import (
	"github.com/dairongpeng/ds/heap/minheap"
	"github.com/dairongpeng/ds/pkg"
	"github.com/dairongpeng/ds/unionfind"
)

// KruskalMST 克鲁斯卡尔最小生成树算法。返回set。
// 在不破坏原有图点与点的连通性基础上，让连通的边的整体权值最小。返回最小权值或者边的集合
// comparator: 两个边如何比较大小
// 1. 将所有边按照权值从小到大排序。（小根堆）
// 2. 初始化一个空集合，作为生成树，保存生成树的各个边。（hash set）
// 3. 依次遍历原始图的所有边，若发现该边所连接的两个顶点在当前集合中不连通，则将该边加入生成树中。（并查集）
// 4. 重复步骤3，直至生成整个图的最小生成树为止。
// A -> B 且 B -> A只会在并查集中保留一个，k算法只验证了连通性，未区分联通结构。
func (g *Graph[T]) KruskalMST(comparator pkg.Comparator[*Edge[T]]) map[*Edge[T]]string {
	values := make([]T, 0)
	for k := range g.nodes {
		values = append(values, k)
	}
	// 初始化一个并查集结构
	unionFindSet := unionfind.NewUnionFind[T](values)

	// 初始化一个小根堆
	edgesMinHeap := minheap.NewMaxHeap[*Edge[T]](len(g.edges), comparator)
	// 边按照权值从小到大排序，加入到堆
	for edge := range g.edges {
		_ = edgesMinHeap.Push(edge) // limit等于len(edges)，不会越界报错
	}

	resultSet := make(map[*Edge[T]]string)

	// 堆不为空，弹出小根堆的堆顶
	for !edgesMinHeap.IsEmpty() {
		// 假设M条边，O(logM), 选择小根堆的最小的边，进行贪心。
		edge := edgesMinHeap.Pop()
		// 如果该边的左右两侧不在同一个集合中， 否则已经贪心到这两个点更小的边了，不需要再收集
		if !unionFindSet.Find(edge.from.value, edge.to.value) {
			// 需要收集这条边
			resultSet[edge] = ""
			// 联合from和to， 加入到并查集中去
			unionFindSet.Union(edge.from.value, edge.to.value)
		}
	}
	return resultSet
}

// PrimMST prim算法实现图的最小生成树。
// 1. 随机选定一个起点，将其标记为已访问，将与之相邻的边（权值）加入到堆或者优先队列中。
// 2. 从堆或者优先队列中选出代价最小的边，若该边所连接的顶点未被访问过，则将该顶点标记为已访问，并将与该顶点相连的边（权值）加入堆或队列中。
// 3. 重复步骤2，直至所有顶点都被访问，此时形成的边就是最小生成树。
func (g *Graph[T]) PrimMST(comparator pkg.Comparator[*Edge[T]]) map[*Edge[T]]string {
	// 哪些点被处理过
	nodeSet := make(map[*Node[T]]string, 0)

	// 初始化一个边的小根堆
	edgesMinHeap := minheap.NewMaxHeap[*Edge[T]](len(g.edges), comparator)

	// 哪些边被处理过（加入了堆）
	edgeSet := make(map[*Edge[T]]string, 0)
	// 依次挑选的的边在resultSet里
	resultSet := make(map[*Edge[T]]string, 0)

	// 随便挑了一个点,进入循环处理完后直接break
	// 随便挑一个点的实现，由一个点，解锁所有相连的边，需要打开下文的break
	for _, node := range g.nodes {
		// 该点是否已经被访问过
		if _, ok := nodeSet[node]; !ok {
			// 节点保留，标志为已经访问过
			nodeSet[node] = ""
			// 当前节点的所有与之相邻的边放到小根堆
			// 即由一个点，解锁所有相连的边
			for _, edge := range node.edges {
				// 这个边没有被处理，要加入到小根堆
				if _, ok := edgeSet[edge]; !ok {
					// 标记这个边被考虑过并加入到小根堆了
					edgeSet[edge] = ""
					// 这个边加入小根堆
					_ = edgesMinHeap.Push(edge)
				}
			}

			// 图的边还未全部考虑完，要依次考虑（贪心）
			for !edgesMinHeap.IsEmpty() {
				// 弹出这个点解锁的边中，最小的边（本质还是贪心）
				edge := edgesMinHeap.Pop()
				// 可能的一个新的点,from已经被考虑了，只需要看to
				toNode := edge.to
				// 不含有的时候，就是新的点
				if _, ok := nodeSet[toNode]; !ok {
					// 标记边的右节点被访问过
					nodeSet[toNode] = ""
					// 收集该边，是最小生成树需要的一个边
					resultSet[edge] = ""

					// 由toNode再蔓延开去，收集更多的边加入堆中，再贪心
					for _, nextEdge := range toNode.edges {
						// 没加过的，放入小根堆，并标记为已经处理过
						if _, ok := edgeSet[nextEdge]; !ok {
							edgeSet[nextEdge] = ""
							_ = edgesMinHeap.Push(edge)
						}
					}
				}
			}
		}
		// 直接break意味着我们不用考虑森林的情况，只需要从图的某个节点进入即可
		// 如果不加break我们可以兼容多个无向图的森林的生成树，但会增加循环，即图中的所有节点都进行一次prim
		// break;
	}
	return resultSet
}
