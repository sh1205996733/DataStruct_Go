package linked_list

import (
	"DataStruct_Go/utils"
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	//list := new(SingleLinkedList)
	//list := new(SingleLinkedListV2)
	//list.head = NewNode(nil, nil, nil)
	//list := new(DoubleLinkedList)
	list := new(SingleCircleLinkedList)
	list.Add(11)
	list.Add(22)
	list.Add(33)
	list.Add(44)

	list.AddAtIndex(0, 55)           // [55, 11, 22, 33, 44]
	list.AddAtIndex(2, 66)           // [55, 11, 66, 22, 33, 44]
	list.AddAtIndex(list.Size(), 77) // [55, 11, 66, 22, 33, 44, 77]

	list.Remove(0)               // [11, 66, 22, 33, 44, 77]
	list.Remove(2)               // [11, 66, 33, 44, 77]
	list.Remove(list.Size() - 1) // [11, 66, 33, 44]

	utils.Asserts(list.IndexOf(44) == 3)
	utils.Asserts(list.IndexOf(22) == utils.ELEMENT_NOT_FOUND)
	utils.Asserts(list.Contains(33))
	utils.Asserts(list.Get(0) == 11)
	utils.Asserts(list.Get(1) == 66)
	utils.Asserts(list.Get(list.Size()-1) == 44)

	fmt.Println(list)
}

func TestCircleLinkedList(t *testing.T) {
	list := new(DoubleCircleLinkedList)
	for i := 1; i <= 8; i++ {
		list.Add(i)
	}

	// 指向头结点（指向1）
	list.Reset()
	fmt.Println(list)
	//for !list.IsEmpty() {
	//	list.Next()
	//	list.Next()
	//	fmt.Println(list.RemoveCurrent())
	//}
}
