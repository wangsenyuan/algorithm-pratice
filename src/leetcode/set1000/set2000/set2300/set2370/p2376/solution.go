package p2376

func countSpecialNumbers(n int) int {
	arr := digits(n)
	l := len(arr)
	P := make([]int64, l)
	P[0] = 1
	for i := 1; i < l; i++ {
		P[i] = int64(i) * P[i-1]
	}
	C := make([][]int64, 11)
	for i := 0; i < 11; i++ {
		C[i] = make([]int64, 11)
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	var res int64

	// 长度为i且小于l的计数
	for i := 1; i < l; i++ {
		// most sig digit is 1...9
		// 从剩余的9个数字中选择i-1个，排列组合
		res += 9 * C[9][i-1] * P[i-1]
	}

	// l长度的满足条件的个数
	prev := make(map[int]bool)
	for i := l - 1; i >= 0 && len(prev)+1+i == l; i-- {
		var x int
		if i == l-1 {
			x++
		}

		for x < arr[i] {
			if !prev[x] {
				// not included yet
				res += C[10-len(prev)-1][i] * P[i]
			}
			x++
		}

		prev[arr[i]] = true
	}

	if len(prev) == l {
		res++
	}

	return int(res)
}

func digits(num int) []int {
	var res []int

	for num > 0 {
		res = append(res, num%10)
		num /= 10
	}

	// reverse(res)

	return res
}
