package array

import "sort"

// SearchInsertPosition using Binary Search
func SearchInsertPosition(nums []int, target int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		}
		return 1
	}
	mid := len(nums) / 2
	half := nums[mid]
	if half == target {
		return mid
	}
	if half > target {
		// left
		return SearchInsertPosition(nums[:mid], target)
	}
	// right
	return SearchInsertPosition(nums[mid+1:], target) + mid + 1
}

// SearchInsertPositionSort uses sort package.
// This solution requires less code but about 30% less performant.
func SearchInsertPositionSort(nums []int, target int) int {
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	return i
}
