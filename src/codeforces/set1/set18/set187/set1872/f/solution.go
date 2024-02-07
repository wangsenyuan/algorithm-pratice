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
		n := readNum(reader)
		a := readNNums(reader, n)
		c := readNNums(reader, n)
		res := solve(c, a)
		s := fmt.Sprintf("%v", res)
		s = s[1 : len(s)-1]
		buf.WriteString(fmt.Sprintf("%s\n", s))
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

func solve(c []int, a []int) []int {
	n := len(a)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		deg[a[i]-1]++
	}

	vis := make([]bool, n)

	ans := make([]int, n)
	l, r := 0, n-1

	handle := func(cycle []int) {
		var it int
		m := len(cycle)
		for i := 0; i < m; i++ {
			if c[cycle[i]] < c[cycle[it]] {
				it = i
			}
		}

		// it 是最后一个
		for i := 0; i < m; i++ {
			j := (it - i + m) % m
			ans[r] = cycle[j] + 1
			r--
		}
	}

	que := make([]int, n)
	var front, tail int
	for u := 0; u < n; u++ {
		if deg[u] == 0 {
			vis[u] = true
			que[front] = u
			front++
		}
	}

	for tail < front {
		u := que[tail]
		tail++
		ans[l] = u + 1
		l++
		v := a[u] - 1
		deg[v]--
		if deg[v] == 0 {
			que[front] = v
			vis[v] = true
			front++
		}
	}

	// only cycles left
	for u := 0; u < n; u++ {
		if !vis[u] {
			vis[u] = true
			var pos int
			y := a[u] - 1
			for y != u {
				vis[y] = true
				que[pos] = y
				pos++
				y = a[y] - 1
			}
			que[pos] = y
			pos++
			handle(que[:pos])
		}
	}

	return ans
}
