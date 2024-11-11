package p3352

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func countKReducibleNumbers(s string, k int) int {
	n := len(s)

	f := make([]int, n+1)
	f[1] = 0
	for i := 2; i <= n; i++ {
		var cnt int
		for tmp := i; tmp > 0; tmp = tmp & (tmp - 1) {
			cnt++
		}
		f[i] = f[cnt] + 1
	}

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j], C[i-1][j-1])
		}
	}

	var res int

	var cnt int

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			// 当前位置放置0, 可以保证比s小
			for x := 0; x <= n-i-1; x++ {
				// 选择x个位置放置1
				if cnt+x == 0 {
					// 必须要有1， 否则无法简化到1
					continue
				}
				if f[x+cnt] < k {
					res = add(res, C[n-i-1][x])
				}
			}

			cnt++
		}
	}

	return res
}
