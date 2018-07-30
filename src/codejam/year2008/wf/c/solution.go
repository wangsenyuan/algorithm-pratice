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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		n, m := readTwoNums(scanner)
		G := make([][]int, n)
		for i := 0; i < n; i++ {
			G[i] = readNNums(scanner, m)
		}
		fmt.Printf("Case #%d: %d\n", x, solve(n, m, G))
	}
}

func solve(n, m int, G [][]int) int {
	sumRowsCentered := func(r int) int {
		var res int
		var i int
		if m%3 == 0 {
			i = 1
		}

		for i < m {
			res += G[r][i]
			i += 3
		}
		return res
	}

	var total int

	var row int
	if n%3 == 0 {
		row = 1
	}
	for row < n {
		total += sumRowsCentered(row)
		row += 3
	}

	h := n / 2
	var S int

	if h%3 == 1 {
		for i := h - 1; i >= 0; i -= 3 {
			S += sumRowsCentered(i)
		}
		for i := h + 1; i < n; i += 3 {
			S += sumRowsCentered(i)
		}
		return S - total
	}

	for i := h - 2; i >= 0; i -= 3 {
		S += sumRowsCentered(i)
	}
	for i := h + 2; i < n; i += 3 {
		S += sumRowsCentered(i)
	}
	return total - S
}
