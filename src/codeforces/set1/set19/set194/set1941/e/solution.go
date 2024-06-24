package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		line := readNNums(reader, 4)
		n, m, k, d := line[0], line[1], line[2], line[3]
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			grid[i] = readNNums(reader, m)
		}
		res := solve(grid, k, d)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

type Pair struct {
	first  int
	second int
}

func solve(grid [][]int, k int, d int) int {
	n := len(grid)
	m := len(grid[0])

	que := make([]Pair, m)
	calc := func(r int) int {
		row := grid[r]

		var front, tail int
		que[front] = Pair{1, 0}
		front++

		for i := 1; i < m; i++ {
			for tail < front && i-que[tail].second-1 > d {
				tail++
			}
			// tail < front 肯定是成立的
			cur := que[tail].first + row[i] + 1
			for front > tail && que[front-1].first >= cur {
				front--
			}
			que[front] = Pair{cur, i}
			front++
		}
		return que[front-1].first
	}

	row := make([]int, n)
	var sum int
	best := 1 << 60
	for i := 0; i < n; i++ {
		row[i] = calc(i)
		sum += row[i]
		if i >= k-1 {
			if i >= k {
				sum -= row[i-k]
			}
			best = min(best, sum)
		}
	}

	return best
}
