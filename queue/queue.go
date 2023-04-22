package dsqueue

type DSQueue[T any] interface {
	Enqueue(value T)
	Dequeue() (T, bool)
	Front() (T, bool)
	Size() int
	IsEmpty() bool
	Print()
}
