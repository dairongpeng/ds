package stack

type Stack[T any] interface {
	Push(v T)
	Pop() (T, bool)
	Top() (T, bool)
	Size() int
	IsEmpty() bool
	Print()
}
