package p5550

func decrypt(code []int, k int) []int {
	if k == 0 {
		for i := 0; i < len(code); i++ {
			code[i] = 0
		}
		return code
	}

	if k < 0 {
		reverse(code)
		res := decrypt(code, -k)
		reverse(res)
		return res
	}

	n := len(code)
	sum := make([]int, n)
	for i := 0; i < n; i++ {
		sum[i] = code[i]
		if i > 0 {
			sum[i] += sum[i-1]
		}
	}

	for i := 0; i < n; i++ {
		j := i + k
		if j < n {
			code[i] = sum[j] - sum[i]
		} else {
			code[i] = sum[n-1] - sum[i] + sum[j-n]
		}
	}

	return code
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
