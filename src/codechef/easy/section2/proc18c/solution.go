package main

import "fmt"

func main() {
	var tc int
	fmt.Scanf("%d", &tc)

	for tc > 0 {
		tc--

		var L, R uint64
		fmt.Scanf("%d %d", &L, &R)
		fmt.Println(solve(L, R))
	}
}

const MOD = 1e9 + 7
const MAX_SUM = 180

var cop [][]bool
var cnt1 []uint64
var cnt2 []uint64
var cnt [][]uint64

func init() {
	cop = make([][]bool, MAX_SUM+1)
	for i := 0; i <= MAX_SUM; i++ {
		cop[i] = make([]bool, MAX_SUM+1)
	}

	for i := 1; i <= MAX_SUM; i++ {
		cop[1][i] = true
	}
	for i := 2; i <= MAX_SUM; i++ {
		for j := i + 1; j <= MAX_SUM; j++ {
			g := gcd(i, j)
			if g == 1 {
				cop[i][j] = true
				cop[j][i] = true
			}
		}
	}

	cnt = make([][]uint64, MAX_SUM+1)
	for i := 0; i <= MAX_SUM; i++ {
		cnt[i] = make([]uint64, 20)
	}

	for x := 0; x < 10; x++ {
		cnt[x][1] = 1
	}

	for k := 2; k <= 19; k++ {
		for x := 0; x <= MAX_SUM; x++ {
			for y := 0; y < 10 && y <= x; y++ {
				cnt[x][k] += cnt[x-y][k-1]
				if cnt[x][k] >= MOD {
					cnt[x][k] -= MOD
				}
			}
		}
	}

	cnt1 = make([]uint64, MAX_SUM+1)
	cnt2 = make([]uint64, MAX_SUM+1)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func digits(x uint64) []int {
	res := make([]int, 20)
	var i int
	for x > 0 {
		res[i] = int(x % 10)
		i++
		x /= 10
	}
	res = res[:i]
	for j, k := 0, i-1; j < k; j, k = j+1, k-1 {
		res[j], res[k] = res[k], res[j]
	}
	return res
}

func count(x int, C uint64) uint64 {
	ds := digits(C)
	sz := len(ds)
	var ans uint64
	for i := 0; i < sz && x >= 0; i++ {
		if i == sz-1 {
			if x <= ds[i] {
				ans++
			}
		} else {
			for j := 0; j < ds[i] && j <= x; j++ {
				ans = (ans + cnt[x-j][sz-1-i]) % MOD
			}
			x -= ds[i]
		}
	}
	return ans
}

func solve(L, R uint64) int {
	for i := 1; i <= MAX_SUM; i++ {
		cnt1[i] = 0
		cnt1[i] = count(i, L-1) % MOD
		cnt2[i] = 0
		cnt2[i] = count(i, R) % MOD
	}
	var ans uint64
	for x := 1; x <= MAX_SUM; x++ {
		for y := x + 1; y <= MAX_SUM; y++ {
			if cop[x][y] {
				a := sub(cnt2[x], cnt1[x])
				b := sub(cnt2[y], cnt1[y])
				ans = (ans + a*b) % MOD
			}
		}
	}

	return int(ans)
}

func sub(a, b uint64) uint64 {
	if a >= b {
		return a - b
	}
	return a + MOD - b
}
