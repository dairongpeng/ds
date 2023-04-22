package sort

type DSSort[T any] interface {
	Sort(arr []T, cmp func(item1, item2 any) int)
}
