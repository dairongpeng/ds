package pkg

// Comparator 比较器
// if item1 < item2时, 返回负数
// if item1 == item2时，返回0
// if item1 > item2时，返回正数
type Comparator[T any] func(item1, item2 T) int

// NumberComparator 数值类型的比较器
func NumberComparator[T int | int32 | int64 | float32 | float64](a, b T) int {
	if a-b < 0 {
		return -1
	} else if a-b == 0 {
		return 0
	} else {
		return 1
	}
}
