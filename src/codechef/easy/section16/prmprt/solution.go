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

	n := readNum(scanner)
	m := readNum(scanner)

	for i := 0; i < m; i++ {
		u, v := readTwoNums(scanner)
		fmt.Println(solve(n, u, v))
	}
}

const MAX_N = 1000007

var factors []int

func init() {
	factors = make([]int, MAX_N)

	set := make([]bool, MAX_N)

	for x := 2; x < MAX_N; x++ {
		if !set[x] {
			factors[x] = x
			Y := int64(x) * int64(x)
			if Y >= MAX_N {
				continue
			}
			y := int(Y)

			for y < MAX_N {
				if factors[y] == 0 {
					factors[y] = x
				}
				set[y] = true
				y += x
			}
		}
	}
}

func solve(n int, u, v int) int {
	if u == 1 || v == 1 {
		return -1
	}

	if u == v {
		return 0
	}

	if u%v == 0 && factors[v] == v {
		return 1
	}

	if v%u == 0 && factors[u] == u {
		return 1
	}

	uu := u

	for uu > 1 {
		x := factors[uu]
		if v%x == 0 {
			return 2
		}
		uu /= x
	}

	N := int64(n)

	x := factors[u]
	y := factors[v]
	M := int64(x) * int64(y)
	if M > N {
		return -1
	}
	r := 2

	if x != u {
		r++
	}
	if y != v {
		r++
	}
	return r
}
