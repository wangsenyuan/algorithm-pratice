package maxmatrix

func getMaxMatrix(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	sum := make([][]int, m)

	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)

		for j := 0; j < n; j++ {
			sum[i][j] = matrix[i][j]
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
		}
	}

	num := make([]int, m)

	var best int = -(1 << 30)
	ans := make([]int, 4)
	for j := 0; j < n; j++ {
		for k := j; k < n; k++ {
			for i := 0; i < m; i++ {
				num[i] = sum[i][k]
				if j > 0 {
					num[i] -= sum[i][j-1]
				}
			}
			tmp, a, b := findBest(num)
			if tmp > best {
				best = tmp
				ans[0] = a
				ans[1] = j
				ans[2] = b
				ans[3] = k
			}
		}
	}
	return ans
}

func findBest(num []int) (int, int, int) {
	a, b := -1, -1
	best := -(1 << 30)
	var sum int

	for i := 0; i < len(num); i++ {
		sum += num[i]
		if sum > best {
			best = sum
			b = i
		}
		if sum < 0 {
			sum = 0
		}
	}
	sum = 0
	for i := b; i >= 0; i-- {
		sum += num[i]
		if sum == best {
			a = i
			break
		}
	}
	return best, a, b
}
