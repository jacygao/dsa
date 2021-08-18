package array

/* MoveZeroes solves Leetcode 283:

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:

Input: nums = [0]
Output: [0]

*/
func MoveZeroesSol1(nums []int) {
	l, r := 0, 1
	for r < len(nums) {
		if nums[l] != 0 {
			l++
			r++
			continue
		}
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r++
			continue
		}
		r++
	}
}

func MoveZeroesSol2(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}
