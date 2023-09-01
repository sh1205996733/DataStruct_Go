package skiplist

import (
	"DataStruct_Go/utils"
	"testing"
)

func TestSkipList(t *testing.T) {
	//fmt.Println(randomLevel())
	list := NewSkipList()
	testSkipList(list, 30, 10)
}

func TestSkipListTime(t *testing.T) {
	count := 100_0000
	delta := 10
	utils.Times("SkipList", func() {
		list := NewSkipList()
		testSkipList(list, count, delta)
	})
	utils.Times("TreeMap", func() {
		list := map[int]int{}
		testMap(list, count, delta)
	})
}

func testSkipList(list *skiplist, count, delta int) {
	for i := 0; i < count; i++ {
		list.Put(i, i+delta)
	}
	for i := 0; i < count; i++ {
		utils.Asserts(list.Get(i) == i+delta)
	}
	utils.Asserts(list.Size() == count)
	for i := 0; i < count; i++ {
		list.Remove(i)
		//utils.Asserts(list.remove(i) == i+delta)
	}
	utils.Asserts(list.Size() == 0)
}

func testMap(m map[int]int, count, delta int) {
	for i := 0; i < count; i++ {
		m[i] = i + delta
	}
	for i := 0; i < count; i++ {
		utils.Asserts(m[i] == i+delta)
	}
	utils.Asserts(len(m) == count)
	for i := 0; i < count; i++ {
		//utils.Asserts(m.remove(i) == i+delta)
		delete(m, i)
	}
	utils.Asserts(len(m) == 0)
}
