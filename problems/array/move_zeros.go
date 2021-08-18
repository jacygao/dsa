package array

func moveZeroes(nums []int) {
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
