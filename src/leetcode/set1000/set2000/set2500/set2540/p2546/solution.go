package p2546

func makeStringsEqual(s string, target string) bool {
	n := len(s)
	// s[i] = s[i] | s[j]
	// s[j] = s[i] ^ s[j]
	// 0 1 => 1 1
	// 0 0 => 0 0
	// 1 0 => 1 1
	// 1 1 => 1 0
	// 如果有一个1，那么可以将所有0变成1
	// 如果有两个1，可以将其中的一个1变成0
	if s == target {
		return true
	}
	// s != target
	zero := true
	for i := 0; i < n && zero; i++ {
		zero = target[i] == '0'
	}
	if zero {
		return false
	}
	zero = true
	for i := 0; i < n && zero; i++ {
		zero = s[i] == '0'
	}
	if zero {
		return false
	}
	// 不全为0
	return true
}
