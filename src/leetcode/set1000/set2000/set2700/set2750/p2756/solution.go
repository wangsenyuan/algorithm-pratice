package p2756

func findPrimePairs(n int) [][]int {
	spf := make([]int, n+1)
	var primes []int
	for i := 2; i <= n; i++ {
		if spf[i] == 0 {
			spf[i] = i
			primes = append(primes, i)
			if n/i < i {
				continue
			}
			for j := i * i; j <= n; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}

	var res [][]int

	for _, x := range primes {
		y := n - x
		if x > y {
			break
		}
		if spf[y] == y {
			res = append(res, []int{x, y})
		}
	}

	return res
}
