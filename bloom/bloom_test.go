package bloom

import (
	"fmt"
	"testing"
)

// === RUN   TestNewFilter
// Element 'apple' may exist in the Bloom filter
// Element 'banana' may exist in the Bloom filter
// Element 'cherry' may exist in the Bloom filter
// Element 'durian' does not exist in the Bloom filter
// --- PASS: TestNewFilter (0.00s)
// PASS
//
// Process finished with the exit code 0
func TestNewFilter(t *testing.T) {
	// 100万数据量：1000000 预期失误率控制在万分之一以下为：0.0001
	n := 1000000
	p := 0.0001
	bf := NewFilter(n, p) // 真实失误率：0.0001007858789369095 k=14 len(arr)=599067

	elements := []string{"apple", "banana", "cherry"}
	// 将元素添加到布隆过滤器中
	for _, element := range elements {
		bf.Add(element)
	}

	// 检查元素是否可能存在于布隆过滤器中
	// 万分之一的概率，元素不在bloom中，被判定存在bloom中！
	testElements := []string{"apple", "banana", "cherry", "durian"}
	for _, element := range testElements {
		if bf.Contains(element) {
			fmt.Printf("Element '%s' may exist in the Bloom filter\n", element)
		} else {
			fmt.Printf("Element '%s' does not exist in the Bloom filter\n", element)
		}
	}
}
