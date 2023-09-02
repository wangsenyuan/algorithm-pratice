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

	t := readNum(reader)

	for t > 0 {
		t--
		line := readNNums(reader, 4)
		n, k, r, c := line[0], line[1], line[2], line[3]
		res := solve(n, k, r, c)
		for i := 0; i < n; i++ {
			buf.WriteString(res[i])
			buf.WriteByte('\n')
		}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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
func solve(n int, k int, row int, col int) []string {
	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			grid[i][j] = '.'
		}
	}
	row--
	col--

	id := make([]int, n)
	// 0, 1, 2, 3
	y := col % k
	for i, j := 0, row-y; i < n; i++ {
		id[i] = (j + n) % n
		j = (j + 1) % n
	}

	fill := func(a, b int) {
		var it int
		for i := 0; i < k; i++ {
			grid[a][b+it] = 'X'
			a = (a + 1) % n
			it++
		}
	}

	for i := 0; i < n/k; i++ {
		for j := 0; j < n/k; j++ {
			fill(id[i*k], j*k)
		}
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = string(grid[i])
	}
	return res
}
