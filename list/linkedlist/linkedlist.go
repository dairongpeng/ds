package linkedlist

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	Head *Node[T]
}

// New 初始化一个链表结构
func New[T any](values ...T) *List[T] {
	var l = &List[T]{}
	if len(values) > 0 {
		for _, v := range values {
			l.Add(v)
		}
	}
	return l
}

// Add 添加一个元素到链表中
func (l *List[T]) Add(v T) {
	newNode := &Node[T]{Value: v}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	newNode.Next = l.Head
	l.Head = newNode
}

// Remove 从链表中移出一个元素
func (l *List[T]) Remove() (T, bool) {
	if l.Head == nil {
		var zeroValue T
		return zeroValue, false
	}
	v := l.Head.Value
	l.Head = l.Head.Next
	return v, true
}

// Get 通过下标获取链表中的元素
func (l *List[T]) Get(index int) (T, bool) {
	var zeroValue T

	if index < 0 {
		return zeroValue, false
	}

	if l.Head == nil {
		return zeroValue, false
	}

	cur := l.Head
	for cur != nil {
		if index == 0 {
			return cur.Value, true
		}
		index--
		cur = cur.Next
	}

	return zeroValue, false
}

// Reverse 翻转链表
func (l *List[T]) Reverse() {
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	var prev *Node[T]
	current := l.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// 此时链表的头结点为prev
	l.Head = prev
}

