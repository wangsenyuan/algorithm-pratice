package main

import "fmt"

func main() {
	var tc int

	fmt.Scanf("%d", &tc)

	for tc > 0 {
		tc--
		var num uint64
		fmt.Scanf("%d", &num)

		fmt.Println(solve(num))
	}
}

var C [][]int64

func init() {
	C = make([][]int64, 4)

	for i := 0; i < 4; i++ {
		C[i] = make([]int64, 4)
		C[i][0] = 1
		C[i][i] = 1

		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
		}
	}
}

func nCr(a, b int) int64 {
	if a < b {
		return 0
	}

	return C[a][b]
}

const MOD = int64(1e9 + 7)

func solve(num uint64) int {
	bits := ternary(num)

	n := len(bits)

	var get func(i int) (int64, int64)

	get = func(i int) (int64, int64) {
		if i == n {
			return 1, 0
		}

		x, y := get(i + 1)

		if bits[i] == 0 {
			return x, y
		}

		if bits[i] == 1 {
			return 2 * x, 2 * y
		}

		return 2*x + y, x + 2*y
	}

	x, y := get(0)

	res := x + 2*y

	res %= MOD

	return int(res)
}

func solve1(num uint64) int {
	bits := ternary(num)

	n := len(bits)

	mem := make([][][][]int64, n+1)

	for i := 0; i <= n; i++ {
		mem[i] = make([][][]int64, 2)
		for j := 0; j < 2; j++ {
			mem[i][j] = make([][]int64, 3)
			for k := 0; k < 3; k++ {
				mem[i][j][k] = make([]int64, 2)
				mem[i][j][k][0] = -1
				mem[i][j][k][1] = -1
				// mem[i][j][k][2] = -1
			}
		}
	}

	var loop func(i int, tight int, val int64, needed int) int64

	loop = func(i int, tight int, val int64, needed int) int64 {
		if mem[i][tight][val][needed-1] >= 0 {
			return mem[i][tight][val][needed-1]
		}

		var res int64

		if i == n {
			if int(val) == needed {
				res = 1
			}
		} else {
			if tight == 1 {
				for x := 0; x < bits[i]; x++ {
					res += loop(i+1, 0, val*nCr(bits[i], x)%3, needed)
				}
				res += loop(i+1, 1, val, needed)
			} else {
				for x := 0; x < 3; x++ {
					res += loop(i+1, 0, val*nCr(bits[i], x)%3, needed)
				}
			}
		}

		res %= MOD

		mem[i][tight][val][needed-1] = res

		return res
	}

	var res = loop(0, 1, 1, 1)
	res += loop(0, 1, 1, 2) * 2

	res %= MOD

	return int(res)
}

func ternary(num uint64) []int {
	res := make([]int, 0, 100)

	for num > 0 {
		res = append(res, int(num%3))
		num /= 3
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
