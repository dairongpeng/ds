package mergesort

type MergeSorter[T any] struct {
	Arr []T
}

// NewMergeSorter 初始化一个归并排序结构
func NewMergeSorter[T any](values []T) *MergeSorter[T] {
	ms := &MergeSorter[T]{
		Arr: values,
	}

	return ms
}

// Sort 归并排序递归实现
func (ms *MergeSorter[T]) Sort(arr []T, cmp func(item1, item2 any) int) {
	// 空数组或者只存在1个元素
	if len(arr) < 2 {
		return
	}

	// 传入被排序数组，以及左右边界到递归函数
	process(arr, 0, len(arr)-1, cmp)
}

// process 使得数组arr的L到R位置变为有序
func process[T any](arr []T, L, R int, cmp func(item1, item2 any) int) {
	if L == R { // base case
		return
	}

	mid := L + (R-L)/2
	process(arr, L, mid, cmp)
	process(arr, mid+1, R, cmp)
	// 当前栈顶左右已经排好序，准备左右merge，注意这里的merge动作递归的每一层都会调用
	merge(arr, L, mid, R, cmp)
}

// merge arr L到M有序 M+1到R有序 变为arr L到R整体有序
func merge[T any](arr []T, L, M, R int, cmp func(item1, item2 any) int) {
	// merge过程申请辅助数组，准备copy
	help := make([]T, 0)
	p1 := L
	p2 := M + 1
	// p1未越界且p2未越界
	for p1 <= M && p2 <= R {
		if cmp(arr[p1], arr[p2]) <= 0 { // arr[p1] <= arr[p2]
			help = append(help, arr[p1])
			p1++
		} else {
			help = append(help, arr[p2])
			p2++
		}
	}

	// p2越界的情况, 把p1剩余部分直接追加
	for p1 <= M {
		help = append(help, arr[p1])
		p1++
	}

	// p1越界的情况, 把p2剩余部分直接追加
	for p2 <= R {
		help = append(help, arr[p2])
		p2++
	}

	// 把辅助数组help中整体merge后的有序数组，copy回原数组arr中去
	for j := 0; j < len(help); j++ {
		arr[L+j] = help[j]
	}
}
