package rbtree

import (
	"fmt"
	"leetcode/print/printer"
	"testing"
)

type Person struct {
	id   int
	name string
}

func (t Person) CompareTo(e interface{}) (int, error) {
	t1, ok := e.(Person)
	var err error
	if !ok {
		err = fmt.Errorf("%s", "case conversion failed")
	}
	return int(t.id - t1.id), err
}
func (t Person) Print() int {
	return t.id
}

type Integer int

func (t Integer) CompareTo(e interface{}) (int, error) {
	t1, ok := e.(Integer)
	var err error
	if !ok {
		err = fmt.Errorf("%s", "case conversion failed")
	}
	return int(t - t1), err
}

func TestAVl(t *testing.T) {
	data := []Integer{
		67, 52, 92, 96, 53, 95, 13, 63, 34, 82, 76, 54, 9, 68, 39,
	}

	rb := NewRBTree()
	for i := 0; i < len(data); i++ {
		rb.Add(data[i])
	}

	//printer.Println(avl.RootNode())
	printer.Println(rb.RootNode())
	fmt.Println("---------------------------------------------")
	//ndex := Person{100+7,"test7"}
	rb.Remove(Integer(34))
	printer.Println(rb.RootNode())
	//rb.InorderTraversal()
	//fmt.Println()

}
