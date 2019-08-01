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
		R, C := readTwoNums(scanner)
		grid := make([]string, R)
		for i := 0; i < R; i++ {
			scanner.Scan()
			grid[i] = scanner.Text()
		}
		fmt.Println(solve(R, C, grid))
	}
}

func solve(R, C int, grid []string) int {
	var res int

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] != '#' {
				var can = (1<<4 - 1)
				for k := 1; can > 0; k++ {
					if can&1 == 1 && (i-k < 0 || grid[i-k][j] == '#') {
						can ^= 1
					}
					if can&2 == 2 && (j-k < 0 || grid[i][j-k] == '#') {
						can ^= 2
					}
					if can&4 == 4 && (i+k >= R || grid[i+k][j] == '#') {
						can ^= 4
					}
					if can&8 == 8 && (j+k >= C || grid[i][j+k] == '#') {
						can ^= 8
					}
					var cnt int
					if can&1 == 1 && grid[i-k][j] == 'D' {
						cnt++
					}
					if can&2 == 2 && grid[i][j-k] == 'R' {
						cnt++
					}
					if can&4 == 4 && grid[i+k][j] == 'U' {
						cnt++
					}
					if can&8 == 8 && grid[i][j+k] == 'L' {
						cnt++
					}
					res += cnt * (cnt - 1) / 2
				}
			}
		}
	}

	return res
}
