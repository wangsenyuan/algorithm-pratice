package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readOneNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := readOneNum(scanner)

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = readNNums(scanner, n)
	}
	fmt.Println(solve(n, grid))
}

const oo = 1 << 30

func solve(n int, grid [][]int) int {
	sr := make([][]int, n)
	sc := make([][]int, n)

	for i := 0; i < n; i++ {
		sr[i] = make([]int, n)
		sc[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := grid[i][j]
			sr[i][j] = x
			sc[j][i] = x
		}
	}
	maxR := -oo
	minC := oo
	for i := 0; i < n; i++ {
		sort.Ints(sr[i])
		sort.Ints(sc[i])
		maxR = max(maxR, sr[i][0])
		minC = min(minC, sc[i][n-1])
	}

	rowChangeMax := make([]int, n)
	colChangeMax := make([]int, n)

	check := func(rowChanges, colChanges int) bool {
		rMax := maxR
		if rowChanges > 0 {
			if rowChangeMax[rowChanges] == 0 {
				for i := 0; i < n; i++ {
					if sr[i][rowChanges] > rMax {
						rMax = sr[i][rowChanges]
					}
				}
				rowChangeMax[rowChanges] = rMax
			}
			rMax = rowChangeMax[rowChanges]
		}
		cMin := minC
		if colChanges > 0 {
			if colChangeMax[colChanges] == 0 {
				for i := 0; i < n; i++ {
					if sc[i][n-1-colChanges] < cMin {
						cMin = sc[i][n-1-colChanges]
					}
				}
				colChangeMax[colChanges] = cMin
			}
			cMin = colChangeMax[colChanges]
		}

		return rMax >= cMin
	}

	for changes := 0; changes < n; changes++ {
		for rowChanges := 0; rowChanges <= changes; rowChanges++ {
			colChanges := changes - rowChanges
			if check(rowChanges, colChanges) {
				return changes
			}
		}
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
