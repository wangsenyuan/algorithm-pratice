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
		s := readString(reader)
		w, m := readTwoNums(reader)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 3)
		}
		res := solve(w, s, Q)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
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

func solve(w int, s string, Q [][]int) [][]int {
	rec := make([][]int, 9)
	for i := 0; i < 9; i++ {
		rec[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			rec[i][j] = -1
		}
	}
	ans := make([][][]int, 9)
	for i := 0; i < 9; i++ {
		ans[i] = make([][]int, 9)
		for j := 0; j < 9; j++ {
			ans[i][j] = make([]int, 2)
			for k := 0; k < 2; k++ {
				ans[i][j][k] = -1
			}
		}
	}
	n := len(s)
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		x := int(s[i]-'0') % 9
		suf[i] = (x + suf[i+1]) % 9
		if i+w <= n {
			cur := (suf[i] - suf[i+w] + 9) % 9
			rec[cur][1] = rec[cur][0]
			rec[cur][0] = i
		}
	}

	process := func(l1 int, x int) {
		for v := 0; v < 9; v++ {
			y := (x * v) % 9
			for u := 0; u < 9; u++ {
				l2 := rec[u][0]
				if l2 == l1 {
					l2 = rec[u][1]
				}
				if l2 < 0 {
					continue
				}
				z := (y + u) % 9
				ans[v][z][0] = l1 + 1
				ans[v][z][1] = l2 + 1
			}
		}
	}

	for l1 := n - w; l1 >= 0; l1-- {
		cur := (suf[l1] - suf[l1+w] + 9) % 9
		process(l1, cur)
	}

	res := make([][]int, len(Q))

	for i, q := range Q {
		l1, l2, k := q[0], q[1], q[2]
		l1--
		v := (suf[l1] - suf[l2] + 9) % 9
		res[i] = ans[v][k]
	}
	return res
}
