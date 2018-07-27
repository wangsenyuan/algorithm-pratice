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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		line := readNNums(scanner, 3)
		n, k, p := line[0], line[1], line[2]
		res := solve(n, k, p)
		fmt.Printf("Case #%d: %d\n", i, res)
	}
}

const MOD = 30031

var help [][]int

func init() {
	N := 1 << 10

	help = make([][]int, N)
	for i := 0; i < N; i++ {
		help[i] = make([]int, N)
	}
}

func solve(n, k, p int) int {
	fmt.Fprintf(os.Stderr, "[debug] solve(%d %d %d)\n", n, k, p)
	K := 1 << uint(k)
	P := 1 << uint(p)

	M := make([][]int, P)
	for i := 0; i < P; i++ {
		M[i] = make([]int, P)
	}

	mp := make([]int, P)
	for i := 0; i < P; i++ {
		mp[i] = -1
	}
	var idx int
	for mask := K - 1; mask < P; mask++ {
		if bitcount(mask) == k {
			mp[mask] = idx
			idx++
		}
	}

	for mask := 1; mask < P; mask++ {
		if mask&1 == 1 && bitcount(mask) == k {
			newMask := mask >> 1
			for i := 0; i < p; i++ {
				if newMask&(1<<uint(i)) == 0 && mp[newMask|(1<<uint(i))] >= 0 {
					M[mp[newMask|(1<<uint(i))]][mp[mask]] = 1
				}
			}
		}
	}
	res := make([][]int, P)
	for i := 0; i < P; i++ {
		res[i] = make([]int, P)
	}
	for i := 0; i < P; i++ {
		if i&1 == 1 && bitcount(i) == k {
			res[mp[i]][mp[K-1]] = 1
		}
	}
	fmt.Fprintf(os.Stderr, "[debug] solve(%d %d %d) => starting pow matrix\n", n, k, p)

	todo := n - k
	for todo > 0 {
		if todo&1 == 1 {
			res = multiple(res, M, idx)
		}
		M = multiple(M, M, idx)
		fmt.Fprintf(os.Stderr, "[debug] solve(%d %d %d) pow at %d\n", n, k, p, todo)
		todo >>= 1
	}

	return res[mp[K-1]][mp[K-1]]
}

func bitcount(x int) int {
	var cnt int
	for x > 0 {
		cnt += x & 1
		x >>= 1
	}
	return cnt
}

func multiple(a, b [][]int, n int) [][]int {
	// n := len(a)
	// a, b, c has to be n * n matrix

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var res int
			for k := 0; k < n; k++ {
				res += (a[i][k] * b[k][j]) % MOD
				if res >= MOD {
					res -= MOD
				}
			}
			help[i][j] = res
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = help[i][j]
		}
	}

	return a
}
