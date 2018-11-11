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
		line := readNNums(scanner, 4)
		fmt.Println(solve(line[0], line[1], line[2], line[3]))
	}
}

func solve(n, m int, x, y int) int64 {

	find := func(u, v int) int64 {
		//all cells after column v, except the same row
		res := int64(m-v) * int64(n-1)
		//right - left diagonal
		res -= int64(min(u-1, m-v))
		// left - right diagnao
		res -= int64(min(n-u, m-v))

		if u+v == x+y && y > v {
			res += int64(min(x-1, m-y))
		}

		if v-u == y-x && y > v {
			res += int64(min(n-x, m-y))
		}

		if u == x && y > v {
			res += int64(m - y)
		}
		// discount the (x, y) self
		if y > v && u+v != x+y && v-u != y-x && u != x {
			res--
		}
		return res
	}

	res := int64(x-1) * int64(n-x)

	for j := 1; j <= m; j++ {
		for i := 1; i <= n; i++ {
			if i == x && j == y {
				continue
			}
			res += find(i, j)
		}
	}

	return res * 2
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
