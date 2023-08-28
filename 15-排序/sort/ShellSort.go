package sort

import (
	"math"
)

// ShellSort 希尔排序 [不稳定排序、原地算法] 希尔排序也被称为递减增量排序
type ShellSort struct {
	BaseSort
}

/**
 * 希尔排序把序列看作是一个矩阵，分成 𝑚 列，逐列进行排序
 * 	✓ 𝑚 从某个整数逐渐减为1
 *	✓ 当 𝑚 为1时，整个序列将完全有序
 * 矩阵的列数取决于步长序列（step sequence）
 * 	✓ 比如，如果步长序列为{1,5,19,41,109,...}，就代表依次分成109列、41列、19列、5列、1列进行排序
 * 	✓ 不同的步长序列，执行效率也不同
 */

func (s *ShellSort) Sort() {
	stepSequence := sedgewickStepSequence(len(s.Array))
	for _, step := range stepSequence {
		s.shellSort(step)
	}
}

// 因此希尔排序底层一般使用插入排序对每一列进行排序，也很多资料认为希尔排序是插入排序的改进版
func (s ShellSort) shellSort(step int) {
	// col : 第几列，column的简称
	for col := 0; col < step; col++ {
		// col、col+step、col+2*step、col+3*step
		for begin := col + step; begin < len(s.Array); begin += step {
			//每一列进行排序(插入排序)
			cur := begin
			//for cur > col && s.cmp(cur, cur-step) < 0 {
			//	s.swap(cur, cur-step)
			//	cur -= step
			//}
			v := s.Array[cur]
			for cur > col && s.cmpVal(v, s.Array[cur-1]) < 0 {
				s.Array[cur] = s.Array[cur-step]
				cur -= step
			}
			s.Array[cur] = v
		}
	}
}

// 获取步长序列 16 8 4 2 1
func shellStepSequence(size int) []int {
	stepSequence := make([]int, 0)
	for step := size >> 1; step > 0; step >>= 1 {
		stepSequence = append(stepSequence, step)
	}
	return stepSequence
}

// 获取步长序列
func sedgewickStepSequence(size int) []int {
	stepSequence := make([]int, 0)
	k, step := 0, 0
	for {
		if k%2 == 0 {
			pow := int(math.Pow(2.0, float64(k>>1)))
			step = 1 + 9*(pow*pow-pow)
		} else {
			pow1 := int(math.Pow(2.0, float64((k-1)>>1)))
			pow2 := int(math.Pow(2.0, float64((k+1)>>1)))
			step = 1 + 8*pow1*pow2 - 6*pow2
		}
		if step >= size {
			break
		}
		stepSequence = append([]int{step}, stepSequence...)
		k++
	}
	//fmt.Println("stepSequence = ", stepSequence)
	return stepSequence
}
