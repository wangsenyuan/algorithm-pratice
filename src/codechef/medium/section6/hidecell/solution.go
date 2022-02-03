package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(grid [][]byte) int {
		fmt.Println("?")
		for _, row := range grid {
			fmt.Printf("%s\n", string(row))
		}
		return readNum(reader)
	}

	tc := readNum(reader)
	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n, ask)
		fmt.Printf("!\n%d %d\n", res[0], res[1])
		ans := readNum(reader)
		if ans != 1 {
			panic("fail")
		}
	}
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n int, ask func([][]byte) int) []int {
	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]byte, n)
	}

	fillRow := func(r int, x byte) {
		for j := 0; j < n; j++ {
			grid[r][j] = x
		}
	}

	fill := func(x byte) {
		for i := 0; i < n; i++ {
			fillRow(i, x)
		}
	}

	isRow := func(r int) bool {
		fill('1')
		fillRow(r, '0')
		return ask(grid) == 1
	}

	getCol := func(a int) int {
		lo, hi := 1, n-2
		for lo < hi {
			c := (lo + hi) / 2
			fill('0')
			fillRow(n-1, '1')
			for i := 0; i < n-1; i++ {
				if i != a {
					for j := 0; j < c+1; j++ {
						grid[i][j] = '1'
					}
					for j := c + 1; j < n; j++ {
						grid[i][j] = '0'
					}
				}
			}
			if ask(grid) == 1 {
				hi = c
			} else {
				lo = c + 1
			}
		}

		return hi
	}

	getRow := func(c int) int {
		lo, hi := 1, n-2
		for lo < hi {
			r := (lo + hi) / 2
			for i := 0; i < n; i++ {
				if i <= r {
					for j := 0; j < c; j++ {
						grid[i][j] = '1'
					}
					grid[i][c] = '0'
					for j := c + 1; j < n; j++ {
						grid[i][j] = '1'
					}
				} else {
					for j := 0; j < n-1; j++ {
						grid[i][j] = '0'
					}
					grid[i][n-1] = '1'
				}
			}

			if ask(grid) == 1 {
				hi = r
			} else {
				lo = r + 1
			}
		}
		return hi
	}

	isOneOfCols := func(S []int) bool {
		in := make([]bool, n)
		for _, j := range S {
			in[j] = true
		}
		fill('0')
		for i := 0; i < n; i++ {
			if i == 0 || i == n-1 {
				fillRow(i, '1')
			} else {
				for j := 0; j < n; j++ {
					if j != n-1 && in[j+1] && i != n-2 {
						grid[i][j] = '1'
					} else if j != 0 && in[j-1] && i != 1 {
						grid[i][j] = '1'
					} else {
						grid[i][j] = '0'
					}
				}
			}
		}
		return ask(grid) == 1
	}

	if isRow(1) {
		j := getCol(1)
		return []int{1, j}
	}
	if isRow(n - 2) {
		j := getCol(n - 2)
		return []int{n - 2, j}
	}

	cols := make([][]int, 3)
	for j := 0; j < 3; j++ {
		cols[j] = make([]int, 0, n/3+1)
	}
	for j := 1; j < n-1; j++ {
		cols[j%3] = append(cols[j%3], j)
	}

	var id int
	if isOneOfCols(cols[1]) {
		id = 1
	} else if isOneOfCols(cols[2]) {
		id = 2
	}

	S := cols[id]

	for len(S) > 1 {
		l := len(S) / 2
		a, b := S[:l], S[l:]
		if isOneOfCols(a) {
			S = a
		} else {
			S = b
		}
	}

	return []int{getRow(S[0]), S[0]}
}
