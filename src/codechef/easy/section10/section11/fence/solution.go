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
		line := readNNums(scanner, 3)
		n, m, k := line[0], line[1], line[2]
		plants := make([][]int, k)
		for i := 0; i < k; i++ {
			plants[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, m, k, plants))
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(n, m, k int, plants [][]int) int {
	res := k * 4

	mm := make(map[int]bool)

	for _, plant := range plants {
		r, c := plant[0] - 1, plant[1] - 1
		x := r * m + c
		mm[x] = true
	}


	for _, plant := range plants {
		r, c := plant[0] - 1, plant[1] - 1

		for i := 0; i < 4; i++ {
			u, v := r + dd[i], c + dd[i+1]
			if u >= 0 && u < n && v >= 0 && v < m {
				y := u * m + v
				if mm[y] {
					res--
				}
			}
		}
	}

	return res
}