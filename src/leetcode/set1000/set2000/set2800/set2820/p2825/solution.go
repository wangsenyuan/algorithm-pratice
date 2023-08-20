package p2825

func canMakeSubsequence(str1 string, str2 string) bool {
	if len(str2) > len(str1) {
		return false
	}

	n := len(str1)
	m := len(str2)

	// 需要找出str1中的某些下标，将它们变成下一个字符
	// 然后判断str2是否是str1的子序列
	// 从i开始，把那些需要变化的记录下来，即可
	for i, j := 0, 0; i < n && j < m; i++ {
		if str1[i] == str2[j] {
			j++
		} else if next(str1[i]) == str2[j] {
			j++
		}
		if j == m {
			return true
		}
	}

	return false
}

func next(x byte) byte {
	if x < 'z' {
		return byte(int(x-'a') + 'a' + 1)
	}
	return 'a'
}
