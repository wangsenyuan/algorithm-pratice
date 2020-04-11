package main

import "fmt"

func main() {
	var tc int
	fmt.Scanf("%d", &tc)

	for tc > 0 {
		tc--

		var k int
		var n uint64

		fmt.Scanf("%d %d", &k, &n)

		fmt.Println(solve1(n, k))
	}
}

const MOD = 1000000009

func solve(num uint64, k int) int {
	mem := make([][]int64, 64)
	for i := 0; i < 64; i++ {
		mem[i] = make([]int64, 20)
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = -1
		}
	}

	var loop func(b int, carry int) int64

	loop = func(b int, carry int) int64 {
		if b < 0 {
			if carry == 0 {
				return 1
			}
			return 0
		}
		if carry >= len(mem[b]) {
			return 0
		}
		if mem[b][carry] >= 0 {
			return mem[b][carry]
		}
		mem[b][carry] = 0

		newCarry := carry + int((num>>uint(b))&1)

		for d := 0; d <= min(newCarry, k); d++ {
			mem[b][carry] += loop(b-1, 2*(newCarry-d))
			if mem[b][carry] >= MOD {
				mem[b][carry] -= MOD
			}
		}

		return mem[b][carry]
	}

	return int(loop(63, 0))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(n uint64, k int) int {
	if k == 3 {
		return solve3(n)
	}
	return solve4(n)
}

func solve3(n uint64) int {

	var f func(num uint64) int64

	mem := make(map[uint64]int64)

	f = func(num uint64) int64 {
		if num == 0 || num == 1 {
			return 1
		}
		if v, found := mem[num]; found {
			return v
		}

		res := f(num/2) + f(num/2-1)

		res %= MOD
		mem[num] = res
		return res
	}

	return int(f(n))
}

func solve4(n uint64) int {
	var f func(num uint64) int64

	mem := make(map[uint64]int64)

	f = func(num uint64) int64 {
		if num == 0 || num == 1 {
			return 1
		}

		if v, found := mem[num]; found {
			return v
		}

		res := f(num/2) + f(num/2-1)

		if num&1 == 0 {
			if num >= 4 {
				res += f(num/2 - 2)
			}
		}

		res %= MOD

		mem[num] = res

		return res
	}

	return int(f(n))
}
