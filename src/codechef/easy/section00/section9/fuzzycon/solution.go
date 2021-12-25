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
		A := readNNums(scanner, 6)
		fmt.Println(solve(A[0], A[1], A[2], A[3], A[4], A[5]))
	}
}

func solve(x, y, z, a, b, c int) int {
	if x == a && y == b && z == c {
		return 0
	}

	if x+y+z == 0 {
		return min3(solve(1, 0, 0, a, b, c), solve(0, 1, 0, a, b, c), solve(0, 0, 1, a, b, c)) + 1
	}

	if a+b+c == 0 {
		return min3(solve(1, 0, 0, x, y, z), solve(0, 1, 0, x, y, z), solve(0, 0, 1, x, y, z)) + 1
	}

	X := mask(x, y, z)
	Y := mask(a, b, c)

	return min2(dist(X, Y), dist(flip(X), Y))
}

func min2(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func mask(a, b, c int) int {
	a &= 1
	b &= 1
	c &= 1
	return (a << 2) | (b << 1) | c
}

func dist(X, Y int) int {
	X &= 7
	Y &= 7
	z := X ^ Y

	var res int

	for z > 0 {
		res += z & 1
		z >>= 1
	}

	return res
}

func flip(num int) int {
	return 7 & (^num)
}
