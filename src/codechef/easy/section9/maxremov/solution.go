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
		N, K := readTwoNums(scanner)
		ops := make([][]int, N)
		for i := 0; i < N; i++ {
			ops[i] = readNNums(scanner, 2)
		}
		res := solve(N, K, ops)
		fmt.Println(res)
	}
}

const MAX_N = 1e5 + 10

var pk []int
var pk1 []int
var cnt []int

func init() {
	pk = make([]int, MAX_N)
	pk1 = make([]int, MAX_N)
	cnt = make([]int, MAX_N)
}

func solve(N int, K int, ops [][]int) int {
	for i := 0; i < MAX_N; i++ {
		pk[i] = 0
		pk1[i] = 0
		cnt[i] = 0
	}

	for _, op := range ops {
		x, y := op[0], op[1]
		cnt[x]++
		cnt[y+1]--
	}

	for i := 1; i < MAX_N; i++ {
		cnt[i] += cnt[i-1]
	}

	for i := 1; i < MAX_N; i++ {
		pk[i] += pk[i-1]
		pk1[i] += pk1[i-1]
		if cnt[i] == K+1 {
			pk1[i]++
		} else if cnt[i] == K {
			pk[i]++
		}
	}

	var best int

	for _, op := range ops {
		x, y := op[0], op[1]
		a := pk[x-1]
		b := pk1[y] - pk1[x-1]
		c := pk[MAX_N-1] - pk[y]
		if a+b+c > best {
			best = a + b + c
		}
	}

	return best
}
