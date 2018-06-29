package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		scanner.Scan()
		s := scanner.Bytes()
		m, p := readTwoNums(scanner)
		fmt.Println(solve(s, m, p))
		t--
	}
}

func solve(s []byte, m, p int) int {
	n := len(s)
	m %= n
	p %= n

	N := 2 * n

	S := make([]byte, N)
	copy(S, s)
	copy(S[n:], s)

	can := make([]bool, n+1)
	can[0] = true

	f := make([]int, N+1)
	var c int

	for i, j := 2, 1; i <= N; i, j = i+1, j+1 {
		for c > 0 && S[c] != S[j] {
			c = f[c]
		}
		if S[c] == S[j] {
			c++
		}
		f[i] = c
		if i > n && c == n {
			can[i-n] = true
		}

		if c == n {
			c = f[c]
		}
	}

	var ans, ind int

	for k := 0; k < n; k++ {
		ind -= m
		ans++
		if ind < 0 {
			ind += n
		}
		if can[ind] {
			break
		}
		ind -= p
		ans++
		if ind < 0 {
			ind += n
		}
		if can[ind] {
			break
		}
	}

	return ans
}
