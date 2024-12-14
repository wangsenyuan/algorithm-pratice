package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	res := solve(n)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(n int) [][]int {
	// n % 4 == 0
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	swap := func(i int, j int, u int, v int) {
		grid[i][j], grid[u][v] = grid[u][v], grid[i][j]
	}

	var v int
	for i := 0; i < n; i += 4 {
		for j := 0; j < n; j += 4 {
			for r := i; r < i+4; r++ {
				for c := j; c < j+4; c++ {
					grid[r][c] = v
					v++
				}
			}
			swap(i+1, j, i+1, j+1)
			swap(i+1, j+2, i+1, j+3)
			swap(i+2, j, i+2, j+2)
			swap(i+2, j+1, i+2, j+3)
			swap(i+3, j, i+3, j+3)
			swap(i+3, j+1, i+3, j+2)
		}

	}

	return grid
}
