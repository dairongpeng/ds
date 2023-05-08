package binarytree

import (
	"github.com/dairongpeng/ds/pkg"
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

// FindSuccessorNode 给定一个二叉搜索树的节点，求该节点的后继结点。后继结点是中序遍历中一个节点的下一个节点
// 1. 如果该节点的右子节点存在，那么该节点的后继节点是其右子节点的子树中最左侧的节点。
// 2. 如果该节点的右子节点不存在，则需要向上遍历其祖先节点，直到找到一个祖先节点，该祖先节点的左子节点是该节点的祖先节点之一。这个祖先节点就是该节点的后继节点。
// 3. 该算法常常被用在二叉搜索树中（平衡）所以该算法算法复杂度O(h)，h为树的高度。一个节点的后继结点是整个树节点中大于当前节点的最小节点
// 4. 二叉搜索树的中序遍历，就是一个从小到达排列的序列
// 5. 如果不是二叉搜索树，寻找某节点的后继，则可能需要给节点增加一个parent指针，或者对该树进行中序遍历，再线性查找。
func (t *Tree[T]) FindSuccessorNode(node *Node[T], comparator pkg.Comparator[T]) *Node[T] {
	root := t.Root

	// 如果右子树不为空，后继节点就是右子树的最左节点
	if node.Right != nil {
		curr := node.Right
		for curr.Left != nil {
			curr = curr.Left
		}
		return curr
	} else {
		var succ *Node[T]
		curr := root
		// 否则，从树的根节点开始遍历找。实际上就是二分查找，二分查找到了node，那么上一个维护的节点就是node的后继节点
		// 因为二叉搜索树，根节点划分了左右子树。
		for curr != nil {
			if comparator(curr.Value, node.Value) > 0 { // 如果当前节点比 node 大 curr.Value > node.Value
				// 则当前节点有可能为 node 的后继节点，
				// 将当前节点赋给后继节点变量 succ, 后面看情况决定要不要更新为更确切的后继节点
				// 然后向左子树遍历以寻找更准确的后继节点
				succ = curr
				curr = curr.Left
			} else if comparator(curr.Value, node.Value) < 0 { // 如果当前节点比 node 小 curr.Value < node.Value
				// 则当前节点不能为 node 的后继节点，
				// 向右子树继续寻找后继节点
				curr = curr.Right
			} else { // 否则说明找到了 node 节点，退出循环，标记其后继节点为当前找到的最新的后继节点
				break
			}
		}
		return succ
	}
}

// IsBalanced 判断一颗二叉树是不是平衡二叉树
func (t *Tree[T]) IsBalanced() bool {
	head := t.Root

	// 平衡信息
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

		// 左子树的平衡信息
		rightInfo := f(head.Left)
		// 右孩子的平衡信息
		leftInfo := f(head.Right)

		// 当前节点左右子树的最大高度
		var maxh int
		if leftInfo.height > rightInfo.height {
			maxh = leftInfo.height
		} else {
			maxh = rightInfo.height
		}
		// 当前节点的高度
		ch := maxh + 1

		// 当前节点确认的树是否是平衡的
		var cok = true
		if !leftInfo.ok || !rightInfo.ok || math.Abs(float64(leftInfo.height)-float64(rightInfo.height)) > 1 {
			cok = false
		}

		// 构建当前节点的平衡信息返回
		return &info{
			ok:     cok,
			height: ch,
		}
	}

	return f(head).ok
}

// FindMaxDistance 返回二叉树中任意两节点的最大距离
func (t *Tree[T]) FindMaxDistance() int {
	head := t.Root
	if head == nil {
		return 0
	}

	// 距离信息
	type info struct {
		// 当前节点为树根的情况下，该树的最大距离
		maxdistance int
		// 当前节点为树根的情况下，该树的高度
		height int
	}

	var f func(node *Node[T]) *info

	f = func(node *Node[T]) *info {
		if node == nil {
			return &info{
				maxdistance: 0,
				height:      0,
			}
		}

		// 左树信息
		leftInfo := f(node.Left)
		// 右树信息
		rightInfo := f(node.Right)

		// 用左右树的信息，加工当前节点自身的info
		// 自身的高度是，左右较大的高度加上自身节点高度1
		ch := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1
		// 自身最大距离，(左右树最大距离)和(左右树高度相加再加1)，求最大值
		cd := int(math.Max(
			math.Max(float64(leftInfo.maxdistance), float64(rightInfo.maxdistance)),
			float64(leftInfo.height+rightInfo.height+1)))
		// 自身的info返回
		return &info{
			maxdistance: cd,
			height:      ch,
		}
	}

	return f(head).maxdistance
}

