package strint

func ReverseInt(num int) int {
	// negative number cannot be reversed
	if num < 0 {
		return 0
	}
	// single digits should return itself
	if num < 10 {
		return num
	}
	// numbers end with 0 cannot be reversed
	if num % 10 == 0 {
		return 0
	}

	reversedNum := 0
	for num > 0 {
		reversedNum = reversedNum * 10 + num % 10
		num /= 10
	}

	return reversedNum
}