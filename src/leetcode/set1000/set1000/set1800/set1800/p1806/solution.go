package p1806

func reinitializePermutation(n int) int {
	// arr[i] = arr[i/2] if i is even
	// arr[i] = arr[n/2 + (i - 1) / 2) is i is odd
	// n is even
	h := n / 2
	//0, 1, 2, 3, 4, 5, 6, 7
	// arr[i] = arr[h], i has to be odd, i is odd
	// so i has to 1
	var res int = 1
	cur := 1
	for cur != h {
		if cur > h {
			// some(i) n/2 + (i - 1)/2 = cur
			res++
			cur = (cur-n/2)*2 + 1
		} else {
			res++
			cur *= 2
		}
	}

	return res
}
