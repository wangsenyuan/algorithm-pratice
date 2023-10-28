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
		n, m, d := readThreeNums(reader)
		p := readNNums(reader, n)
		a := readNNums(reader, m)
		res := solve(p, a, d)
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

func solve(p []int, a []int, d int) int {
	// pos[a[0]] < pos[a[1]] <= pos[a[0]] + d
	// 如果能够将pos[a[0]] 往前移动，或者 last(a] 往后移动
	// 只需要考虑相邻的两个数即可
	n := len(p)
	pos := make([]int, n+1)
	for i := 0; i < n; i++ {
		pos[p[i]] = i
	}

	best := n

	for i := 0; i+1 < len(a); i++ {
		// move a[i] left
		if pos[a[i]] > pos[a[i+1]] || pos[a[i+1]] > pos[a[i]]+d {
			return 0
		}
		// move a[i] left, or move a[i+1] right
		// or move a[i] right after a[i+1]
		tmp := pos[a[i+1]] - pos[a[i]]
		x := max(0, pos[a[i+1]]-d-1)
		y := min(n-1, x+d+1)

		if x <= pos[a[i]] && y >= pos[a[i+1]] && y-x > d {
			tmp = min(tmp, pos[a[i]]-x+y-pos[a[i+1]])
		}

		best = min(best, tmp)
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
