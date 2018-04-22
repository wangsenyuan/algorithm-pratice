package p821

func shortestToChar(S string, C byte) []int {
	n := len(S)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = n
	}

	var i int
	for i < n && S[i] != C {
		i++
	}

	for i < n {
		res[i] = 0
		j := i - 1
		for j >= 0 && i-j < res[j] {
			res[j] = i - j
			j--
		}
		j = i + 1
		for j < n && S[j] != C {
			res[j] = j - i
			j++
		}

		i = j
	}

	return res
}
