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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
	}
}

const MOD = 1e9 + 7

func solve(n int, A []int) int {
	f := make([]int64, n)
	f[0] = 1
	var res int64
	for i := 0; i < n; i++ {
		if i > 0 {
			modAdd(&f[i], f[i-1])
		}
		if i >= 2 && A[i-2] == 2 {
			modAdd(&f[i], f[i-2])
		}

		if i >= 3 && A[i-3] == 2 && A[i-2] == 2 {
			// i - 3 -> i - 1 -> i - 2 -> i
			modAdd(&f[i], f[i-3])
		}
		modAdd(&res, f[i])
	}

	nxt := n
	for i := n - 3; i >= 0; i-- {
		if A[i+2] == 1 {
			// must end at i + 2, otherwise can't reach i + 1
			nxt = i + 2
		}
		if A[i] == 1 {
			// can't go after i + 1
			continue
		}

		var g int64
		if nxt == n {
			g = int64(nxt - i - 2)
		} else {
			x := int64(nxt - i - 1)
			g = max(0, x-1)
			if x%2 == 1 {
				g++
			}
			if x%2 == 1 && nxt+1 < n && A[nxt+1] == 2 {
				g++
			}
		}
		tmp := g * f[i]
		tmp %= MOD
		modAdd(&res, tmp)
	}

	return int(res)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func modAdd(res *int64, num int64) {
	*res += num
	if *res >= MOD {
		*res -= MOD
	}
}
