package p1566

func containsPattern(arr []int, m int, k int) bool {
	for i := 0; i+m*k <= len(arr); i++ {
		var ok = true
		for j := 0; j < k && ok; j++ {
			for u := 0; u < m && ok; u++ {
				if arr[i+j*m+u] != arr[i+u] {
					ok = false
				}
			}
		}
		if ok {
			return true
		}
	}
	return false
}
