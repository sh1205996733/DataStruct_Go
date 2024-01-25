package search

// target 应该所在位置 [left,right]
func lower_bound1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>2
		if nums[mid] < target {
			left = mid + 1 //[mid+1,right]
		} else {
			right = mid - 1 //[left,mid-1]
		}
	}
	return left
}

// target 应该所在位置 [left,right)
func lower_bound2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>2
		if nums[mid] < target {
			left = mid + 1 //[mid+1,right)
		} else {
			right = mid //[left,mid)
		}
	}
	return left //right
}

// target 应该所在位置 (left,right)
func lower_bound3(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)>>2
		if nums[mid] < target {
			left = mid // (mid,right)
		} else {
			right = mid // (left,mid)
		}
	}
	return right
}
