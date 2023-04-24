package quicksort

import (
	"math/rand"
)

type QuickSorter[T any] struct {
	Arr []T
}

// NewQuickSorter 初始化一个快速排序结构
func NewQuickSorter[T any](values []T) *QuickSorter[T] {
	ms := &QuickSorter[T]{
		Arr: values,
	}

	return ms
}

// Sort 快速排序递归实现
func (qs *QuickSorter[T]) Sort(cmp func(item1, item2 any) int) {
	// 空数组或者只存在1个元素
	if len(qs.Arr) < 2 {
		return
	}

	sortByNetherlandsFlag(qs.Arr, 0, len(qs.Arr)-1, cmp)
}

// sortByNetherlandsFlag 通过荷兰国旗问题，解决快排partition
// 一次partition可以搞定一批位置。小于标志位的区域；等于标志位的区域；大于标志位的区域
func sortByNetherlandsFlag[T any](arr []T, L int, R int, cmp func(item1, item2 any) int) {
	if L >= R {
		return
	}

	// 随机选择排序因子，交换到arr R位置作为基准。达到算法稳定的目的
	arr[L+(int)(rand.Float64()*float64(R-L+1))], arr[R] = arr[R], arr[L+(int)(rand.Float64()*float64(R-L+1))]

	// 每次partition返回等于区域的范围, 荷兰国旗问题
	equalArea := netherlandsFlag(arr, L, R, cmp)
	// 对等于区域左边的小于区域递归，partition
	sortByNetherlandsFlag(arr, L, equalArea[0]-1, cmp)
	// 对等于区域右边的大于区域递归，partition
	sortByNetherlandsFlag(arr, equalArea[1]+1, R, cmp)
}

// arr[L...R] 玩荷兰国旗问题的划分，以arr[R]做划分值
// 小于arr[R]放左侧  等于arr[R]放中间  大于arr[R]放右边
// 返回中间区域的左右边界
func netherlandsFlag[T any](arr []T, L, R int, cmp func(item1, item2 any) int) []int {
	// 不存在荷兰国旗问题
	if L > R {
		return []int{-1, -1}
	}

	// 已经都是等于区域，由于用R做划分返回R位置
	if L == R {
		return []int{L, R}
	}

	// < 区 右边界
	less := L - 1
	// > 区 左边界
	more := R
	index := L
	for index < more {
		// 当前值等于右边界，不做处理，index++
		if cmp(arr[index], arr[R]) == 0 { // arr[index] == arr[R]
			index++
		} else if cmp(arr[index], arr[R]) < 0 { // 小于交换当前值和左边界的值 arr[index] < arr[R]
			less++
			arr[index], arr[less] = arr[less], arr[index]
			index++
		} else { // 大于右边界的值
			more--
			arr[index], arr[more] = arr[more], arr[index]
		}
	}
	// 比较完之后，把R位置的数，调整到等于区域的右边，至此大于区域才是真正意义上的大于区域
	arr[more], arr[R] = arr[R], arr[more]
	return []int{less + 1, more}
}

// QuickSort 快速实现版本
// 1. 一次partition只搞定了一个位置
// 2. 没有随机选择排序因子
// 3. 比较好理解，理解了这个再去看上文会清晰一些
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	left, right := 0, len(arr)-1
	// 选arr[0]作为基准，循环从1开始
	for i := 1; i <= right; {
		// 交换到左区域
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
			i++
		} else if arr[i] > pivot { // 交换到右区域
			arr[right], arr[i] = arr[i], arr[right]
			right--
		} else {
			i++
		}
	}

	// 再对左区域和右区域递归
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}
