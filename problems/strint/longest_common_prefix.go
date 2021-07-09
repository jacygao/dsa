package strint

func longestCommonPrefix(strs []string) string {
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
