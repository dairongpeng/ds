package binarytree

import (
	"fmt"
	"testing"
)

// pre-order:1 2 4 5 3 6 7
// in-order:4 2 5 1 6 3 7
// post-order:4 5 2 6 7 3 1
func TestBinaryTreeR(t *testing.T) {
	tree := NewTree[int](1)
	tree.Root.Left = &Node[int]{Value: 2}
	tree.Root.Right = &Node[int]{Value: 3}
	tree.Root.Left.Left = &Node[int]{Value: 4}
	tree.Root.Left.Right = &Node[int]{Value: 5}
	tree.Root.Right.Left = &Node[int]{Value: 6}
	tree.Root.Right.Right = &Node[int]{Value: 7}

	preOrders := tree.PreOrder()
	inOrders := tree.InOrder()
	postOrders := tree.PostOrder()

	fmt.Print("pre-order:")
	for _, t := range preOrders {
		fmt.Printf("%d ", t)
	}
	fmt.Println()

	fmt.Print("in-order:")
	for _, t := range inOrders {
		fmt.Printf("%d ", t)
	}
	fmt.Println()

	fmt.Print("post-order:")
	for _, t := range postOrders {
		fmt.Printf("%d ", t)
	}
	fmt.Println()
}

// pre-order:1 2 4 5 3 6 7
// in-order:4 2 5 1 6 3 7
// post-order:4 5 2 6 7 3 1
// level-order:1 2 3 4 5 6 7
func TestBinaryTreeNonR(t *testing.T) {
	tree := NewTree[int](1)
	tree.Root.Left = &Node[int]{Value: 2}
	tree.Root.Right = &Node[int]{Value: 3}
	tree.Root.Left.Left = &Node[int]{Value: 4}
	tree.Root.Left.Right = &Node[int]{Value: 5}
	tree.Root.Right.Left = &Node[int]{Value: 6}
	tree.Root.Right.Right = &Node[int]{Value: 7}

	preOrders := tree.PreOrderNonRecursive()
	inOrders := tree.InOrderNonRecursive()
	postOrders := tree.PostOrderNonRecursive()
	levelOrders := tree.LevelOrder()

	fmt.Print("pre-order:")
	for _, t := range preOrders {
		fmt.Printf("%d ", t)
	}
	fmt.Println()

	fmt.Print("in-order:")
	for _, t := range inOrders {
		fmt.Printf("%d ", t)
	}
	fmt.Println()

	fmt.Print("post-order:")
	for _, t := range postOrders {
		fmt.Printf("%d ", t)
	}
	fmt.Println()

	fmt.Print("level-order:")
	for _, t := range levelOrders {
		fmt.Printf("%d ", t)
	}
	fmt.Println()
}
