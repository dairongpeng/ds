package binarytree

import (
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

		// 获取头节点，打印该头结点
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
		// 获取头节点，打印该头结点
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
		// 获取头节点，打印该头结点
		result = append(result, node.Value)
	}

	f(t.Root)

	return result
}

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

func (t *Tree[T]) InOrderNonRecursive() {

}

func (t *Tree[T]) PostOrderNonRecursive() {

}
