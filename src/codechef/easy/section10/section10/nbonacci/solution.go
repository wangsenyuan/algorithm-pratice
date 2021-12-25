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

	N, Q := readTwoNums(scanner)
	F := readNNums(scanner, N)
	K := make([]int, Q)
	for i := 0; i < Q; i++ {
		K[i] = readNum(scanner)
	}
	res := solve(N, Q, F, K)
	for _, num := range res {
		fmt.Println(num)
	}
}

func solve(N int, Q int, F []int, K []int) []int {
	// F[i] = F[i-1] ^ F[i-2] ^ F[i-3] .... F[i-N]
	// F[i + 1] = F[i] ^ F[i-1] ^ F[i-2] ... F[i-N + 1]
	// F[i] ^ F[i+1] = F[i] ^ F[i-N]
	// F[i+1] = F[i-N]
	// F[i] = F[i- N - 1]
	// F[N + 1] = F[0]
	// F[N+ 2] = F[1]
	// S[i-1] = F[0] ^ F[1] .... F[i - 1]
	// S[N] = F[0] ^ F[1] .... F[N] = 0
	// S[N+1] = F[0] ^ F[1] .... F[N] ^ F[N+1] = F[N+1] = F[0] = S[0]
	// S[N+2] = F[N+1] ^ F[N+2] = F[0] ^ F[1] = S[1]
	// S[N+N] = S[N-1]
	// S[N+N+1] = S[N] = 0
	S := make([]int, N+1)

	for i := 0; i < N; i++ {
		S[i] = F[i]
		if i > 0 {
			S[i] ^= S[i-1]
		}
	}

	ans := make([]int, Q)
	for i := 0; i < Q; i++ {
		k := (K[i] - 1) % (N + 1)
		ans[i] = S[k]
	}

	return ans
}
