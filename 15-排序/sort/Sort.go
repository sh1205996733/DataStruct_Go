package sort

import (
	"fmt"
	"reflect"
	"time"
)

type BaseSort struct {
	Name      string
	Array     []int
	CmpCount  int
	SwapCount int
	Time      int64
}

func DoSort(array []int, sort Sort) {
	if array == nil || len(array) < 2 {
		return
	}
	f := reflect.ValueOf(sort).Elem()
	f.FieldByName("Name").SetString(f.Type().Name())
	f.FieldByName("Array").Set(reflect.ValueOf(array))
	begin := time.Now()
	sort.Sort()
	fmt.Println("["+f.Type().Name()+"]	排序后:", array)
	f.FieldByName("Time").SetInt(time.Since(begin).Milliseconds())
}

type Sort interface {
	Sort()
}

/**
 * 返回值等于0，代表 array[i1] == array[i2]
 * 返回值小于0，代表 array[i1] < array[i2]
 * 返回值大于0，代表 array[i1] > array[i2]
 */
func (i *BaseSort) cmp(i1, i2 int) int {
	i.CmpCount++
	return i.Array[i1] - i.Array[i2]
}
func (i *BaseSort) cmpVal(v1, v2 int) int {
	i.CmpCount++
	return v1 - v2
}

func (i *BaseSort) swap(i1, i2 int) {
	i.SwapCount++
	i.Array[i1], i.Array[i2] = i.Array[i2], i.Array[i1]
}

func SortSlice(array []Sort) {
	for end := len(array) - 1; end > 0; end-- { //end从最后一个元素length - 1开始遍历到1
		for begin := 1; begin <= end; begin++ { //begin从第二个元素arr[1]开始遍历到end
			if CompareTo(array[begin], array[begin-1]) < 0 { //比较arr[i]和arr[i-1]的值，若arr[i-1]大，就交换
				array[begin], array[begin-1] = array[begin-1], array[begin]
			}
		}
	}
}
func CompareTo(i, j interface{}) int64 {
	oi := reflect.ValueOf(i).Elem()
	oj := reflect.ValueOf(j).Elem()

	result := oi.FieldByName("Time").Int() - oj.FieldByName("Time").Int()
	if result != 0 {
		return result
	}

	result = oi.FieldByName("CmpCount").Int() - oj.FieldByName("CmpCount").Int()
	if result != 0 {
		return result
	}

	return oi.FieldByName("SwapCount").Int() - oj.FieldByName("SwapCount").Int()
}
func (i *BaseSort) String() string {
	timeStr := fmt.Sprintf("耗时：%ds(%dms)", i.Time/1000, i.Time)
	compareCountStr := "比较：" + numberString(i.CmpCount)
	swapCountStr := "交换：" + numberString(i.SwapCount)
	stableStr := fmt.Sprintf("稳定性：%v", i.isStable())
	inPlaceStr := fmt.Sprintf("原地算法：%v", i.isInPlace())
	return "【" + i.Name + "】\n" +
		stableStr + " \t" +
		inPlaceStr + " \t" +
		timeStr + " \t" +
		compareCountStr + "\t " +
		swapCountStr + "\n" +
		"------------------------------------------------------------------"
}

func numberString(number int) string {
	if number < 10000 {
		return fmt.Sprintf("%d", number)
	}
	if number < 100000000 {
		return fmt.Sprintf("%d万", number/10000.0)
	}
	return fmt.Sprintf("%d亿", number/100000000.0)
}

func (i *BaseSort) isStable() bool {
	m := map[string]bool{
		"BubbleSort":    true,
		"SelectionSort": false,
		"InsertionSort": true,
		"MergeSort":     true,
		"QuickSort":     false,
		"ShellSort":     false,
		"HeapSort":      false,
		"BucketSort":    true,
		"RadixSort":     true,
		"CountingSort":  true,
	}
	return m[i.Name]
}

func (i *BaseSort) isInPlace() bool {
	m := map[string]bool{
		"BubbleSort":    true,
		"InsertionSort": true,
		"SelectionSort": true,
		"MergeSort":     false,
		"QuickSort":     true,
		"ShellSort":     true,
		"HeapSort":      true,
		"BucketSort":    false,
		"RadixSort":     false,
		"CountingSort":  false,
	}
	return m[i.Name]
}
