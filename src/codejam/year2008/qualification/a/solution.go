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
	// f, err := os.Open("./A-small-practice.in")
	f, err := os.Open("./A-large-practice.in")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		n := readNum(scanner)
		S := make([]string, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			S[j] = scanner.Text()
		}

		m := readNum(scanner)
		Q := make([]string, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			Q[j] = scanner.Text()
		}

		fmt.Printf("Case #%d: %d\n", i, solve(S, Q))
	}
}

func solve(S []string, Q []string) int {
	si := make(map[string]bool)
	for i := 0; i < len(S); i++ {
		si[S[i]] = true
	}
	m := len(S)
	n := len(Q)

	pos := make(map[string]bool)
	var ans int
	for i := 0; i < n; i++ {
		q := Q[i]
		if si[q] && !pos[q] {
			pos[q] = true
		}
		if len(pos) == m {
			// can pick search
			ans++
			pos = make(map[string]bool)
			pos[q] = true
		}
	}

	return ans
}
