package array

/*	MaxArea2Ptr solves LeetCode 11 (https://leetcode.com/problems/container-with-most-water/)

	Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
	n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0).
	Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.
*/
func MaxArea2Ptr(height []int) int {
	res := 0
	leftIndex := 0
	rightIndex := len(height) - 1

	for leftIndex < rightIndex {
		res = max(res, min(height[leftIndex], height[rightIndex])*(rightIndex-leftIndex))

		if height[leftIndex] < height[rightIndex] {
			leftIndex++
			continue
		}
		rightIndex--
	}
	return res
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
