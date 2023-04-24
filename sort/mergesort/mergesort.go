package mergesort

import "math"

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
func (ms *MergeSorter[T]) Sort(cmp func(item1, item2 any) int) {
	// 空数组或者只存在1个元素
	if len(ms.Arr) < 2 {
		return
	}

	// 传入被排序数组，以及左右边界到递归函数
	process(ms.Arr, 0, len(ms.Arr)-1, cmp)
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

// SortNonRecursive 非递归会使用循环进行迭代计算
func (ms *MergeSorter[T]) SortNonRecursive(cmp func(item1, item2 any) int) {
	// 空数组或者只存在1个元素
	if len(ms.Arr) < 2 {
		return
	}

	N := len(ms.Arr)
	// 当前有序的，左组长度, 那么实质分组大小是从2开始的
	// 已经有序的容量大小，类似于递归的出口
	mergeSize := 1
	for mergeSize < N {
		// L表示当前分组的左组的位置，初始为第一个分组的左组位置为0
		L := 0
		for L < N {
			// L...M  当前左组（mergeSize）
			M := L + mergeSize - 1
			// 当前左组包含当前分组的所有元素，即没有右组了，无需merge已经有序
			if M >= N {
				break
			}
			//  L...M为左组   M+1...R(mergeSize)为右组。
			//  右组够mergeSize个的时候，右坐标为M + mergeSize，右组不够的情况下右组边界坐标为整个数组右边界N - 1
			R := math.Min(float64(M+mergeSize), float64(N-1))
			// 把当前组进行merge
			merge(ms.Arr, L, M, int(R), cmp)
			L = int(R) + 1
		}
		// 如果mergeSize乘2必定大于N，直接break。
		// 防止mergeSize溢出，有可能N很大，下面乘2有可能范围溢出（整形数大于21亿）
		if mergeSize > N/2 {
			break
		}
		mergeSize *= 2
	}
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
