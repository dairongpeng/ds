package maxheap

import (
	"errors"
	"github.com/dairongpeng/ds/pkg"
)

// MaxHeap 堆结构也被称为优先级队列
type MaxHeap[T any] struct {
	// 大根堆底层数组结构
	heap []T
	// 分配给堆的空间限制
	limit int
	// 表示目前这个堆收集了多少个数，即堆大小。也表示添加的下一个数应该放在哪个位置
	heapSize int
}

// NewMaxHeap 初始化一个大根堆结构
func NewMaxHeap[T any](limit int) *MaxHeap[T] {
	maxHeap := &MaxHeap[T]{
		heap:     make([]T, 0),
		limit:    limit,
		heapSize: 0,
	}
	return maxHeap
}

func (h *MaxHeap[T]) IsEmpty() bool {
	return len(h.heap) == 0
}

func (h *MaxHeap[T]) IsFull() bool {
	return h.heapSize == h.limit
}

func (h *MaxHeap[T]) Push(value T, cmp pkg.Comparator[T]) error {
	if h.heapSize == h.limit {
		return errors.New("heap is full")
	}

	h.heap[h.heapSize] = value
	// heapSize的位置保存当前value
	up(h.heap, h.heapSize, cmp)
	h.heapSize++
	return nil
}

// Pop 返回堆中的最大值，并且在大根堆中，把最大值删掉。弹出后依然保持大根堆的结构
func (h *MaxHeap[T]) Pop(cmp pkg.Comparator[T]) T {
	// 弹出堆顶元素的实现为
	// 1. 交换堆顶和队列末尾元素。
	// 2. 堆大小减一
	// 3. 交换上来的堆顶元素，进行下沉down操作，去到合适的位置
	tmp := h.heap[0]
	h.heapSize--
	h.heap[0], h.heap[h.heapSize] = h.heap[h.heapSize], h.heap[0]
	down(h.heap, 0, h.heapSize, cmp)
	return tmp
}

// 往堆上添加数，需要从当前位置找父节点比较。实质上是从数组的末尾添加节点，往整个树根节点方向去PK
func up[T any](arr []T, index int, cmp pkg.Comparator[T]) {
	for cmp(arr[index], arr[(index-1)/2]) > 0 { // arr[index] > arr[(index-1)/2]
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

// 从index位置，不断的与左右孩子比较，下沉。下沉终止条件为：1. 左右孩子都不大于当前值 2. 没有左右孩子了
func down[T any](arr []T, index int, heapSize int, cmp pkg.Comparator[T]) {
	// 左孩子的位置
	left := index*2 + 1
	// 左孩子越界，右孩子一定越界。退出循环的条件是：
	// 1. 左孩子越界（左右孩子越界）
	for left < heapSize {
		var largestIdx int
		rigth := left + 1
		// 存在右孩子，且右孩子的值比左孩子大，选择右孩子的位置
		if rigth < heapSize && cmp(arr[rigth], arr[left]) > 0 { // rigth < heapSize && arr[rigth] > arr[left]
			largestIdx = rigth
		} else {
			largestIdx = left
		}

		// 1. 左右孩子的最大值都不大于当前值，终止寻找。无需继续下沉
		if cmp(arr[largestIdx], arr[index]) <= 0 { // arr[largestIdx] <= arr[index]
			break
		}
		// 左右孩子的最大值大于当前值
		arr[largestIdx], arr[index] = arr[index], arr[largestIdx]
		// 当前位置移动到交换后的位置，继续寻找
		index = largestIdx
		// 移动后左孩子理论上的位置，下一次循环判断越界情况
		left = index*2 + 1
	}
}

// HeapSort 堆排序额外空间复杂度O(1)。时间复杂度为 O(nlogn)
// 整个堆排序的过程，都只需要极个别临时存储空间，所以堆排序是原地排序算法。
// 堆排序不是稳定的排序算法，因为在排序的过程，存在将堆的最后一个节点跟堆顶节点互换的操作，所以就有可能改变值相同数据的原始相对顺序。
// 1. 建堆过程的时间复杂度是 O(n)
// 2. 排序时间复杂度为 O(nlogn)。整体O(nlogn)
func HeapSort[T any](arr []T, cmp pkg.Comparator[T]) {
	if len(arr) < 2 {
		return
	}

	// 原始版本, 调整arr满足大根堆结构。O(N*logN)
	//for i := 0; i < len(arr); i++ { // O(N)
	//	heapInsert(arr, i) // O(logN)
	//}

	// 优化版本：heapInsert改为heapify。从末尾开始看是否需要heapify=》O(N)复杂度。
	// 但是这只是优化了原有都是构建堆（O(NlogN)），最终的堆排序仍然是O(NlogN)。比原始版本降低了常数项
	for i := len(arr) - 1; i >= 0; i-- {
		down(arr, i, len(arr), cmp)
	}

	// 实例化一个大根堆,此时arr已经是调整后满足大根堆结构的arr
	mh := MaxHeap[T]{
		heap:     arr,
		limit:    len(arr),
		heapSize: len(arr),
	}

	mh.heapSize--
	arr[0], arr[mh.heapSize] = arr[mh.heapSize], arr[0]
	// O(N*logN)
	for mh.heapSize > 0 { // O(N)
		down(arr, 0, mh.heapSize, cmp) // O(logN)
		mh.heapSize--
		arr[0], arr[mh.heapSize] = arr[mh.heapSize], arr[0] // O(1)
	}

}
