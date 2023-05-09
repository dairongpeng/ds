package unionfind

import "testing"

func TestUnionFind(t *testing.T) {
	type args struct {
		a int
		b int
	}

	type test struct {
		name string
		args args
		want bool
	}

	tests := []test{
		{
			name: "case1",
			args: args{
				a: 1,
				b: 8,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				a: 1,
				b: 9,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				a: 3,
				b: 8,
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				a: 2,
				b: 8,
			},
			want: false,
		},
	}

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	uf := NewUnionFind[int](values)

	uf.Union(1, 3)
	uf.Union(7, 9)
	uf.Union(8, 3)

	for _, tt := range tests {
		if got := uf.Find(tt.args.a, tt.args.b); got != tt.want {
			t.Errorf("Find(1, 8) = %v, want %v | Case = %s", got, tt.want, tt.name)
		}
	}

}
