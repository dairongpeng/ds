package queue

type Queue[T any] interface {
	Enqueue(v T)
	Dequeue() (T, bool)
	Front() (T, bool)
	Size() int
	IsEmpty() bool
	Print()
}
