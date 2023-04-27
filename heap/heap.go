package dsheap

// DSHeap 堆结构（完全二叉树的子集）
// 每个子树父节点都比孩子大称为大根堆
// 每个子树父节点都比孩子小称为小根堆
// 其他完全二叉树不能称为堆
type DSHeap[T any] interface {
	IsEmpty() bool
	IsFull() bool
	Push(value T) error
	Pop() (T, bool)
}
