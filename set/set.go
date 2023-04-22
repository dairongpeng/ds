package dsset

type DSSet[T any] interface {
	Add(value T)
	Remove(value T)
	Contains(value T) bool
}
