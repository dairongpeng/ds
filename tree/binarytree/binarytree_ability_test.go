package binarytree

import (
	"github.com/dairongpeng/ds/pkg"
	"testing"
)

func TestTree_IsBalanced(t1 *testing.T) {
	type testCase[T any] struct {
		name string
		t    Tree[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "case1_1",
			t: Tree[int]{
				Root: &Node[int]{Value: 1},
			},
			want: true,
		},
		{
			name: "case2_12",
			t: Tree[int]{
				Root: &Node[int]{Value: 1, Left: &Node[int]{Value: 2}},
			},
			want: true,
		},
		{
			name: "case2_123_all_left",
			t: Tree[int]{
				Root: &Node[int]{Value: 1, Left: &Node[int]{Value: 2, Left: &Node[int]{Value: 3}}},
			},
			want: false,
		},
		{
			name: "case2_1234",
			t: Tree[int]{
				Root: &Node[int]{Value: 1, Left: &Node[int]{Value: 2, Left: &Node[int]{Value: 4}}, Right: &Node[int]{Value: 3}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.IsBalanced(); got != tt.want {
				t1.Errorf("IsBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSuccessorNode(t *testing.T) {
	// 构建以下二叉树
	//     5
	//    / \
	//   2   8
	//  / \   \
	// 1   3   10
	//
	// inorder: 1 2 3 5 8 10
	root := &Node[int]{
		Value: 5,
		Left: &Node[int]{
			Value: 2,
			Left: &Node[int]{
				Value: 1,
			},
			Right: &Node[int]{
				Value: 3,
			},
		},
		Right: &Node[int]{
			Value: 8,
			Right: &Node[int]{
				Value: 10,
			},
		},
	}

	tree := &Tree[int]{
		Root: root,
	}

	// 情况1：查找一个有右子节点的节点的后继节点
	node1 := root.Left
	expected1 := root.Left.Right
	if res := tree.FindSuccessorNode(node1, pkg.NumberComparator[int]); res != expected1 {
		t.Errorf("expected %v but got %v", expected1, res)
	}

	// 情况2：查找一个没有右子节点的节点的后继节点
	node2 := root.Left.Left
	expected2 := root.Left
	if res := tree.FindSuccessorNode(node2, pkg.NumberComparator[int]); res != expected2 {
		t.Errorf("expected %v but got %v", expected2, res)
	}

	// 情况3：查找一个最右侧节点的后继节点
	node3 := root.Right.Right
	var expected3 *Node[int]
	if res := tree.FindSuccessorNode(node3, pkg.NumberComparator[int]); res != expected3 {
		t.Errorf("expected %v but got %v", expected3, res)
	}

	// 情况4：查找根节点的后继节点
	node4 := root
	expected4 := root.Right
	if res := tree.FindSuccessorNode(node4, pkg.NumberComparator[int]); res != expected4 {
		t.Errorf("expected %v but got %v", expected4, res)
	}

	// 情况5：查找一个不存在的节点的后继节点
	node5 := &Node[int]{
		Value: 100,
	}
	var expected5 *Node[int]
	if res := tree.FindSuccessorNode(node5, pkg.NumberComparator[int]); res != expected5 {
		t.Errorf("expected %v but got %v", expected5, res)
	}
}
