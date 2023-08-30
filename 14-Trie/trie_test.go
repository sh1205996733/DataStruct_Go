package trie

import (
	"DataStruct_Go/utils"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Add("cat", 1)
	trie.Add("dog", 2)
	trie.Add("catalog", 3)
	trie.Add("cast", 4)
	trie.Add("小码哥", 5)
	utils.Asserts(trie.Size() == 5)
	utils.Asserts(trie.StartsWith("do"))
	utils.Asserts(trie.StartsWith("c"))
	utils.Asserts(trie.StartsWith("ca"))
	utils.Asserts(trie.StartsWith("cat"))
	utils.Asserts(trie.StartsWith("cata"))
	utils.Asserts(!trie.StartsWith("hehe"))
	utils.Asserts(trie.Get("小码哥") == 5)
	utils.Asserts(trie.Remove("cat") == 1)
	utils.Asserts(trie.Remove("catalog") == 3)
	utils.Asserts(trie.Remove("cast") == 4)
	utils.Asserts(trie.Size() == 2)
	utils.Asserts(trie.StartsWith("小"))
	utils.Asserts(trie.StartsWith("do"))
	utils.Asserts(!trie.StartsWith("c"))
}
