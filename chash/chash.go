package chash

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Ring struct {
	virtualNodes int               // 虚拟节点的数量
	nodes        []string          // 物理节点标识
	ring         map[uint32]string // 哈希环，存储虚拟节点的哈希值与真实节点的映射关系
}

func NewRing(virtualNodes int, nodes []string) *Ring {
	r := &Ring{
		virtualNodes: virtualNodes,
		nodes:        nodes,
		ring:         make(map[uint32]string),
	}

	r.generateRing()
	return r
}

// generateRing generates the hash ring
func (h *Ring) generateRing() {
	for _, node := range h.nodes {
		for i := 0; i < h.virtualNodes; i++ {
			virtualKey := fmt.Sprintf("%s-%d", node, i)
			hash := crc32.ChecksumIEEE([]byte(virtualKey))
			h.ring[hash] = node
		}
	}
}

// GetNode returns the node responsible for the given key
func (h *Ring) GetNode(key string) string {
	hash := crc32.ChecksumIEEE([]byte(key))

	// 查找离hash最近的节点
	var hashKeys []uint32
	for k := range h.ring {
		hashKeys = append(hashKeys, k)
	}
	sort.Slice(hashKeys, func(i, j int) bool { return hashKeys[i] < hashKeys[j] })

	// 可采用二分查找加速
	for _, k := range hashKeys {
		if hash <= k {
			return h.ring[k]
		}
	}

	// 若未找到节点，则返回第一个节点
	return h.ring[hashKeys[0]]
}
