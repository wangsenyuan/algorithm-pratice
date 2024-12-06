package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	statements := make([]string, n)
	for i := range n {
		statements[i] = readString(reader)
	}
	return solve(statements)
}

func solve(statments []string) int {
	n := len(statments)

	g := make([][]int, n)
	// 每个只维护最多两个依赖
	pos := make(map[string]int)

	// 居然是计算haha的数量

	type state struct {
		cnt  int // haha count
		pref string
		suf  string
		ln   int
	}

	getPref := func(s string) string {
		n := len(s)
		return s[:min(n, 3)]
	}

	getSuf := func(s string) string {
		n := len(s)
		return s[max(0, n-3):]
	}

	count := func(s string) int {
		var cnt int
		for i := 0; i+4 <= len(s); i++ {
			if s[i:i+4] == "haha" {
				cnt++
			}
		}
		return cnt
	}

	merge := func(b, c state) state {
		cnt := b.cnt + c.cnt
		pref, suf := b.pref, c.suf

		cnt += count(b.suf + c.pref)

		if b.ln == len(b.suf) {
			pref = getPref(b.suf + c.pref)
		}

		if c.ln == len(c.pref) {
			suf = getSuf(b.suf + c.pref)
		}

		return state{cnt, pref, suf, b.ln + c.ln}
	}

	process := func(s string) state {
		return state{count(s), getPref(s), getSuf(s), len(s)}
	}

	not_processed := state{-1, "", "", -1}

	val := make([]state, n)

	for i, cur := range statments {
		j := strings.Index(cur, ":=")
		val[i] = not_processed
		if j < 0 {
			// a = b + c
			j = strings.Index(cur, "=")
			a := cur[:j-1]
			cur = cur[j+2:]
			k := strings.Index(cur, "+")
			b := pos[cur[:k-1]]
			c := pos[cur[k+2:]]
			pos[a] = i
			g[i] = append(g[i], b, c)
		} else {
			a := cur[:j-1]
			pos[a] = i
			val[i] = process(cur[j+3:])
		}
	}

	var dfs func(u int) state

	dfs = func(u int) state {
		if val[u] != not_processed {
			return val[u]
		}
		b, c := g[u][0], g[u][1]
		x := dfs(b)
		y := dfs(c)
		val[u] = merge(x, y)
		return val[u]
	}

	res := dfs(n - 1)

	return res.cnt
}
