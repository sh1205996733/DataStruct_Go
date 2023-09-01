package bloom

import (
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(1_00_0000, 0.25)
	for i := 1; i <= 1_0000; i++ {
		bf.Put(i)
	}

	var count = 0
	for i := 1_0001; i <= 20000; i++ {
		if bf.Contains(i) {
			count++
		}
	}
	fmt.Println(count)
	// 数组
	//		String[] urls = {};
	//		BloomFilter<String> bf = new BloomFilter<String>(10_0000_0000, 0.01);
	//		/*
	//		for (String url : urls) {
	//			if (bf.contains(url)) continue;
	//			// 爬这个url
	//			// ......
	//
	//			// 放进BloomFilter中
	//			bf.put(url);
	//		}*/
	//
	//		for (String url : urls) {
	//			if (bf.put(url) == false) continue;
	//			// 爬这个url
	//			// ......
	//		}
}
