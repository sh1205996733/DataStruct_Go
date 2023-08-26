package array

import (
	"fmt"
	"testing"
)

// 动态数组
func TestArrayList(t *testing.T) {
	list := NewArrayList()
	list.Add(10)
	list.Add(Person{10, "Jack"})
	list.Add(22)

	list.IndexOf(Person{10, "Jack"})
	fmt.Println(list)

	//persons := NewArrayList()
	//persons.Add(Person{10, "Jack"})
	//persons.Add(nil)
	//persons.Add(Person{15, "Rose"})
	//persons.Add(nil)
	//persons.Add(Person{12, "James"})
	//persons.Add(nil)
	//fmt.Println(persons)
	//fmt.Println(persons.IndexOf(nil))
}

func TestArrayList2(t *testing.T) {
	persons := NewArrayList()
	persons.Add(Person{10, "Jack"})
	persons.Add(Person{12, "James"})
	persons.Add(Person{15, "Rose"})
	persons.Clear()
	persons.Add(Person{22, "abc"})

	fmt.Println(persons)
	ints := NewArrayList()
	ints.Add(10)
	ints.Add(10)
	ints.Add(22)
	ints.Add(33)
	fmt.Println(ints)
}

type Person struct {
	age  int
	name string
}

func (p *Person) String() string {
	return fmt.Sprintf("Person [age=%d, name=%s]", p.age, p.name)
}

func (p *Person) Equals(obj any) bool {
	if obj == nil {
		return false
	}
	if p1, ok := obj.(*Person); ok {
		return p.age == p1.age
	}
	if p1, ok := obj.(Person); ok {
		return p.age == p1.age
	}
	return false
}
