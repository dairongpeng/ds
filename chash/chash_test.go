package chash

import (
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	nodes := []string{"NodeA", "NodeB", "NodeC"}
	virtualNodes := 1000

	hashRing := NewRing(virtualNodes, nodes)

	// 假设有一些键需要映射到节点
	keys := []string{"KeyA", "KeyB", "KeyC", "KeyD", "KeyE", "apple", "banana", "cherry", "durian"}

	for _, key := range keys {
		node := hashRing.GetNode(key)
		fmt.Printf("Key %s maps to node %s\n", key, node)
	}
}
