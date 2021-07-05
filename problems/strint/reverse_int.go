package strint

import "math"

func ReverseInt(num int) int {
	// single digits should return itself
	if num < 10 && num > -10 {
		return num
	}

	reversedNum := 0
	for num != 0 {
		reversedNum = reversedNum*10 + num%10
		num /= 10
	}

	if reversedNum > math.MaxInt32 || reversedNum < math.MinInt32 {
		return 0
	}
	return reversedNum
}
