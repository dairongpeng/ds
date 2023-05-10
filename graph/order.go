package graph

/** 图的遍历算法 **/

import (
	"github.com/dairongpeng/ds/queue/arrayqueue"
	"github.com/dairongpeng/ds/stack/arraystack"
)

// Bfs 从图中选择一个node，从node出发，对图进行宽度优先遍历, 借助队列
// 选择的node不同，结果一般也不同，选择起始节点遍历图应该根据具体问题的要求来进行决策。
func (node *Node[T]) Bfs() []T {
	if node == nil {
		return nil
	}
	bfsorder := make([]T, 0)

	queue := arrayqueue.New[*Node[T]]()
	// 图需要用set结构，因为图相比于二叉树有可能存在环
	// 即有可能存在某个点多次进入队列的情况。使用Set可以防止相同节点重复进入队列
	Set := make(map[*Node[T]]string, 0)
	queue.Enqueue(node)
	Set[node] = ""

	for !queue.IsEmpty() {
		// 出队列
		cur, _ := queue.Dequeue()

		bfsorder = append(bfsorder, cur.value)
		for _, next := range cur.nexts {
			// 直接邻居，没有进入过Set的进入Set和队列
			// 用set限制队列的元素，防止有环队列一直会加入元素
			if _, ok := Set[next]; !ok { // Set中不存在, 则加入队列
				Set[next] = ""
				queue.Enqueue(next)
			}
		}
	}

	return bfsorder
}

// Dfs 从图中选择一个node，从node出发，对图进行深度优先遍历。借助栈
// 选择的node不同，结果一般也不同，选择起始节点遍历图应该根据具体问题的要求来进行决策。
func (node *Node[T]) Dfs() []T {
	if node == nil {
		return nil
	}
	dfsorder := make([]T, 0)

	stack := arraystack.New[*Node[T]]()
	// Set的作用和宽度优先遍历类似，保证重复的点不要进栈
	set := make(map[*Node[T]]string, 0)
	// 进栈
	stack.Push(node)
	set[node] = ""
	// 收集的时机是在进栈的时候
	dfsorder = append(dfsorder, node.value)

	for !stack.IsEmpty() {
		cur, _ := stack.Pop()

		// 枚举当前弹出节点的后代
		for _, next := range cur.nexts {
			// 只要某个后代没进入过栈，进栈
			if _, ok := set[next]; !ok {
				// 把该节点的父亲节点重新压回栈中
				stack.Push(cur)
				// 再把自己压入栈中
				stack.Push(next)
				set[next] = ""
				// 收集当前节点的值
				dfsorder = append(dfsorder, next.value)
				// 直接break，此时栈顶是当前next节点，达到深度优先的目的
				break
			}
		}
	}

	return dfsorder
}

// Topology 有向无环图图DAG的拓扑排序, 返回拓扑排序的顺序list; 可以变种用来检查一张图是否存在环
func (g *Graph[T]) Topology() []*Node[T] {
	// 提取入度信息：节点->入度
	inMap := make(map[*Node[T]]int)
	// 提取入度为零的节点信息，剩余入度为0的点，才能进这个队列
	zeroInQueue := arrayqueue.New[*Node[T]]()
	// 拿到该图中所有的点集
	for _, node := range g.nodes {
		// 初始化每个点，每个点的入度是原始节点的入度信息
		// 加入inMap
		inMap[node] = node.in
		// 由于是有向无环图，则必定有入度为0的起始点。放入到zeroInQueue
		if node.in == 0 {
			zeroInQueue.Enqueue(node)
		}
	}

	// 拓扑排序的结果，依次加入result
	result := make([]*Node[T], 0)

	for !zeroInQueue.IsEmpty() {
		// 该有向无环图初始入度为0的点，直接出队放入结果集中
		cur, _ := zeroInQueue.Dequeue()
		result = append(result, cur)
		// 该节点的下一层邻居节点，入度减一且加入到入度的map中
		for _, next := range cur.nexts {
			inMap[next] = inMap[next] - 1
			// 如果下一层存在入度变为0的节点，加入到0入度的队列中
			if inMap[next] == 0 {
				zeroInQueue.Enqueue(next)
			}
		}
	}
	// 如果所有节点都被加入到结果序列中，则说明该有向图是有向无环图，否则就是有环的。
	// 1. 不存在入度为0的节点了，存在环，这时result结果节点比图节点少
	// 2. 所有节点被加入了result，result节点和图节点一样多
	return result
}
