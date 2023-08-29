package set

import (
	"DataStruct_Go/utils"
	"fmt"
)

// hashMapSet hashmap集合
type hashMapSet struct {
	m map[any]any
}

func (h *hashMapSet) Size() int {
	return len(h.m)
}

func (h *hashMapSet) IsEmpty() bool {
	return len(h.m) == 0
}

func (h *hashMapSet) Clear() {
	for k, _ := range h.m {
		delete(h.m, k)
	}
}

func (h *hashMapSet) Contains(element any) bool {
	_, ok := h.m[element]
	return ok
}

func (h *hashMapSet) Add(element any) {
	h.m[element] = element
}

func (h *hashMapSet) Remove(element any) {
	delete(h.m, element)
}

func (h *hashMapSet) Traversal(visitor ...utils.Visitor) {
	for _, v := range h.m {
		fmt.Println(v)
	}
}

func NewHashMapSet() utils.Set {
	return &hashMapSet{
		m: make(map[any]any),
	}
}
