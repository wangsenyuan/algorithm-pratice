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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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
		var L, R uint64
		scanner.Scan()
		pos := readUint64(scanner.Bytes(), 0, &L)
		readUint64(scanner.Bytes(), pos+1, &R)
		A := readNNums(scanner, 10)
		fmt.Println(solve(L, R, A))
	}
}

var C [][]uint64
var P [][]uint64

func init() {
	C = make([][]uint64, 20)

	for i := 0; i < 20; i++ {
		C[i] = make([]uint64, 20)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	P = make([][]uint64, 11)
	for i := 1; i <= 10; i++ {
		P[i] = make([]uint64, 19)
		P[i][0] = 1
		for j := 1; j < 19; j++ {
			P[i][j] = uint64(i) * P[i][j-1]
		}
	}
}

func solve(L, R uint64, A []int) uint64 {
	return count(R, A) - count(L-1, A)
}

func count(X uint64, A []int) uint64 {
	if X == 0 {
		return 0
	}

	stack := make([]int, 20)
	var top int

	for X != 0 {
		stack[top] = int(X % 10)
		top++
		X /= 10
	}

	process := func(mask int) uint64 {
		B := make([]int, 10)
		for i := 0; i < 10; i++ {
			B[i] = A[i]
		}

		fixed := make([]int, 0, 10)

		for i := 0; i < 10; i++ {
			if mask&(1<<uint(i)) > 0 {
				fixed = append(fixed, i)
			}
		}

		mul := func(c int) uint64 {
			var res uint64 = 1

			for i := 0; i < len(fixed); i++ {
				if B[fixed[i]] < 0 {
					return 0
				}
				res *= C[c][B[fixed[i]]]
				c -= B[fixed[i]]
				if c < 0 {
					return 0
				}
			}
			res *= P[10-len(fixed)][c]
			return res
		}
		var res uint64
		for i := 1; i < top; i++ {
			// try number with digit length i ( < top), those are definitely less than X
			// the only requirement is not set 0 at the MSB
			for j := 1; j < 10; j++ {
				B[j]--
				res += mul(i - 1)
				B[j]++
			}
		}
		// try length top, from MSB
		for i := top; i > 0; i-- {
			var st = 0
			if i == top {
				st++
			}
			for st < stack[i-1] {
				B[st]--
				res += mul(i - 1)
				B[st]++
				st++
			}
			B[stack[i-1]]--
		}

		// mul(0) is the number just equal X
		res += mul(0)

		return res
	}

	var res uint64

	for mask := 0; mask < (1 << 10); mask++ {
		tmp := process(mask)
		var cnt int
		x := mask
		for x > 0 {
			cnt += x & 1
			x >>= 1
		}
		if cnt&1 == 1 {
			res -= tmp
		} else {
			res += tmp
		}
	}
	return res
}
