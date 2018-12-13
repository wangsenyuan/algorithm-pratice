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
		k, n := readTwoNums(scanner)

		if k >= 4 {
			for i := 0; i < k; i++ {
				scanner.Scan()
			}
			fmt.Println("yes")
		} else {
			coords := make([][]int, k)
			for i := 0; i < k; i++ {
				coords[i] = readNNums(scanner, 2)
			}
			res := solve(k, n, coords)
			if res {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}

	}
}

func solve(k int, n int, coords [][]int) bool {
	if k >= 4 {
		return true
	}

	if k == 0 {
		return false
	}

	if k == 1 {
		return corner(coords[0], n)
	}
	if k == 2 {
		if corner(coords[0], n) || corner(coords[1], n) {
			return true
		}

		return pair(coords, n)
	}

	// k == 3
	if corner(coords[0], n) || corner(coords[1], n) || corner(coords[2], n) {
		return true
	}

	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			a := coords[i]
			b := coords[j]
			if pair([][]int{a, b}, n) {
				return true
			}
		}
	}
	return false
}

func corner(point []int, n int) bool {
	x, y := point[0], point[1]
	if x == 0 && y == 0 {
		return true
	}

	if x == 0 && y == n-1 {
		return true
	}

	if x == n-1 && y == 0 {
		return true
	}

	return x == n-1 && y == n-1
}

func pair(coords [][]int, n int) bool {
	point0, point1 := coords[0], coords[1]
	a, b := point0[0], point0[1]
	c, d := point1[0], point1[1]
	if a == 0 && (c == n-1 || c == 0) {
		return true
	}

	if a == n-1 && (c == 0 || c == n-1) {
		return true
	}

	if b == 0 && (d == n-1 || d == 0) {
		return true
	}

	if b == n-1 && (d == 0 || d == n-1) {
		return true
	}

	return false
}

type Points [][]int

func (this Points) Len() int {
	return len(this)
}

func (this Points) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this Points) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
