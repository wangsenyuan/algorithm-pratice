package p1888

func minFlips(s string) int {
	n := len(s)
	a, b := make([]int, 2), make([]int, 2)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			if s[i] == '0' {
				a[0]++
			} else {
				a[1]++
			}
		} else {
			if s[i] == '0' {
				b[0]++
			} else {
				b[1]++
			}
		}
	}
	ans := min(flips(a, b, '0'), flips(a, b, '1'))

	if n&1 == 0 {
		return ans
	}

	A, B := make([]int, 2), make([]int, 2)
	C, D := make([]int, 2), make([]int, 2)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			if s[i] == '0' {
				A[0]++
			} else {
				A[1]++
			}
		} else {
			if s[i] == '0' {
				B[0]++
			} else {
				B[1]++
			}
		}
		// 将i(包括i)前的字符都移到后端
		if (i+1)%2 == 0 {
			// 移动了偶数个字符，后半部分的奇偶性不改变，前部分的改变
			C[0] = a[0] - A[0] + B[0]
			C[1] = a[1] - A[1] + B[1]
			D[0] = b[0] - B[0] + A[0]
			D[1] = b[1] - B[1] + A[1]
		} else {
			// 移动了奇数个字符，后半部分的奇偶改变，前部分没有变
			C[0] = A[0] + b[0] - B[0]
			C[1] = A[1] + b[1] - B[1]
			D[0] = B[0] + a[0] - A[0]
			D[1] = B[1] + a[1] - A[1]
		}
		ans = min(ans, min(flips(C, D, '0'), flips(C, D, '1')))
	}

	return ans
}

func flips(a, b []int, first byte) int {
	// a[0] 是偶数位上面0的数量，a[1] 是偶数位上面1的数量
	// b[0] 是奇数位上面0的数量，b[1] 是奇数位上面1的数量
	var ans int

	if first == '1' {
		ans = a[0] + b[1]
	} else {
		ans = a[1] + b[0]
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
