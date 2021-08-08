package p1963

func minSwaps(s string) int {
	// 只需要把当前的]替换到最后面,
	// ] 越到后面越有利
	n := len(s)
	var level int
	var swaps int
	for i, j := 0, n-1; i < j; i++ {
		if s[i] == '[' {
			level++
		} else {
			level--
		}
		if level < 0 {
			swaps++
			// unbalance
			for s[j] != '[' {
				// should be safe
				j--
			}
			// move one more step
			j--
			level = 1
		}
	}
	return swaps
}
