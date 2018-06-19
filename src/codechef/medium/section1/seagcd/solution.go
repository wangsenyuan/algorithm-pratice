package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		tmp := readNNums(scanner, 4)
		fmt.Println(solve(tmp[0], tmp[1], tmp[2], tmp[3]))
		t--
	}
}

const MOD = 1000000007

var MAX_M = 10000000

var B = make([]uint64, MAX_M+1)

func solve(N, M, L, R int) uint64 {
	var ans uint64

	for D := M; D > 0; D-- {
		AD := fastPow(M/D, N)
		var S uint64

		for k := 2 * D; k <= M; k += D {
			S = (S + B[k])
			if S >= MOD {
				S -= MOD
			}
		}
		B[D] = AD + MOD - S
		if B[D] >= MOD {
			B[D] -= MOD
		}
		if D >= L && D <= R {
			ans = (ans + B[D])
			if ans >= MOD {
				ans -= MOD
			}
		}
	}
	return ans
}

func fastPow(a, n int) uint64 {
	var res uint64 = 1
	b := uint64(a)

	for n > 0 {
		if n&1 == 1 {
			res = (res * b)
			if res >= MOD {
				res %= MOD
			}
		}
		b = (b * b)
		if b >= MOD {
			b %= MOD
		}
		n >>= 1
	}

	return res
}
