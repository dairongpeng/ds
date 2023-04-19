package list

type List[T any] interface {
	Add(T)
	Remove() (T, bool)
	Get(index int) (T, bool)
	Print()
}
