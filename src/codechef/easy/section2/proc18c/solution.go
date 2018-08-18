package main

func main() {

}

const MOD = 1e9 + 7
const MAX_SUM = 162 // 18 * 9
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
		cnt[i] = make([]uint64, 19)
	}

	for x := 0; x < 10; x++ {
		cnt[x][0] = 1
	}

	for k := 1; k < 19; k++ {
		for x := 0; x <= MAX_SUM; x++ {
			for y := 0; y < 10; y++ {
				z := x - y
				if z < 0 {
					break
				}
				cnt[x][k] += cnt[z][k-1]
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
	res := make([]int, 18)
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
	L := len(ds)
	if L == 1 {
		if uint64(x) <= C {
			return cnt[x][L-1]
		}
		return 0
	}
	var ans uint64
	for y := 0; y <= ds[0]; y++ {
		z := x - y
		if z >= 0 {
			ans += cnt[z][L-2]
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
				cx1, cy1 := cnt1[x], cnt1[y]
				tmp1 := (cx1 * cy1) % MOD

				cx2, cy2 := cnt2[x], cnt2[y]
				tmp2 := (cx2 * cy2) % MOD

				tmp3 := (cx1 * cy2) % MOD
				tmp4 := (cx2 * cy1) % MOD

				tmp := sub(tmp2, tmp1)
				tmp = sub(tmp, tmp3)
				tmp = sub(tmp, tmp4)
				ans += tmp
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
