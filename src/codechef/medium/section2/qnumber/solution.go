package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var N int64
	pos := readInt64(scanner.Bytes(), 0, &N)
	var Q int
	readInt(scanner.Bytes(), pos+1, &Q)
	queries := make([][]int64, Q)

	for i := 0; i < Q; i++ {
		queries[i] = make([]int64, 2)
		scanner.Scan()
		pos = readInt64(scanner.Bytes(), 0, &queries[i][0])
		readInt64(scanner.Bytes(), pos+1, &queries[i][1])
	}
	res := solve(N, Q, queries)

	for _, ans := range res {
		fmt.Printf("%d\n", ans)
	}
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

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

const MAX_N = 1000007

var primes []int

func init() {
	set := make([]bool, MAX_N+1)
	primes = make([]int, 0, 10000)
	for x := 2; x <= MAX_N; x++ {
		if !set[x] {
			primes = append(primes, x)
			for y := x * x; y > x && y <= MAX_N; y += x {
				set[y] = true
			}
		}
	}
}

func primeFactors(X int64) [][]int {
	res := make([][]int, 0, 32)
	for i := 0; i < len(primes) && X > int64(primes[i]); i++ {
		p := int64(primes[i])
		if X%p == 0 {
			var cnt int
			for X%p == 0 {
				cnt++
				X /= p
			}
			res = append(res, []int{primes[i], cnt})
		}
	}
	if X > 1 {
		res = append(res, []int{int(X), 1})
	}

	return res
}

func solve(N int64, Q int, queries [][]int64) []int64 {
	P := primeFactors(N)

	cal1 := func(S [][]int) int64 {
		var res int64 = 1
		for i, j := 0, 0; i < len(P) && j < len(S); {
			a, b := P[i], S[j]
			if a[0] < b[0] {
				i++
			} else if a[0] > b[0] {
				j++
			} else {
				x := min(a[1], b[1])
				res *= int64(x + 1)
				i++
				j++
			}
		}
		return res
	}

	cal2 := func(S [][]int) int64 {
		var res int64 = 1
		var i, j int
		for i < len(P) && j < len(S) {
			a, b := P[i], S[j]
			if a[0] < b[0] {
				// it is ok
				res *= int64(a[1] + 1)
				i++
			} else if a[0] > b[0] {
				// not ok, S has something P don't have
				return 0
			} else {
				if a[1] < b[1] {
					return 0
				}
				res *= int64(a[1] - b[1] + 1)
				i++
				j++
			}
		}
		if j < len(S) {
			return 0
		}
		if i < len(P) {
			res *= int64(P[i][1] + 1)
			i++
		}
		return res
	}

	var X int64 = 1
	for i := 0; i < len(P); i++ {
		X *= int64(P[i][1] + 1)
	}

	cal3 := func(S [][]int) int64 {
		Y := cal2(S)
		return X - Y
	}

	res := make([]int64, Q)

	for i := 0; i < Q; i++ {
		query := queries[i]
		S := primeFactors(query[1])
		if query[0] == 1 {
			res[i] = cal1(S)
		} else if query[0] == 2 {
			res[i] = cal2(S)
		} else {
			res[i] = cal3(S)
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
