package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	line := readNNums(scanner, 3)
	N, M, T := line[0], line[1], line[2]
	X := make([][]int, N)

	for i := 0; i < N; i++ {
		X[i] = readNNums(scanner, 3)
	}

	Y := make([][]int, M)
	for i := 0; i < M; i++ {
		Y[i] = readNNums(scanner, 3)
	}
	base := readNNums(scanner, 2)

	fmt.Println(solve(N, M, T, X, Y, base))
}

const MAX_D = 10000007

var min_prime_factors []int

func init() {
	min_prime_factors = make([]int, MAX_D)
	not_primes := make([]bool, MAX_D)

	for x := 2; x < MAX_D; x++ {
		if not_primes[x] {
			continue
		}

		min_prime_factors[x] = x

		Y := int64(x) * int64(x)

		if Y >= MAX_D {
			continue
		}

		for y := x * x; y < MAX_D; y += x {
			not_primes[y] = true
			if min_prime_factors[y] == 0 {
				min_prime_factors[y] = x
			}
		}
	}
}

func solve(N, M, T int, X [][]int, Y [][]int, base []int) int {
	conns := make([][]int, M)
	for i, y := range Y {
		a := distance(y, base)
		conns[i] = make([]int, 0, 10)
		for j, x := range X {
			b := distance(x, y)
			if a+b <= T {
				conns[i] = append(conns[i], j)
			}
		}
	}
	pp := make([]int, N)
	for i := 0; i < N; i++ {
		pp[i] = -1
	}

	var dfs func(u int, c int) bool

	vis := make([]int, N)

	dfs = func(u int, c int) bool {
		if vis[u] == c {
			return false
		}
		vis[u] = c
		for _, v := range conns[u] {
			if pp[v] == -1 || dfs(pp[v], c) {
				pp[v] = u
				return true
			}
		}
		return false
	}
	var res int

	for i := 0; i < M; i++ {
		if dfs(i, i+1) {
			res++
		}
	}

	return res
}

func distance(from []int, to []int) int {
	x, y, id := from[0], from[1], from[2]
	u, v := to[0], to[1]
	d := abs(x-u) + abs(y-v)

	factors := make([]int, 0, 20)
	for d > 1 {
		factors = append(factors, min_prime_factors[d])
		d = d / min_prime_factors[d]
	}

	sort.Ints(factors)

	if len(factors) == 0 {
		return id
	}

	i := sort.Search(len(factors), func(i int) bool {
		return factors[i] >= id
	})
	if i == len(factors) {
		return id
	}
	return factors[i]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
