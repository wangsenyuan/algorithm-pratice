package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)

	queries := make([][]int, k)
	for i := 0; i < k; i++ {
		queries[i] = readNNums(reader, 2)
	}

	res := solve(n, m, queries)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(n int, m int, queries [][]int) []int {
	var ans int
	for y := 0; y < m; y++ {
		for k := 1; ; k++ {
			nx := k / 2
			ny := y + (k+1)/2
			if nx == n || ny == m {
				break
			}
			ans += k
		}
	}
	// at least two cells
	for x := 0; x < n; x++ {
		for k := 1; ; k++ {
			nx := x + (k+1)/2
			ny := k / 2
			if nx == n || ny == m {
				break
			}
			ans += k
		}
	}

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			grid[i][j] = 1
		}
	}

	// only one cell staircase
	ans += n * m

	res := make([]int, len(queries))

	for i, cur := range queries {
		x, y := cur[0], cur[1]
		x--
		y--
		for c := 0; c < 2; c++ {
			l, r := 1, 1
			for {
				nx := x + (l+c)/2
				ny := y + (l+1-c)/2
				if nx == n || ny == m || grid[nx][ny] == 0 {
					break
				}
				l++
			}
			for {
				nx := x - (r+1-c)/2
				ny := y - (r+c)/2
				if nx < 0 || ny < 0 || grid[nx][ny] == 0 {
					break
				}
				r++
			}

			if grid[x][y] == 0 {
				ans += l * r
			} else {
				ans -= l * r
			}
		}
		ans += grid[x][y]
		grid[x][y] ^= 1
		ans -= grid[x][y]
		res[i] = ans
	}

	return res
}
