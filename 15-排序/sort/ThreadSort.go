package sort

import (
	"fmt"
	"time"
)

// ThreadSort 史上“最强”排序 – 休眠排序
type ThreadSort struct {
	Value int
}

func NewThreadSort(value int) *ThreadSort {
	return &ThreadSort{Value: value}
}
func (t *ThreadSort) Start() {
	time.Sleep(time.Duration(t.Value) * time.Second)
	fmt.Println(t.Value)
}
func main() {
	array := []int{5, 4, 3, 2, 1}
	for i := 0; i < len(array); i++ {
		NewThreadSort(array[i]).Start()
	}
}
