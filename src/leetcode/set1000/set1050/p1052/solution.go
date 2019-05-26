package p1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)
	left := make([]int, n)

	for i := 0; i < n; i++ {
		left[i] = customers[i] * (1 - grumpy[i])
		if i > 0 {
			left[i] += left[i-1]
		}
	}

	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		right[i] = customers[i] * (1 - grumpy[i])
		if i < n-1 {
			right[i] += right[i+1]
		}
	}

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + customers[i]
	}

	var best int

	for i := 0; i+X <= n; i++ {
		j := i + X
		var a int
		if i > 0 {
			a = left[i-1]
		}
		var b = 0
		if j < n {
			b = right[j]
		}
		c := sum[j] - sum[i]

		if a+b+c > best {
			best = a + b + c
		}
	}

	return best
}
