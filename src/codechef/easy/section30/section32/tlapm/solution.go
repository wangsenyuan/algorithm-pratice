package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// hint(105)
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		line := readNNums(reader, 4)
		a, b, c, d := line[0], line[1], line[2], line[3]
		res := solve(a, b, c, d)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
		if s[i] == '\n' {
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

const MAX_N = 2010

var row [][]int64
var col [][]int64

func init() {
	grid := make([][]int, MAX_N)
	for i := 0; i < MAX_N; i++ {
		grid[i] = make([]int, MAX_N)
	}
	cur := 1
	for k := 0; k < MAX_N; k++ {
		for j := k; j >= 0; j-- {
			grid[k-j][j] = cur
			cur++
		}
	}
	col = make([][]int64, MAX_N/2)
	for j := 0; j < len(col); j++ {
		col[j] = make([]int64, MAX_N/2)
		for i := 0; i < len(col[j]); i++ {
			col[j][i] = int64(grid[i][j])
			if i > 0 {
				col[j][i] += col[j][i-1]
			}
		}
	}
	row = make([][]int64, MAX_N/2)
	for i := 0; i < len(row); i++ {
		row[i] = make([]int64, MAX_N/2)
		for j := 0; j < len(row[i]); j++ {
			row[i][j] = int64(grid[i][j])
			if j > 0 {
				row[i][j] += row[i][j-1]
			}
		}
	}
}

func solve(x1, y1, x2, y2 int) int64 {
	// seems better to go down, then right
	x1--
	y1--
	x2--
	y2--
	res := col[y1][x2]
	if x1 > 0 {
		res -= col[y1][x1-1]
	}
	res += row[x2][y2]
	res -= row[x2][y1]
	return res
}
