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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		queries := make([][]int, n)
		for i := 0; i < n; i++ {
			queries[i] = readNNums(scanner, 3)
		}
		res := solve(n, queries)

		for i := 0; i < n; i++ {
			if res[i] {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}

func solve(n int, queries [][]int) []bool {
	var x, y int

	ans := make([]bool, n)

	for i := 0; i < n; i++ {
		query := queries[i]
		t := query[0]
		a := query[1]
		b := query[2]

		if t == 1 {
			ans[i] = true
			x = a
			y = b
		} else {
			// check x -> a, y -> b
			// or x -> b, y -> a can both valid
			if a == b && a >= x && b >= y {
				ans[i] = true
				x = a
				y = b
			} else if (a >= x && b >= y) && (b >= x && a >= y) {
				ans[i] = false
			} else if (a >= x && b >= y) || (b >= x && a >= y) {
				ans[i] = true
				if a >= x && b >= y {
					x = a
					y = b
				} else {
					x = b
					y = a
				}
			}
		}
	}

	return ans
}
