package main

import (
	"bufio"
	"os"
	"fmt"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
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

func readPair(scanner *bufio.Scanner) (a int, b int) {
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
		N, K := readPair(scanner)
		res := solve(N, K)
		for i := 0; i < N; i++ {
			if i < N-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
		t--
	}
}

func solve(N int, K int) []int {
	var ms int

	for k := K; k > 0; k >>= 1 {
		ms++
	}
	ms--

	// g[0] ^ g[1] ^ .. g[n-1] = 1111
	g := make([]int, N)

	if N == 1 {
		g[0] = K
	} else if K == 1 {
		for i := 0; i < N; i++ {
			g[i] = 1
		}
	} else if K == 2 {
		g[0] = 2
		for i := 1; i < N; i++ {
			g[i] = 1
		}
	} else if K == 3 {
		if N%2 == 0 {
			g[0] = 2
		} else {
			g[0] = 3
		}
		for i := 1; i < N; i++ {
			g[i] = 1
		}
	} else if N%2 == 0 {
		g[0] = 1 << uint(ms)
		g[1] = g[0] - 1
		for i := 2; i < N; i++ {
			g[i] = 1
		}
	} else {
		g[0] = 1 << uint(ms)
		g[1] = g[0] - 2
		for i := 2; i < N; i++ {
			g[i] = 1
		}
	}

	return g
}
