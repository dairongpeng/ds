package arraystack

import "fmt"

type Stack[T any] []T

// New 初始化一个栈
func New[T any](values ...T) *Stack[T] {
	var s = &Stack[T]{}
	if len(values) > 0 {
		for _, v := range values {
			s.Push(v)
		}
	}
	return s
}

// Push 将元素v压入栈顶
func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Pop 弹出栈顶元素，返回弹出的元素和一个bool值，如果栈为空则返回一个零值和false。
func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	index := len(*s) - 1
	val := (*s)[index]
	*s = (*s)[:index]
	return val, true
}

// Top 返回栈顶元素，但不将其弹出，如果栈为空则返回一个零值和false。
func (s *Stack[T]) Top() (T, bool) {
	if len(*s) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	index := len(*s) - 1
	val := (*s)[index]
	return val, true
}

// Size 返回栈的元素个数。
func (s *Stack[T]) Size() int {
	return len(*s)
}

// IsEmpty 判断栈是否为空。
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Print 打印栈
func (s *Stack[T]) Print() {
	fmt.Println("Array Stack: ")
	for _, curs := range *s {
		fmt.Print(curs, " ")
	}

	fmt.Println()
}
