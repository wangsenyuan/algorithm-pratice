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

func solve(N int64, Q int, queries [][]int64) []int64 {
	// at most 64
	p := make([]int64, 64)
	np := make([]int, 64)
	var ps int
	var cnt int

	for N&1 == 0 {
		cnt++
		N >>= 1
	}
	if cnt > 0 {
		p[ps] = 2
		np[ps] = cnt
		ps++
	}

	for i := int64(3); i*i <= N; i += 2 {
		cnt = 0
		for N%i == 0 {
			cnt++
			N /= i
		}
		if cnt > 0 {
			p[ps] = i
			np[ps] = cnt
			ps++
		}
	}
	if N > 1 {
		p[ps] = N
		np[ps] = 1
		ps++
	}

	var X int64 = 1
	for i := 0; i < ps; i++ {
		X *= int64(np[i] + 1)
	}

	res := make([]int64, Q)

	kp := make([]int, 64)

	for i := 0; i < Q; i++ {
		query := queries[i]
		T := query[0]
		K := query[1]

		for j := 0; j < ps; j++ {
			kp[j] = 0
			for kp[j] < np[j] && K%p[j] == 0 {
				kp[j]++
				K /= p[j]
			}
		}
		if T == 1 {
			res[i] = 1
			for j := 0; j < ps; j++ {
				res[i] *= int64(kp[j] + 1)
			}
			continue
		}

		var ans int64
		if K == 1 {
			// K divides N
			ans = 1
			for j := 0; j < ps; j++ {
				ans *= int64(np[j] - kp[j] + 1)
			}
		}

		if T == 3 {
			ans = X - ans
		}
		res[i] = ans
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
