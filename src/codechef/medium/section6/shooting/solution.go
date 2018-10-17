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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			grid[i] = scanner.Text()
		}
		res := solve(n, m, grid)
		if res {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}
	}
}

func solve(n, m int, grid []string) bool {
	lasers := make([]int, 0, 16)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'L' {
				lasers = append(lasers, i*m+j)
			}
		}
	}
	l := len(lasers)

	N := 1 << uint(l)

	col := make([]int, m)
	row := make([]int, n)
	check := func(state int) bool {
		for i := 0; i < m; i++ {
			col[i] = -1
		}

		for i := 0; i < n; i++ {
			row[i] = -1
		}

		for i := 0; i < l; i++ {
			pos := lasers[i]
			x, y := pos/m, pos%m
			if state&(1<<uint(i)) > 0 {
				// vertical laser
				col[y] = max(col[y], x)
			} else {
				if row[x] < 0 {
					row[x] = y
				} else if row[x] < 50 {
					row[x] = 50
				}
			}
		}

		for i := 0; i < n; i++ {
			if row[i] == 50 {
				// two horizontal lasers
				continue
			} else if row[i] < 0 {
				// no horizontal lasers
				for j := 0; j < m; j++ {
					if grid[i][j] == 'E' && col[j] < i {
						return false
					}
				}
			} else {
				y := row[i]
				killRow := true
				// laser at (i, y) kill right
				for j := 0; j < y; j++ {
					if grid[i][j] == 'E' && col[j] < i {
						killRow = false
						break
					}
				}
				if !killRow {
					// laser kill left
					for j := y + 1; j < m; j++ {
						if grid[i][j] == 'E' && col[j] < i {
							return false
						}
					}
				}
			}
		}
		return true
	}

	for state := 0; state < N; state++ {
		if check(state) {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
