package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		h := readNNums(reader, n)
		p := readNNums(reader, n)
		res := solve(k, h, p)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(k int, h []int, p []int) bool {
	n := len(h)
	monsters := make([]Pair, n)
	for i := 0; i < n; i++ {
		monsters[i] = Pair{h[i], p[i]}
	}

	sort.Slice(monsters, func(i, j int) bool {
		return monsters[i].first < monsters[j].first
	})

	suf := make([]int, n+1)
	suf[n] = 1 << 30
	//suf[n-1] = monsters[n-1].second

	for i := n - 1; i >= 0; i-- {
		suf[i] = min(suf[i+1], monsters[i].second)
	}

	var damage int
	for i := 0; i < n && k > 0; {
		//j := i
		for i < n && damage+k >= monsters[i].first {
			i++
		}
		if i == n {
			return true
		}
		// k最大1e5
		damage += k
		k -= suf[i]
	}

	return false
}

type Pair struct {
	first  int
	second int
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
