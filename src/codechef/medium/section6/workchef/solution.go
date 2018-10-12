package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		bs := scanner.Bytes()
		var L, R uint64

		pos := readUint64(bs, 0, &L)
		pos = readUint64(bs, pos+1, &R)
		var K int
		readInt(bs, pos+1, &K)
		fmt.Println(solve(L, R, K))
	}
}

const LCM = 1520
const MAXDIG = 20

var memo [][][][]int64

func init() {
	memo = make([][][][]int64, MAXDIG)
	for i := 0; i < MAXDIG; i++ {
		memo[i] = make([][][]int64, 2)
		for j := 0; j < 2; j++ {
			memo[i][j] = make([][]int64, LCM)
			for k := 0; k < LCM; k++ {
				memo[i][j][k] = make([]int64, 1<<9+5)
			}
		}
	}
}

func resetMemo() {
	for i := 0; i < MAXDIG; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < LCM; k++ {
				for l := 0; l < 1<<9+5; l++ {
					memo[i][j][k][l] = -1
				}
			}
		}
	}
}

func solve(L, R uint64, K int) int64 {
	return solveN(R, K) - solveN(L-1, K)
}

func solveN(n uint64, K int) int64 {
	resetMemo()
	dig := make([]int, 0, 64)

	if n == 0 {
		dig = append(dig, 0)
	} else {
		for n != 0 {
			dig = append(dig, int(n%uint64(10)))
			n /= uint64(10)
		}

		for i, j := 0, len(dig)-1; i < j; i, j = i+1, j-1 {
			dig[i], dig[j] = dig[j], dig[i]
		}
	}

	var dp func(index int, tight int, rem int, mask int) int64

	dp = func(index int, tight int, rem int, mask int) int64 {
		if memo[index][tight][rem][mask] < 0 {
			var res int64

			if index == len(dig) {
				var cnt int
				for d := 1; d < 10; d++ {
					if mask&(1<<uint(d-1)) > 0 {
						if rem%d == 0 {
							cnt++
						}
					}
				}
				if cnt >= K {
					res = 1
				}
			} else {
				for d := 0; d < 10; d++ {
					if tight > 0 && d > dig[index] {
						continue
					}
					var newTight int
					if tight > 0 && d == dig[index] {
						newTight = 1
					}
					newRem := (rem*10 + d) % LCM
					newMask := mask
					if d != 0 {
						newMask |= 1 << uint(d-1)
					}
					res += dp(index+1, newTight, newRem, newMask)
				}
			}
			memo[index][tight][rem][mask] = res
		}

		return memo[index][tight][rem][mask]
	}

	ans := dp(0, 1, 0, 0)

	return ans
}
