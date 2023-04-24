package sort

type DSSort[T any] interface {
	Sort(cmp func(item1, item2 any) int)
}
