// Package dstrie 前缀树又叫字典树, 是一颗多叉树
package dstrie

// Node 前缀树的节点
type Node struct {
	Pass    int
	End     int
	Childes []*Node // 也可以实现成map，达到容纳更多字符的目的
}

type Trie struct {
	Head *Node
}

func NewTrie() (root *Trie) {
	head := &Node{
		Pass: 0,
		End:  0,
		// 默认保存26个英文字符a~z
		// 0 a
		// 1 b
		// .. ..
		// 25 z
		// Childes[i] == nil 表示i方向的路径不存在
		Childes: make([]*Node, 26),
	}

	trie := &Trie{
		Head: head,
	}
	return trie
}

// Insert 往该前缀树中添加字符串
func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	// 字符串转字符数组,每个元素是字符的ascii码
	chs := []byte(word)
	node := t.Head
	// 头结点的pass首先++
	node.Pass++
	// 路径的下标
	var path int
	// 从左往右遍历字符
	for i := 0; i < len(chs); i++ {
		// 当前字符减去'a'的ascii码得到需要添加的下个节点下标。即当前字符去往的路径
		path = int(chs[i] - 'a')
		// 当前方向上没有建立节点，即一开始不存在这条路，新开辟
		if node.Childes[path] == nil {
			node.Childes[path] = &Node{
				Pass:    0,
				End:     0,
				Childes: make([]*Node, 26),
			}
		}
		// 引用指向当前来到的节点
		node = node.Childes[path]
		// 当前节点的pass++
		node.Pass++
	}
	// 当新加的字符串所有字符处理结束，最后引用指向的当前节点就是该字符串的结尾节点，end++
	node.End++
}

// Search 在该前缀树中查找word这个单词之前加入过几次
func (t *Trie) Search(word string) int {
	if len(word) == 0 {
		return 0
	}
	chs := []byte(word)
	node := t.Head
	path := 0
	for i := 0; i < len(chs); i++ {
		path = int(chs[i] - 'a')
		// 寻找该字符串的路径中如果提前找不到path，就是未加入过，0次
		if node.Childes[path] == nil {
			return 0
		}
		node = node.Childes[path]
	}
	// 如果顺利把word字符串在前缀树中走完路径，那么此时的node对应的end值就是当前word在该前缀树中添加了几次
	return node.End
}

// Delete 删除该前缀树的某个字符串
func (t *Trie) Delete(word string) {
	// 首先要查一下该字符串是否加入过, 该字符串在我们的前缀树中
	if t.Search(word) != 0 {
		// 沿途pass--
		chs := []byte(word)
		node := t.Head
		node.Pass--
		path := 0
		for i := 0; i < len(chs); i++ {
			path = int(chs[i] - 'a')
			// 在寻找的过程中，pass为0，提前可以得知在本次删除之后，该节点以下的路径不再需要，可以直接删除。
			// 那么该节点之下下个方向的节点引用置为空，GC会回收内存
			node.Childes[path].Pass--
			if node.Childes[path].Pass == 0 {
				node.Childes[path] = nil
				return
			}
			node = node.Childes[path]
		}
		// 最后end--
		node.End--
	}
}

// PrefixNumber 所有加入的字符串中，有几个是以pre这个字符串作为前缀的
func (t *Trie) PrefixNumber(pre string) int {
	if len(pre) == 0 {
		return 0
	}

	chs := []byte(pre)
	node := t.Head
	path := 0
	for i := 0; i < len(chs); i++ {
		path = int(chs[i] - 'a')
		// pre走不到最后，就没有以pre作为前缀的字符串存在
		if node.Childes[path] == nil {
			return 0
		}
		node = node.Childes[path]
	}
	// 顺利走到最后，返回的pass就是有多少个字符串以当前pre为前缀的
	return node.Pass
}
