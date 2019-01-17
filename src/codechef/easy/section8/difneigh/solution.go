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
		n, m := readTwoNums(scanner)
		k, res := solve(n, m)
		fmt.Println(k)

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if j < m-1 {
					fmt.Printf("%d ", res[i][j])
				} else {
					fmt.Printf("%d\n", res[i][j])
				}
			}
		}
	}
}

var magic = [4][4]int{
	{1, 1, 4, 4}, {2, 3, 3, 2}, {4, 4, 1, 1}, {3, 2, 2, 3},
}

func solve(n, m int) (int, [][]int) {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	var k int
	if n == 1 || m == 1 {
		if max(n, m) <= 2 {
			k = 1
		} else {
			k = 2
		}
		if n == 1 {
			for i := 0; i < m; i++ {
				res[0][i] = 1 + (i%4)/2
			}
		} else {
			for i := 0; i < n; i++ {
				res[i][0] = 1 + (i%4)/2
			}
		}
	} else if min(n, m) == 2 {
		if max(n, m) == 2 {
			k = 2
		} else {
			k = 3
		}
		if n == 2 {
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					res[i][j] = 1 + j%3
				}
			}
		} else {
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					res[i][j] = 1 + i%3
				}
			}
		}
	} else {
		k = 4
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				res[i][j] = magic[i%4][j%4]
			}
		}
	}
	return k, res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
