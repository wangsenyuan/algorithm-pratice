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

func readNum(scanner *bufio.Scanner) (a int) {
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

	t := readNum(scanner)
	for t > 0 {
		n, m := readTwoNums(scanner)
		field := make([][]int, n)
		for i := 0; i < n; i++ {
			field[i] = readNNums(scanner, m)
		}
		fmt.Println(solve(n, m, field))
		t--
	}
}

func solve(n int, m int, field [][]int) int {
	cells := make(Cells, n*m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cells[i*m+j] = Cell{i, j, field[i][j]}
		}
	}
	sort.Sort(cells)

	checked := make([]bool, n*m)
	que := make([]int, n*m)

	dd := []int{-1, 0, 1, 0, -1}

	check := func(x, y int) {
		front, tail := 0, 0
		que[tail] = x*m + y
		tail++
		checked[x*m+y] = true
		for front < tail {
			tt := tail
			for front < tt {
				u, v := que[front]/m, que[front]%m
				front++
				for k := 0; k < 4; k++ {
					a, b := u+dd[k], v+dd[k+1]
					if a >= 0 && a < n && b >= 0 && b < m && !checked[a*m+b] && field[u][v] >= field[a][b] {
						checked[a*m+b] = true
						que[tail] = a*m + b
						tail++
					}
				}
			}
		}
	}

	var ans int
	for i := 0; i < len(cells); i++ {
		x, y := cells[i].x, cells[i].y
		if !checked[x*m+y] {
			ans++
			check(x, y)
		}
	}
	return ans
}

type Cell struct {
	x, y   int
	height int
}

type Cells []Cell

func (cells Cells) Len() int {
	return len(cells)
}

func (cells Cells) Less(i, j int) bool {
	return cells[i].height > cells[j].height
}

func (cells Cells) Swap(i, j int) {
	cells[i], cells[j] = cells[j], cells[i]
}
