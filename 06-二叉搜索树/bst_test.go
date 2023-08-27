package binarytree

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

func TestBST(t *testing.T) {
	data := []int{
		7, 4, 9, 2, 5, 8, 11, 3, 12, 1,
	}

	bst := new(BST)
	//bst.Comparator = func(e1,e2 interface{}) int {
	//	return e1.(int) - e2.(int)
	//}

	for i := 0; i < len(data); i++ {
		bst.Add(data[i])
	}

	printer.Println(bst.RootNode())

	bst.Remove(7)

	printer.Println(bst.RootNode())

}
func TestBST2(t *testing.T) {

	bst := new(BST)
	for i := 0; i < 10; i++ {
		bst.Add(Person{100 + i, fmt.Sprintf("test%d", i)})
	}

	printer.Println(bst.RootNode())

	//index := Person{100+7,"test7"}
	bst.Remove(107)

	printer.Println(bst.RootNode())
	bst.InorderTraversal()
	fmt.Println()

}

func TestBST_InorderTraversal(t *testing.T) {
	data := []int{
		7, 4, 9, 2, 5, 8, 11, 3, 12, 1,
	}

	bst := new(BST)
	for i := 0; i < len(data); i++ {
		bst.Add(data[i])
	}

	printer.Println(bst.RootNode())

	//Visitor := new(binarytree.Visitor)
	//Visitor.Visit = func(value interface{}) bool {
	//	fmt.Println(value," ",value == 2)
	//	return value == 2
	//}
	bst.PreorderTraversal(struct {
		Stop  bool
		Visit func(interface{}) bool
	}{Stop: false, Visit: func(value interface{}) bool {
		fmt.Println(value, " ", value == 2)
		return value == 2
	}})
	//bst.PreorderTraversal(Visitor)
	fmt.Println()
}
