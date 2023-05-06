package binarytree

import (
	"github.com/dairongpeng/ds/queue/arrayqueue"
	"math"
)

// MaxWidthUseMap 找到该二叉树的最大宽度，本质使用二叉树的按层遍历，借助map结构实现
func (t *Tree[T]) MaxWidthUseMap() int {
	if t.Root == nil {
		return 0
	}
	hd := t.Root

	queue := arrayqueue.New[*Node[T]]()
	queue.Enqueue(hd)

	// 节点和节点所在层的映射
	levelMap := make(map[*Node[T]]int, 0)
	// 头节点head属于第一层
	levelMap[hd] = 1
	// 当前正在统计哪一层的宽度
	curLevel := 1
	// 当前curLevel层，宽度目前是多少
	curLevelNodes := 0
	// 用来保存所有层的最大宽度的值
	max := 0
	for !queue.IsEmpty() {
		cur, _ := queue.Dequeue()
		curNodeLevel := levelMap[cur]
		// 当前节点的左孩子不为空，队列加入左孩子，层数在之前层上加1
		if cur.Left != nil {
			levelMap[cur.Left] = curNodeLevel + 1
			queue.Enqueue(cur.Left)
		}
		// 当前节点的右孩子不为空，队列加入右孩子，层数也变为当前节点的层数加1
		if cur.Right != nil {
			levelMap[cur.Right] = curNodeLevel + 1
			queue.Enqueue(cur.Right)
		}
		// 当前层等于正在统计的层数，不结算，宽度增加
		if curNodeLevel == curLevel {
			curLevelNodes++
		} else {
			// 新的一层，需要结算
			// 得到目前为止的最大宽度，等于原本收集的最大宽度和本层收集的宽度的较大者
			max = int(math.Max(float64(max), float64(curLevelNodes)))
			// 当前统计的层要进行到新的一层级
			curLevel++
			// 结算后，当前层节点数设置为1，因为已经来到新的一层的第一个节点了
			curLevelNodes = 1
		}
	}
	// 由于最后一层，没有新的一层去结算和对比最后一层收集到宽度curLevelNodes，所以这里单独结算最后一层
	max = int(math.Max(float64(max), float64(curLevelNodes)))
	return max
}

// MaxWidthNoMap 找到该二叉树的最大宽度，不借助map实现
func (t *Tree[T]) MaxWidthNoMap() int {
	if t.Root == nil {
		return 0
	}
	hd := t.Root

	queue := arrayqueue.New[*Node[T]]()
	queue.Enqueue(hd)

	// 当前层，最右节点是谁，初始root的就是本身
	curEnd := t.Root
	// 如果有下一层，下一层最右节点是谁
	var nextEnd *Node[T] = nil
	// 全局最大宽度
	max := 0
	// 当前层的节点数
	curLevelNodes := 0
	for !queue.IsEmpty() {
		cur, _ := queue.Dequeue()
		// 左边不等于空，加入左
		if cur.Left != nil {
			queue.Enqueue(cur.Left)
			// 孩子的最右节点暂时为左节点
			nextEnd = cur.Left
		}
		// 右边不等于空，加入右
		if cur.Right != nil {
			queue.Enqueue(cur.Right)
			// 如果有右节点，孩子层的最右要更新为右节点
			nextEnd = cur.Right
		}
		// 由于最开始弹出当前节点，那么该层的节点数加一
		curLevelNodes++
		// 当前节点是当前层最右的节点，进行结算
		if cur == curEnd {
			// 当前层的节点和max进行比较，计算当前最大的max
			max = int(math.Max(float64(max), float64(curLevelNodes)))
			// 即将进入下一层，重置下一层节点为0个节点
			curLevelNodes = 0
			// 当前层的最右，直接更新为找出来的下一层最右
			curEnd = nextEnd
		}
	}
	return max
}

// IsBalanced 判断一颗二叉树是不是平衡二叉树
func (t *Tree[T]) IsBalanced() bool {
	head := t.Root

	type info struct {
		ok     bool
		height int
	}

	var f func(head *Node[T]) *info

	f = func(head *Node[T]) *info {
		if head == nil {
			i := &info{
				ok:     true,
				height: 0,
			}

			return i
		}

		leftInfo := f(head.Right)
		rightInfo := f(head.Left)

		var maxh int
		if leftInfo.height > rightInfo.height {
			maxh = leftInfo.height
		} else {
			maxh = rightInfo.height
		}
		ch := maxh + 1

		var cok = true
		if !leftInfo.ok || !rightInfo.ok || math.Abs(float64(leftInfo.height)-float64(rightInfo.height)) > 1 {
			cok = false
		}

		return &info{
			ok:     cok,
			height: ch,
		}
	}

	return f(head).ok
}
