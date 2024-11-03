package p3343

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(num ...int) int {
	res := 1
	for _, x := range num {
		res *= x
		res %= mod
	}
	return res
}

func countBalancedPermutations(num string) int {
	freq := make([]int, 10)
	var sum int
	for _, x := range []byte(num) {
		a := int(x - '0')
		sum += a
		freq[a]++
	}

	if sum%2 != 0 {
		return 0
	}
	sum /= 2

	n := len(num)

	C := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]int, n+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j], C[i-1][j-1])
		}
	}

	n1 := n / 2
	// n2 := n - n1

	dp := make([][]int, n1+1)
	fp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, sum+1)
		fp[i] = make([]int, sum+1)
	}

	dp[0][0] = 1
	var tot int
	for x := 0; x <= 9; x++ {
		for i := 0; i <= tot && i <= n1; i++ {
			for s := 0; s <= sum; s++ {
				for j := 0; j <= freq[x] && i+j <= n1 && s+j*x <= sum; j++ {
					// 如果把j个分配给偶数下标
					// 那么剩下的 freq[x] - j个要被分配个奇数下标
					tmp := mul(C[i+j][j], C[freq[x]-j+tot-i][freq[x]-j], dp[i][s])
					fp[i+j][s+j*x] = add(fp[i+j][s+j*x], tmp)
				}
			}
		}
		tot += freq[x]

		for i := 0; i <= tot && i <= n1; i++ {
			for s := 0; s <= sum; s++ {
				dp[i][s] = fp[i][s]
				fp[i][s] = 0
			}
		}
	}

	return dp[n1][sum]
}
