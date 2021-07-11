package strint

/* longestCommonPrefix solves the following problem.

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/
func LongestCommonPrefix(strs []string) string {
	firstStr := strs[0]
	total := len(strs)
	if total == 1 {
		return firstStr
	}

	for i := 0; i < len(firstStr); i++ {
		char := firstStr[i]
		for j := 1; j < total; j++ {
			strLen := len(strs[j])
			if strLen == 0 {
				return ""
			}
			if strLen <= i {
				return firstStr[:i]
			}
			if char != strs[j][i] {
				return firstStr[:i]
			}
		}
	}

	return firstStr
}
