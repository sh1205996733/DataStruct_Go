package avl

import (
	"DataStruct_Go/print/printer"
	"fmt"
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
	//data := []Integer {
	//	67, 52, 92, 96, 53, 95, 13, 63, 34, 82, 76, 54, 9,
	//}
	//
	//avl := NewAVLTree()
	//for i := 0; i < len(data); i++ {
	//	avl.Add(data[i])
	//}
	//
	//printer.Println(avl.RootNode())

	avl := NewAVLTree()
	avl.Comparator = func(e1, e2 interface{}) int {
		return (e1.(Person)).id - (e2.(Person)).id
	}
	for i := 0; i < 10; i++ {
		avl.Add(Person{100 + i, fmt.Sprintf("test%d", i)})
	}
	printer.Println(avl.RootNode())
	fmt.Println("---------------------------------------------")
	index := Person{100 + 7, "test7"}
	avl.Remove(index)
	printer.Println(avl.RootNode())
	avl.InorderTraversal()
	fmt.Println()

}
