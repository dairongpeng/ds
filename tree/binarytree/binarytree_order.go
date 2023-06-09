package binarytree

import (
	"github.com/dairongpeng/ds/pkg"
	"github.com/dairongpeng/ds/queue/arrayqueue"
	"github.com/dairongpeng/ds/stack/arraystack"
)

// Node 二叉树节点
type Node[T any] struct {
	// 二叉树节点上的值
	Value T
	// 左孩子
	Left *Node[T]
	// 右孩子
	Right *Node[T]
}

type Tree[T any] struct {
	Root *Node[T]
}

func NewTree[T any](value T) *Tree[T] {
	root := &Node[T]{
		Value: value,
	}

	tree := &Tree[T]{
		Root: root,
	}

	return tree
}

// PreOrder 给定二叉树头节点，先序遍历该二叉树
func (t *Tree[T]) PreOrder() []T {
	var result []T
	if t.Root == nil {
		return result
	}

	var f func(node *Node[T])
	f = func(node *Node[T]) {
		if node == nil {
			return
		}

		// 获取头节点，收集或打印该头结点
		result = append(result, node.Value)
		// 递归遍历左子树
		f(node.Left)
		// 递归遍历右子树
		f(node.Right)
	}

	f(t.Root)

	return result
}

// InOrder 给定二叉树头节点，中序遍历该二叉树
func (t *Tree[T]) InOrder() []T {
	var result []T
	if t.Root == nil {
		return result
	}
	var f func(node *Node[T])
	f = func(node *Node[T]) {
		if node == nil {
			return
		}

		// 递归遍历左子树
		f(node.Left)
		// 获取头节点，收集或打印该头结点
		result = append(result, node.Value)
		// 递归遍历右子树
		f(node.Right)
	}

	f(t.Root)

	return result
}

// PostOrder 给定二叉树头节点，后序遍历该二叉树
func (t *Tree[T]) PostOrder() []T {
	var result []T
	if t.Root == nil {
		return result
	}

	var f func(node *Node[T])
	f = func(node *Node[T]) {
		if node == nil {
			return
		}

		// 递归遍历左子树
		f(node.Left)
		// 递归遍历右子树
		f(node.Right)
		// 获取头节点，收集或打印该头结点
		result = append(result, node.Value)
	}

	f(t.Root)

	return result
}

// PreOrderNonRecursive 非递归先序
func (t *Tree[T]) PreOrderNonRecursive() []T {
	var result []T

	if t.Root != nil {
		stack := arraystack.New[*Node[T]]()
		// 入栈
		stack.Push(t.Root)
		for !stack.IsEmpty() {
			// 出栈
			node, _ := stack.Pop()
			result = append(result, node.Value)
			// 右孩子入栈
			if node.Right != nil {
				stack.Push(node.Right)
			}
			// 左孩子入栈
			if node.Left != nil {
				stack.Push(node.Left)
			}
		}
	}

	return result
}

// InOrderNonRecursive 非递归中序
func (t *Tree[T]) InOrderNonRecursive() []T {
	var result []T

	if t.Root != nil {
		root := t.Root
		stack := arraystack.New[*Node[T]]()
		for !stack.IsEmpty() || root != nil {
			// 整条左边界依次入栈
			if root != nil {
				stack.Push(root)
				root = root.Left
			} else { // 左边界到头弹出一个收集或打印，来到该节点右节点，再把该节点的左树以此进栈
				root, _ = stack.Pop()
				result = append(result, root.Value)
				root = root.Right
			}
		}
	}

	return result
}

// PostOrderNonRecursive 非递归后序
func (t *Tree[T]) PostOrderNonRecursive() []T {
	var result []T

	if t.Root != nil {
		root := t.Root

		// 借助两个辅助栈
		stack1 := arraystack.New[*Node[T]]()
		stack2 := arraystack.New[*Node[T]]()
		stack1.Push(root)

		for !stack1.IsEmpty() {
			// 出栈
			root, _ = stack1.Pop()
			stack2.Push(root)
			if root.Left != nil {
				stack1.Push(root.Left)
			}
			if root.Right != nil {
				stack1.Push(root.Right)
			}
		}

		for !stack2.IsEmpty() {
			node, _ := stack2.Pop()
			result = append(result, node.Value)
		}
	}

	return result
}

// LevelOrder 层级遍历
func (t *Tree[T]) LevelOrder() []T {
	if t.Root == nil {
		return nil
	}
	var result []T

	root := t.Root
	// 初始化队列
	queue := arrayqueue.New[*Node[T]]()
	// 加入头结点
	queue.Enqueue(root)

	// 队列不为空出队打印，把当前节点的左右孩子加入队列
	// 每一层的每个节点，都先入队左孩子，后入队右孩子，整体出队顺序就是按层遍历
	for !queue.IsEmpty() {
		// 弹出队列头部的元素
		cur, _ := queue.Dequeue()
		result = append(result, cur.Value)
		if cur.Left != nil {
			queue.Enqueue(cur.Left)
		}
		if cur.Right != nil {
			queue.Enqueue(cur.Right)
		}
	}

	return result
}

// BuildTreeByPreAndInOrder 根据二叉树的先序、中序序列，构建还原该树。（也可以使用中序和后序结合来构建）
// 注意：
// 1. 仅凭先序无法还原该树
// 例如先序序列为[1, 2, 4, 3, 5, 6]，可能对应的是如下两个树
//
//	     1                  1
//	  2     5            2     3
//	4   3 6   nil    nil   4 5    6
//
// 2. 树的节点Value需要唯一，否则无法分清左右树的边界。如果Value不唯一，需要增加额外信息确定唯一。或者再引入二叉树的后续序列，三种序列一块确定和构建该树。
func BuildTreeByPreAndInOrder[T any](preorder []T, inorder []T, comparator pkg.Comparator[T]) *Tree[T] {
	if len(preorder) == 0 {
		return nil
	}
	tree := &Tree[T]{}

	var f func(preorder []T, inorder []T, comparator pkg.Comparator[T]) *Node[T]

	f = func(preorder []T, inorder []T, comparator pkg.Comparator[T]) *Node[T] {
		if len(preorder) == 0 {
			return nil
		}

		headValue := preorder[0]
		headIndex := 0

		// 通过preorder和inorder结合，可以找到当前子树头结点的位置
		for i, v := range inorder {
			if comparator(v, headValue) == 0 { // v == headValue
				headIndex = i
				break
			}
		}

		leftInorder := inorder[:headIndex]
		rightInorder := inorder[headIndex+1:]

		leftPreorder := preorder[1 : len(leftInorder)+1]
		rightPreorder := preorder[len(leftInorder)+1:]

		head := &Node[T]{
			Value: headValue,
		}

		head.Left = f(leftPreorder, leftInorder, comparator)
		head.Right = f(rightPreorder, rightInorder, comparator)

		return head
	}

	tree.Root = f(preorder, inorder, comparator)

	return tree
}
