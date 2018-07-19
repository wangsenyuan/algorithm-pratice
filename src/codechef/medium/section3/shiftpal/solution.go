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

	for tc > 0 {
		scanner.Scan()
		s := scanner.Text()
		fmt.Println(solve(s))
		tc--
	}
}

func solve(s string) int {
	m := len(s)
	s += s
	n := len(s)
	N := 2*n + 1
	S := make([]byte, N)
	// sentinal node
	for i := 0; i < n; i++ {
		S[i<<1] = '|'
		S[i<<1|1] = s[i]
	}
	S[N-1] = '|'

	L := make([]int, N)

	C, R := 0, -1

	for i := 0; i < N; i++ {
		var rad int
		if i <= R {
			rad = min(L[2*C-i], R-i)
		}

		for i+rad < N && i-rad >= 0 && S[i+rad] == S[i-rad] {
			rad++
		}
		L[i] = rad
		if i+rad-1 > R {
			C = i
			R = i + rad - 1
		}
	}

	var ans int
	for i := m; i < N; i += 2 {
		if L[i] > m {
			ans++
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
