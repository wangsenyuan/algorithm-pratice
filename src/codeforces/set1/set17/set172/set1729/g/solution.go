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
		t := readString(reader)
		mv, seq := solve(s, t)
		buf.WriteString(fmt.Sprintf("%d %d\n", mv, seq))
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

const MOD = 1_000_000_000 + 7

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func solve(s string, t string) (int, int) {
	n := len(s)
	m := len(t)
	var p []int
	p = append(p, -n)
	for i := 0; i < n; i++ {
		ok := i+m <= n
		for j := 0; j < m && ok; j++ {
			ok = s[i+j] == t[j]
		}
		if ok {
			p = append(p, i+1)
		}
	}
	p = append(p, n+m)
	tot := len(p) - 1

	f := make([]int, len(p))
	for i := 0; i < len(f); i++ {
		f[i] = 10000
	}
	g := make([]int, len(p))

	f[0] = 0
	g[0] = 1

	for i := 0; i < tot; i++ {
		j := i + 1
		for j <= tot && p[j] <= p[i]+m-1 {
			j++
		}
		for k := j; k <= tot && p[j]+m-1 >= p[k]; k++ {
			if f[i]+1 < f[k] {
				f[k] = f[i] + 1
				g[k] = g[i]
			} else if f[i]+1 == f[k] {
				g[k] = add(g[k], g[i])
			}
		}
	}

	return f[tot] - 1, g[tot]
}
