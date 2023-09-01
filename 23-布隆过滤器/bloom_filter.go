package bloom

import (
	"DataStruct_Go/utils"
	"math"
)

// 布隆过滤器
type bloomFilter struct {
	// 二进制向量的长度(一共有多少个二进制位)
	bitSize int
	// 二进制向量
	bits []int
	// 哈希函数的个数
	hashSize int
}

// NewBloomFilter
// n 数据规模
// p 误判率, 取值范围(0, 1)
func NewBloomFilter(n int, p float64) *bloomFilter {
	if n <= 0 || p <= 0 || p >= 1 {
		panic("wrong n or p")
	}
	ln2 := math.Log(2)
	// 求出二进制向量的长度
	bitSize := int(-(float64(n) * math.Log(p)) / (ln2 * ln2))
	// 求出哈希函数的个数
	hashSize := 1 //int(float64(bitSize) * ln2 / float64(n))
	// bits数组的长度
	bits := make([]int, (bitSize+IntSize-1)/IntSize)
	// 每一页显示100条数据, pageSize
	// 一共有999999条数据, n
	// 请问有多少页 pageCount = (n + pageSize - 1) / pageSize
	return &bloomFilter{
		bitSize:  bitSize,
		hashSize: hashSize,
		bits:     bits,
	}
}

const IntSize = 64

// Put 添加元素
func (b *bloomFilter) Put(value any) bool {
	nullCheck(value)
	// 利用value生成2个整数
	hash1 := utils.Hash(value)
	hash2 := hash1 >> 16
	var result = false
	for i := 1; i <= b.hashSize; i++ {
		index := b.index(i, hash1, hash2)
		// 设置index位置的二进位为1
		if b.Set(index) {
			result = true
		}
		//   101010101010010101
		// | 000000000000000100   1 << index
		//   101010111010010101
	}
	return result
}

// Contains 判断一个元素是否存在
func (b *bloomFilter) Contains(value any) bool {
	nullCheck(value)
	// 利用value生成2个整数
	hash1 := utils.Hash(value)
	hash2 := hash1 >> 16

	for i := 1; i <= b.hashSize; i++ {
		index := b.index(i, hash1, hash2)
		// 查询index位置的二进位是否为0
		if !b.Get(index) {
			return false
		}
	}
	return true
}

// Set 设置index位置的二进位为1
func (b *bloomFilter) Set(index int) bool {
	value := b.bits[index/IntSize]
	bitValue := 1 << (index % IntSize)
	b.bits[index/IntSize] = value | bitValue
	return (value & bitValue) == 0
}

// Get 查看index位置的二进位的值 true代表1, false代表0
func (b *bloomFilter) Get(index int) bool {
	value := b.bits[index/IntSize]
	return (value & (1 << (index % IntSize))) != 0
}

func (b *bloomFilter) index(index int, hash1, hash2 int) int {
	combinedHash := hash1 + (index * hash2)
	if combinedHash < 0 {
		combinedHash = ^combinedHash
	}
	// 生成一个二进位的索引
	return combinedHash % b.bitSize
}
func nullCheck(value any) {
	if value == nil {
		panic("Value must not be null.")
	}
}
