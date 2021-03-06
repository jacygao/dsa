package strint

import "math"

/* ReverseInt solves the following problem.

# Problem Statement

Reverse any 32-bit integer mathmatically without using string.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
Return 0 on integer overflow.
(LeetCode 7)

# Example
Input: 12345 => Output: 54321

Input: 0 => Output: 0

Input: 5 => Output: 5

Input: 3420 => Output: 243

Input: -13 => Output: -31

Input: 1356712359 => 0 (integer overflow)
*/
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
