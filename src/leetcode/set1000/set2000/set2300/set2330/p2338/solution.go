package p2338

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b) % MOD
	return int(r)
}

const H = 16

func idealArrays(n int, maxValue int) int {
	dp := make([][]int, maxValue+1)
	for x := 1; x <= maxValue; x++ {
		dp[x] = make([]int, H+1)
		dp[x][1] = 1
	}
	for x := 1; x <= maxValue; x++ {
		for j := 1; j+1 <= H; j++ {
			for y := 2 * x; y <= maxValue; y += x {
				dp[y][j+1] = modAdd(dp[y][j+1], dp[x][j])
			}
		}
	}

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, H)
		C[i][0] = 1
		for j := 1; j <= i && j < H; j++ {
			C[i][j] = modAdd(C[i-1][j-1], C[i-1][j])
		}
	}

	var ans int

	for x := 1; x <= maxValue; x++ {
		for j := 1; j <= H; j++ {
			ans = modAdd(ans, modMul(C[n-1][j-1], dp[x][j]))
		}
	}

	return ans
}
