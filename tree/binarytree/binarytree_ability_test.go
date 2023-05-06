package binarytree

import "testing"

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
