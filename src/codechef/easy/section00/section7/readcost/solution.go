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

	var n, x int
	var m uint64

	for {
		scanner.Scan()
		pos := readInt(scanner.Bytes(), 0, &n)
		pos = readUint64(scanner.Bytes(), pos+1, &m)
		readInt(scanner.Bytes(), pos+1, &x)

		if n == 0 && m == 0 && x == 0 {
			break
		}
		res := solve(n, int64(m), x)
		fmt.Println(res)
	}
}

func solve(n int, m int64, x int) int64 {
	N := int64(n)
	X := int64(x)
	var res int64

	if m <= N {

		for i := 0; i < int(m); i++ {
			I := int64(i)
			res += (I*N + X) / m
		}
		return res
	}

	a := (X + (m-1)*N) / m

	for i := int64(1); i <= a; i++ {
		p1 := (i*m - X + N - 1) / N
		p2 := ((i+1)*m - X + N - 1) / N
		p1 = max(p1, 0)
		p2 = min(p2, m)
		if p2 > p1 {
			res += (p2 - p1) * i
		}
	}

	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
