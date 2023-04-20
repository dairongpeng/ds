package arrayqueue

import "fmt"

type Queue[T any] []T

// New 初始化一个队列
func New[T any](values ...T) *Queue[T] {
	var q = &Queue[T]{}
	if len(values) > 0 {
		for _, v := range values {
			q.Enqueue(v)
		}
	}
	return q
}

// Enqueue 从队尾加入一个元素
func (q *Queue[T]) Enqueue(v T) {
	*q = append(*q, v)
}

// Dequeue 从队头弹出一个元素，如果队列为空则返回一个零值和false
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(*q) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, true
}

// Front 查看对头的元素，不出队。如果队列为空则返回一个零值和false
func (q *Queue[T]) Front() (T, bool) {
	if len(*q) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	val := (*q)[0]
	return val, true
}

// Size 返回队列的元素个数
func (q *Queue[T]) Size() int {
	return len(*q)
}

// IsEmpty 判断队列是否为空
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// Print 打印队列的元素
func (q *Queue[T]) Print() {
	fmt.Println("Array Queue: ")
	for _, curq := range *q {
		fmt.Print(curq, " ")
	}

	fmt.Println()
}