// RemoveValue 移出链表中值等于target的节点
func (l *List[T]) RemoveValue(target T, cmp func(a, b any) int) {
	// 处理链表头结点的值即等于target的节点
	for l.Head != nil {
		// 头节点不等于target
		if cmp(l.Head.Value, target) != 0 {
			break
		}
		// 从头节点开始，等于target的节点先滤掉
		l.Head = l.Head.Next
	}

	// 1、链表中的节点值全部都等于target
	// 2、原始链表为nil
	if l.Head == nil {
		return
	}

	// head来到第一个不需要删除的位置, 查找pre应该指向的下一个节点是什么？
	pre := l.Head
	cur := l.Head
	for cur != nil {
		// 当前节点cur往下，有多少v等于target的节点，就删除多少节点
		if cmp(cur.Value, target) == 0 { // 当cur等于target就删除cur
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		// 当前节点向下滑动
		cur = cur.Next
	}
}

// HasCycle 检测链表是否成环
func (l *List[T]) HasCycle() bool {
	if l.Head == nil || l.Head.Next == nil {
		return false
	}

	// 快慢指针
	slow := l.Head
	fast := l.Head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		// 慢指针每次走一步
		// 快指针每次走两步
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 有环的话一定追的上，但不一定是第一次成环的节点
	return true
}

// RemoveLastKthNode 删除单链表的倒数第k个节点
func (l *List[T]) RemoveLastKthNode(lastKth int) {
	if l.Head == nil || lastKth < 1 {
		return
	}

	// cur指针也指向链表头节点
	cur := l.Head
	// 检查倒数第lastKth个节点的合法性
	for cur != nil {
		lastKth--
		cur = cur.Next
	}

	// 需要删除的是头结点
	if lastKth == 0 {
		l.Head = l.Head.Next
	}

	if lastKth < 0 {
		// cur回到头结点
		cur = l.Head
		for lastKth != 0 {
			lastKth++
			cur = cur.Next
		}

		// 此次cur就是要删除的前一个节点。把原cur.next删除
		cur.Next = cur.Next.Next
	}

	// lastKth > 0的情况，表示倒数第lastKth节点比原链表程度要大，即不存在
	return
}

// RemoveMidNode 删除链表中间节点
func (l *List[T]) RemoveMidNode() {
	// 无节点，或者只有一个节点的情况，直接返回
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	// 链表两个节点，删除第一个节点
	if l.Head.Next.Next == nil {
		// free first node mem
		l.Head = l.Head.Next
		return
	}

	pre := l.Head
	cur := l.Head.Next.Next

	// 快慢指针
	if cur.Next != nil && cur.Next.Next != nil {
		pre = pre.Next
		cur = cur.Next.Next
	}

	// 快指针走到尽头，慢指针奇数长度停留在中点，偶数长度停留在上中点。删除该节点
	pre.Next = pre.Next.Next

	return
}

// GetLoopNode 给定一个链表，如果成环，返回成环的那个节点
func (l *List[T]) GetLoopNode() *Node[T] {
	// 节点数目不足以成环，返回不存在成环节点
	if l.Head == nil || l.Head.Next == nil || l.Head.Next.Next == nil {
		return nil
	}

	slow := l.Head.Next
	fast := l.Head.Next.Next

	for slow != fast {
		// 快指针提前到达终点，该链表无环
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	// 确定成环，fast回到头节点
	fast = l.Head

	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}

	// 再次相遇节点，就是成环节点
	return slow
}

// Print 打印链表结构
func (l *List[T]) Print() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

// PrintCommonPart 打印两个有序链表的公共部分
func PrintCommonPart[T any](l1 *List[T], l2 *List[T], cmp func(a, b any) int) {
	fmt.Println("Common Part: ")

	for l1.Head != nil && l2.Head != nil {
		if cmp(l1.Head.Value, l2.Head.Value) < 0 { // l.Head.Value < ll.Head.Value
			l1.Head = l1.Head.Next
		} else if cmp(l1.Head.Value, l2.Head.Value) > 0 { // l.Head.Value > ll.Head.Value
			l2.Head = l2.Head.Next
		} else {
			fmt.Println(l1.Head.Value)
			l1.Head = l1.Head.Next
			l2.Head = l2.Head.Next
		}
	}
	fmt.Println()
}

// GetIntersectNode 两个无环链表是否相交, 若相交返回相交的第一个节点, 不相交返回false
func GetIntersectNode[T any](l1 *List[T], l2 *List[T]) (*Node[T], bool) {
	lenA, lenB := 0, 0
	curA, curB := l1.Head, l2.Head

	// 检查链表1的长度
	for curA != nil {
		lenA++
		curA = curA.Next
	}

	// 检查链表2的长度
	for curB != nil {
		lenB++
		curB = curB.Next
	}

	// 长度检查完成后，curA和curB回到两个链表的头结点
	curA, curB = l1.Head, l2.Head

	// 如果链表1的长度大一些
	if lenA > lenB {
		// 先让较长的链表curA追上距离差距
		for i := 0; i < lenA-lenB; i++ {
			curA = curA.Next
		}
	} else { // 如果链表2的长度大一些
		// 先让较长的链表curB追上距离差距
		for i := 0; i < lenB-lenA; i++ {
			curB = curB.Next
		}
	}

	// 此时curA和curB剩余的长度一样长。一起往下移动
	for curA != nil && curB != nil {
		// 一旦发现两个节点的地址相同，即为第一次成环的节点
		if curA == curB {
			return curA, true
		}
		curA = curA.Next
		curB = curB.Next
	}

	// 否则两个链表没有成环的节点
	return nil, false
}

// MergeTwoList 合并两个有序链表
func MergeTwoList[T any](l1, l2 *List[T], cmp func(a, b any) int) *List[T] {
	var L = &List[T]{}

	// base case
	if l1.Head == nil {
		L.Head = l2.Head
		return L
	}
	if l2.Head == nil {
		L.Head = l1.Head
		return L
	}

	var head *Node[T]

	// 选出两个链表较小的头作为整个合并后的头结点
	if cmp(l1.Head.Value, l2.Head.Value) <= 0 { // l1.Head.Value <= l2.Head.Value
		head = l1.Head
	} else {
		head = l2.Head
	}

	// 链表1的准备合并的节点，就是头结点的下一个节点
	cur1 := head.Next
	// 链表2的准备合并的节点，就是另一个链表的头结点
	var cur2 *Node[T]
	if head == l1.Head {
		cur2 = l2.Head
	} else {
		cur2 = l1.Head
	}

	// 使用临时变量pre向下移动，head作为头结点不变
	pre := head
	for cur1 != nil && cur2 != nil {
		if cmp(cur1.Value, cur2.Value) <= 0 { // 比较v1和v2选择小的 cur1.Val <= cur2.Val
			pre.Next = cur1
			// cur1向下滑动
			cur1 = cur1.Next
		} else { // cur1.Val > cur2.Val
			pre.Next = cur2
			// cur2向下滑动
			cur2 = cur2.Next
		}

		// 选中节点后，pre向下滑动，进行下一轮选择
		pre = pre.Next
	}

	// 有一个链表耗尽了，没耗尽的链表直接拼上
	if cur1 != nil {
		pre.Next = cur1
	} else {
		pre.Next = cur2
	}

	L.Head = head
	return L
}