// FindMaxSubBSTSize 查找二叉树中所有子树中含最多节点的二叉搜索子树节点的数量
func (t *Tree[T]) FindMaxSubBSTSize(comparator pkg.Comparator[T]) int {
	head := t.Root
	if head == nil {
		return 0
	}

	type info struct {
		// 以当前节点为头节点的树，整体是否是二叉搜索树
		IsAllBST bool
		// 最大的满足二叉搜索树条件的size
		MaxSubBSTSize int
		// 整棵树的最小值
		Min T
		// 整棵树的最大值
		Max T
	}

	var f func(node *Node[T]) *info

	f = func(node *Node[T]) *info {
		if node == nil {
			return nil
		}

		// 左子树返回的Info信息
		leftInfo := f(node.Left)
		// 右子树返回的Info信息
		rightInfo := f(node.Right)

		// 加工我自身的info。min维护左右手min和当前v的最小值； max维护左右树的max和当前v的最大值
		min, max := node.Value, node.Value
		// 左树不为空，加工min和max
		if leftInfo != nil {
			if comparator(min, leftInfo.Min) > 0 {
				min = leftInfo.Min
			}

			if comparator(max, leftInfo.Max) < 0 {
				max = leftInfo.Max
			}
		}

		// 右树不为空，加工min和max
		if rightInfo != nil {
			if comparator(min, rightInfo.Min) > 0 {
				min = rightInfo.Min
			}

			if comparator(max, rightInfo.Max) < 0 {
				max = rightInfo.Max
			}
		}

		// case1: 与node无关的情况。当前二叉树存在的最大搜索二叉树的最大大小，是左右树存在的最大二叉搜索树的较大的
		maxSubBSTSize := 0
		if leftInfo != nil {
			maxSubBSTSize = leftInfo.MaxSubBSTSize
		}
		if rightInfo != nil {
			maxSubBSTSize = int(math.Max(float64(maxSubBSTSize), float64(rightInfo.MaxSubBSTSize)))
		}
		// 如果当前节点为头的二叉树不是二叉搜索树，则当前Info信息中isAllBST为false
		isAllBST := false

		// case2：与node有关的情况
		// 左树整个是二叉搜索树么
		leftIsAllBST := false
		// 右树整个是二叉搜索树么
		rightIsAllBST := false
		// 左树最大值小于node的值是否
		leftMaxVLessNodeV := false
		// 右树的最小值，大于node的值是否
		rightMinMoreNodeV := false
		if leftInfo == nil {
			leftIsAllBST = true
			leftMaxVLessNodeV = true
		} else {
			leftIsAllBST = leftInfo.IsAllBST
			// leftMaxVLessNodeV = leftInfo.Max < node.Val
			if comparator(leftInfo.Max, node.Value) < 0 {
				leftMaxVLessNodeV = true
			} else {
				leftMaxVLessNodeV = false
			}
		}

		if rightInfo == nil {
			rightIsAllBST = true
			rightMinMoreNodeV = true
		} else {
			rightIsAllBST = rightInfo.IsAllBST
			// rightMinMoreNodeV = rightInfo.Min > node.Val
			if comparator(rightInfo.Min, node.Value) > 0 {
				rightMinMoreNodeV = true
			} else {
				rightMinMoreNodeV = false
			}
		}

		// 如果左树是二叉搜索树，右树也是二叉搜索树，当前节点为树根的左树最大值都比当前值小，当前节点为树根的右树最小值都比当前值大
		// 证明以当前节点node为树根的树，也是一个二叉搜索树。满足case2
		if leftIsAllBST && rightIsAllBST && leftMaxVLessNodeV && rightMinMoreNodeV {
			leftSize := 0
			rightSize := 0
			if leftInfo != nil {
				leftSize = leftInfo.MaxSubBSTSize
			}

			if rightInfo != nil {
				rightSize = rightInfo.MaxSubBSTSize
			}

			// 当前节点为树根的二叉搜索树的节点大小是左树存在的最大二叉搜索树的大小，加上右树存在的最大的二叉搜索树的大小，加上当前node节点1
			maxSubBSTSize = leftSize + rightSize + 1
			// 当前节点整个是二叉搜索树
			isAllBST = true
		}

		return &info{
			IsAllBST:      isAllBST,
			MaxSubBSTSize: maxSubBSTSize,
			Min:           min,
			Max:           max,
		}
	}

	return f(head).MaxSubBSTSize
}

// IsFull 判断一个树是否是满二叉树
func (t *Tree[T]) IsFull() bool {
	head := t.Root
	if head == nil {
		return true
	}

	type info struct {
		// 已x为头节点的树，高度是多少
		height int
		// 已x为头节点的树，整个树的节点个数是多少
		nodes int
	}

	var f func(node *Node[T]) *info

	f = func(node *Node[T]) *info {
		if node == nil {
			return &info{
				height: 0,
				nodes:  0,
			}
		}

		leftInfo := f(node.Left)
		rightInfo := f(node.Right)
		// 当前高度
		ch := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1
		// 当前节点为树根的二叉树所有节点数
		nodes := leftInfo.nodes + rightInfo.nodes + 1
		return &info{
			height: ch,
			nodes:  nodes,
		}
	}

	rootInfo := f(head)

	// 如果一个树的高度 * 2 - 1 等于节点数，那么这个树是满二叉树
	return (rootInfo.height*2 - 1) == rootInfo.nodes
}

