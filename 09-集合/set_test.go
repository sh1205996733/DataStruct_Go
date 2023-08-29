package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	//listSet := NewListSet()
	//listSet.Add(10)
	//listSet.Add(11)
	//listSet.Add(11)
	//listSet.Add(12)
	//listSet.Add(10)
	//listSet.Traversal()

	//treeSet := NewTreeSet()
	//treeSet.Add(12)
	//treeSet.Add(10)
	//treeSet.Add(7)
	//treeSet.Add(11)
	//treeSet.Add(10)
	//treeSet.Add(11)
	//treeSet.Add(9)
	//treeSet.Traversal()

	hashSet := NewHashMapSet()
	hashSet.Add(12)
	hashSet.Add(10)
	hashSet.Add(7)
	hashSet.Add(11)
	hashSet.Add(10)
	hashSet.Add(11)
	hashSet.Add(9)
	hashSet.Remove(7)
	fmt.Println(hashSet.Contains(9))
	hashSet.Traversal()
}
