package gomap

import "testing"

func TestNewTreeMap(t *testing.T) {
	m := NewTreeMap()
	m.Put(1, "A")
	m.Put(2, "B")
	m.Put(3, "C")
	m.Put(4, "D")
	m.Put(4, "D")
	m.Remove(3)
	m.Traversal()
}
