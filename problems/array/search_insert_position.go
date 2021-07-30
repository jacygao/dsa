package array

import "sort"

/* SearchInsertPositionBinarySearch solves the following problem.

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

#Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2

#Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1

#Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4

#Example 4:

Input: nums = [1,3,5,6], target = 0
Output: 0

#Example 5:

Input: nums = [1], target = 0
Output: 0
*/
func SearchInsertPositionBinarySearch(nums []int, target int) int {
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
		return SearchInsertPositionBinarySearch(nums[:mid], target)
	}
	// right
	return SearchInsertPositionBinarySearch(nums[mid+1:], target) + mid + 1
}

// SearchInsertPositionSort uses sort package.
// This solution requires less code but about 30% less performant.
func SearchInsertPositionSort(nums []int, target int) int {
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	return i
}
