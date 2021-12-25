package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		n := readNum(scanner)

		A := make([][]int, n)

		for i := 0; i < n; i++ {
			scanner.Scan()
			var m int
			pos := readInt(scanner.Bytes(), 0, &m)
			A[i] = make([]int, m+1)
			A[i][0] = m
			for j := 0; j < m; j++ {
				pos = readInt(scanner.Bytes(), pos+1, &A[i][j+1])
			}
		}

		res := solve(n, A)

		fmt.Println(res)
	}

}

func solve(n int, A [][]int) int {
	mids := make([]int, n)
	var p int

	var earn int

	for i := 0; i < n; i++ {
		row := A[i]
		m := row[0]

		if m%2 == 0 {
			for i := 0; i*2 < m; i++ {
				earn += row[i+1]
			}
		} else {
			half := m / 2
			for i := 0; i < half; i++ {
				earn += row[i+1]
			}
			mids[p] = row[half+1]
			p++
		}
	}

	if p == 0 {
		return earn
	}

	mids = mids[:p]

	sort.Ints(mids)

	for i := p - 1; i >= 0; i -= 2 {
		earn += mids[i]
	}

	return earn
}
