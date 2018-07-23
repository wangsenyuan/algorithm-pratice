package p805

func splitArraySameAverage(A []int) bool {
	n := len(A)
	var sum int
	for i := 0; i < n; i++ {
		sum += A[i]
	}
	if sum == 0 {
		return true
	}

	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, sum+5)
	}

	f[0][0] = true
	var s int
	for i := 0; i < n; i++ {
		s += A[i]
		for cnt := i; cnt >= 0; cnt-- {
			for j := s - A[i]; j >= 0; j-- {
				f[cnt+1][j+A[i]] = f[cnt+1][j+A[i]] || f[cnt][j]
			}
		}
	}

	for i := 1; i < n; i++ {
		if sum*i%n == 0 && f[i][sum*i/n] {
			return true
		}
	}

	return false
}
