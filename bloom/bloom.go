package bloom

import (
	"hash/fnv"
	"math"
)

// Filter 布隆过滤器
type Filter struct {
	// 位图
	arr []uint32
	// hash函数列表
	hashFs []func(data []byte) uint32
	// 失误率
	p float64
}

// NewFilter 初始化一个布隆过滤器; 指定预期样本量n，指定预期失误率p
//
//  1. 根据n和p可以计算出布隆位数（bit）m, 向上取整
//  2. 根据M和N，可计算出需要多少个hash函数，记为k，向上取整
//  3. 根据1，2信息计算真实失误率为q
//  4. 布隆过滤器不支持删除
func NewFilter(n int, p float64) *Filter {
	// 位图空间
	m := -float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)
	// 哈希函数个数
	k := (m / float64(n)) * math.Log(2)

	upM := int(math.Ceil(m))
	upK := int(math.Ceil(k))

	var size int
	block := upM / 32
	if int32(upM)%32 == 0 {
		size = block
	} else {
		size = block + 1
	}

	// 加工upK个hash函数，作为对比的指纹
	hashFs := make([]func(data []byte) uint32, upK)
	for i := 0; i < upK; i++ {
		hashFs[i] = generateHashFunc(i)
	}

	// 真实失误率
	q := math.Pow(1-math.Exp(-float64(upK)*float64(n)/float64(upM)), float64(upK))
	f := &Filter{
		arr:    make([]uint32, size),
		hashFs: hashFs,
		p:      q,
	}

	return f
}

// 根据给定的种子生成哈希函数, 每个hash函数是独立且均匀分布的
func generateHashFunc(seed int) func(data []byte) uint32 {
	return func(data []byte) uint32 {
		hash := fnv.New32a()
		_, _ = hash.Write(data)
		return hash.Sum32() + uint32(seed)
	}
}

// P 获取布隆过滤器的真实失误率
func (bf *Filter) P() float64 {
	return bf.p
}

// Add 往布隆过滤器中添加一个元素
func (bf *Filter) Add(element string) {
	for _, hf := range bf.hashFs {
		// 每个指纹描摹
		arrIx := hf([]byte(element)) % uint32(len(bf.arr))
		blockIx := hf([]byte(element)) % 32

		bf.arr[arrIx] = bf.arr[arrIx] | (1 << blockIx)
	}
}

// Contains 检查元素是否可能存在于布隆过滤器中
func (bf *Filter) Contains(element string) bool {
	for _, hf := range bf.hashFs {
		// 每个指纹对比
		arrIx := hf([]byte(element)) % uint32(len(bf.arr))
		blockIx := hf([]byte(element)) % 32

		ix := bf.arr[arrIx] >> blockIx
		if ix&1 == 0 {
			return false
		}
	}
	return true
}