// IsCBT 判断一个二叉树是不是一个完全二叉树
// 满二叉树：              完全二叉树：
//
//	    1                     1
//	2       3            2        3
//
// 4    5  6     7      4     5  6
func (t *Tree[T]) IsCBT() bool {
	if t.Root == nil {
		return true
	}
	head := t.Root

	type info struct {
		isFull bool
		isCBT  bool
		height int
	}

	var f func(node *Node[T]) *info

	f = func(node *Node[T]) *info {
		// 如果是空树，我们封装Info而不是返回为空
		// 好处是下文不需要额外增加判空处理
		if node == nil {
			return &info{
				isFull: true,
				isCBT:  true,
				height: 0,
			}
		}

		leftInfo := f(node.Left)
		rightInfo := f(node.Right)

		// 整合当前节点的info
		// 高度信息=左右树最大高度值+1
		currentHeight := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1
		// node是否是满二叉树信息 = 左右都是满且左右高度一样
		currentIsFull := leftInfo.isFull && rightInfo.isFull && leftInfo.height == rightInfo.height
		currentIsBST := false
		if currentIsFull { // 满二叉树一定是完全二叉树
			currentIsBST = true
		} else { // 以node为头整棵树，不是满二叉树，再检查是不是完全二叉树
			// 左右都是完全二叉树才有讨论的必要
			if leftInfo.isCBT && rightInfo.isCBT {
				// 第二种情况。左树是完全二叉树，右树是满二叉树，左树高度比右树高度大1
				//        1
				//    2       3
				// 4
				if leftInfo.isCBT && rightInfo.isFull && leftInfo.height == (rightInfo.height+1) {
					currentIsBST = true
				}
				// 第三种情况。左树满，右树满，且左树高度比右树高度大1
				//        1
				//    2       3
				// 4    5
				if leftInfo.isFull && rightInfo.isFull && leftInfo.height == (rightInfo.height+1) {
					currentIsBST = true
				}
				// 第四种情况。左树满，右树是完全二叉树，且左右树高度相同
				//        1
				//    2       3
				// 4    5  6
				if leftInfo.isFull && rightInfo.isCBT && leftInfo.height == rightInfo.height {
					currentIsBST = true
				}
			}
		}
		return &info{
			isFull: currentIsFull,
			isCBT:  currentIsBST,
			height: currentHeight,
		}
	}

	return f(head).isCBT
}

// LowestCommonAncestor 求二叉树中两个节点的最近公共祖先
func (t *Tree[T]) LowestCommonAncestor(p *Node[T], q *Node[T]) *Node[T] {
	if t.Root == nil {
		return t.Root
	}

	var f func(node *Node[T], p *Node[T], q *Node[T]) *Node[T]

	f = func(node *Node[T], p *Node[T], q *Node[T]) *Node[T] {
		// 如果当前树根为空，或者p和q中有等于 node的，那么它们的最近公共祖先即为node（一个节点也可以是它自己的祖先）
		if node == nil || p == node || q == node {
			return node
		}
		// 递归遍历左子树，只要在左子树中找到了p或q，则先找到谁就返回谁
		left := f(node.Left, p, q)
		// 递归遍历右子树，只要在右子树中找到了p或q，则先找到谁就返回谁
		right := f(node.Right, p, q)

		// left和 right均不为空时，说明 p、q节点分别在 node异侧, 最近公共祖先即为 node;
		// 如果 p 和 q 分别在左右子树中，或者根节点本身就是 p 或 q，则返回根节点
		if left != nil && right != nil {
			return node
		}

		// 如果 p 和 q 都在左子树中，则返回左子树的结果
		if left != nil {
			return left
		}

		// 如果 p 和 q 都在右子树中，则返回右子树的结果
		if right != nil {
			return right
		}

		// 这里表示，p和q不在已node为头节点的树中
		return nil
	}

	return f(t.Root, p, q)
}

// IsSymmetric 判断一个二叉树是不是镜面对称的
func (t *Tree[T]) IsSymmetric(comparator pkg.Comparator[T]) bool {
	if t.Root == nil {
		return true
	}

	// 原始树 node1
	// 镜面树 node2
	var f func(node1 *Node[T], node2 *Node[T]) bool

	f = func(node1 *Node[T], node2 *Node[T]) bool {
		// 当前镜像的节点都为空，也算合法的镜像
		if node1 == nil && node2 == nil {
			return true
		}

		// 互为镜像的两个点不为空
		if node1 != nil && node2 != nil {
			// 当前两个镜像点要是相等的，
			// A树的左树和B树的右树互为镜像且满足，且A树的右树和B树的左树互为镜像，且满足。
			// 那么当前的镜像点下面的都是满足的
			return comparator(node1.Value, node2.Value) == 0 &&
				f(node1.Left, node2.Right) && f(node1.Right, node2.Left)
		}
		// 一个为空，一个不为空 肯定不构成镜像  false
		return false
	}

	return f(t.Root, t.Root)
}
