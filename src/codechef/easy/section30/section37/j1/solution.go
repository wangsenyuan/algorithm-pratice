package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--

		board := make([]string, N)

		for i := 0; i < N; i++ {
			board[i] = readString(reader)
		}

		res := solve(board)

		for _, row := range res {
			buf.WriteString(row)
			buf.WriteByte('\n')
		}

		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const N = 9

func solve(board []string) []string {
	grid := make([][]byte, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]byte, N)
		copy(grid[i], []byte(board[i]))
	}

	row := make([]int, N)
	col := make([]int, N)
	sub := make([]int, N)
	dia := make([]int, 2)

	var placeDia func(x int) bool

	var placeRev func(x int) bool

	var placeOther func(x int, y int) bool

	placeDia = func(x int) bool {
		if x == N {
			return placeRev(0)
		}
		if grid[x][x] != '.' {
			return placeDia(x + 1)
		}
		// x < N
		for i := 1; i <= 9; i++ {
			y := x
			if (dia[0]>>i)&1 == 0 && (row[x]>>i)&1 == 0 && (col[y]>>i)&1 == 0 && (sub[(x/3)*3+(y/3)]>>i)&1 == 0 {
				// i not picked yet
				grid[x][x] = byte('0' + i)
				dia[0] ^= (1 << i)
				row[x] ^= (1 << i)
				col[y] ^= (1 << i)
				sub[(x/3)*3+(y/3)] ^= (1 << i)
				if placeDia(x + 1) {
					return true
				}
				row[x] ^= (1 << i)
				col[y] ^= (1 << i)
				sub[(x/3)*3+(y/3)] ^= (1 << i)
				dia[0] ^= (1 << i)
				grid[x][x] = '.'
			}
		}
		return false
	}

	placeRev = func(x int) bool {
		if x == N {
			return placeOther(0, 0)
		}
		if grid[x][N-1-x] != '.' {
			return placeRev(x + 1)
		}
		for i := 1; i <= 9; i++ {
			y := N - 1 - x
			if (dia[1]>>i)&1 == 0 && (row[x]>>i)&1 == 0 && (col[y]>>i)&1 == 0 && (sub[(x/3)*3+(y/3)]>>i)&1 == 0 {
				grid[x][N-1-x] = byte('0' + i)
				dia[1] ^= (1 << i)
				row[x] ^= (1 << i)
				col[y] ^= (1 << i)
				sub[(x/3)*3+(y/3)] ^= (1 << i)
				if placeRev(x + 1) {
					return true
				}
				row[x] ^= (1 << i)
				col[y] ^= (1 << i)
				sub[(x/3)*3+(y/3)] ^= (1 << i)
				dia[1] ^= (1 << i)
				grid[x][N-1-x] = '.'
			}
		}

		return false
	}

	placeOther = func(x int, y int) bool {
		if x == N {
			return true
		}
		if y == N {
			return placeOther(x+1, 0)
		}
		if grid[x][y] != '.' {
			return placeOther(x, y+1)
		}
		for i := 1; i <= 9; i++ {
			if (row[x]>>i)&1 == 0 && (col[y]>>i)&1 == 0 && (sub[(x/3)*3+(y/3)]>>i)&1 == 0 {
				row[x] ^= (1 << i)
				col[y] ^= (1 << i)
				sub[(x/3)*3+(y/3)] ^= (1 << i)
				grid[x][y] = byte('0' + i)
				if placeOther(x, y+1) {
					return true
				}
				grid[x][y] = '.'
				row[x] ^= (1 << i)
				col[y] ^= (1 << i)
				sub[(x/3)*3+(y/3)] ^= (1 << i)
			}
		}
		return false
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] != '.' {
				a := int(grid[i][j] - '0')
				row[i] |= (1 << a)
				col[j] |= (1 << a)
				sub[(i/3)*3+(j/3)] |= (1 << a)
				if i == j {
					dia[0] |= (1 << a)
				}
				if i+j == N-1 {
					dia[1] |= (1 << a)
				}
			}
		}
	}

	placeDia(0)

	res := make([]string, N)
	for i := 0; i < N; i++ {
		res[i] = string(grid[i])
	}
	return res
}

func next(i, j int) (x int, y int) {
	x = i
	y = j + 1
	if y == N {
		x++
		y = 0
	}
	return x, y
}
