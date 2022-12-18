package p2505

func smallestValue(n int) int {
	lpf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if lpf[i] == 0 {
			for j := i; j <= n; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	for lpf[n] != n {
		var sum int
		tmp := n
		for tmp > 1 {
			sum += lpf[tmp]
			tmp /= lpf[tmp]
		}
		if n == sum {
			break
		}
		n = sum
	}

	return n
}
