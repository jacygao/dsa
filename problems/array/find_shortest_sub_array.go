package array

/* FindShortestSubArray solves LeetCode problem 697.

Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.
Find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

Example 1:

Input: nums = [1,2,2,3,1]
Output: 2
Explanation:
The input array has a degree of 2 because both elements 1 and 2 appear twice.
Of the subarrays that have the same degree:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
The shortest length is 2. So return 2.

Example 2:

Input: nums = [1,2,2,3,1,4,2]
Output: 6
Explanation:
The degree is 3 because the element 2 is repeated 3 times.
So [2,2,3,1,4,2] is the shortest subarray, therefore returning 6.

Constraints:

- nums.length will be between 1 and 50,000.
- nums[i] will be an integer between 0 and 49,999.
*/

func FindShortestSubArray(nums []int) int {
	// the number of appearance of each integer
	mNum := make(map[int]int, len(nums))
	// the first appearance of each integer
	firstAppear := make(map[int]int)

	maxNum := 0
	maxLen := 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		// increase the num of appearance of the number
		mNum[num]++
		// if number never appeared, register it
		if _, ok := firstAppear[num]; !ok {
			firstAppear[num] = i
			continue
		}

		// if new max number appearance found, replace current one
		if mNum[num] > maxNum {
			maxNum = mNum[num]
			maxLen = i - firstAppear[num]
			continue
		}

		// if same max num found, compare the sub array length and take the smaller one
		if mNum[num] == maxNum {
			if maxLen > i-firstAppear[num] {
				maxLen = i - firstAppear[num]
			}
		}
	}
	return maxLen + 1
}
