package p2397

func isStrictlyPalindromic(n int) bool {
	x := 2
	for x <= n-2 && x <= n/x {
		var arr []int
		num := n
		for num > 0 {
			arr = append(arr, num%x)
			num /= x
		}

		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			if arr[i] != arr[j] {
				return false
			}
		}
		x++
	}

	a := n % x
	b := n / x

	return a == b
}
