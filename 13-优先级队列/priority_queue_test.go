package priority_queue

import (
	"errors"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	queue := NewPriorityQueue()
	queue.EnQueue(Person{"Jack", 2})
	queue.EnQueue(Person{"Rose", 10})
	queue.EnQueue(Person{"Jake", 5})
	queue.EnQueue(Person{"James", 15})

	for !queue.IsEmpty() {
		fmt.Println(queue.DeQueue())
	}
}

type Person struct {
	name string
	age  int
}

func (p Person) CompareTo(v interface{}) (int, error) {
	if p1, ok := v.(Person); ok {
		return p.age - p1.age, nil
	}
	return 0, errors.New("类型不匹配")
}
