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
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(a []int) []int {
	n := len(a)
	res := make([]int, n+1)

	// 首先猜测inv = 0
	// x[1] 处开始 + 1
	// x[2] 处开始 + 2
	// 考虑a[n] 是要 + sum(1...n) = (1 + n) * n / 2
	// 考虑 a[n-1] 正常它也是要这么多的，如果 a[n-1] + sum > a[n]
	// 显然 a[n] 有独立的一次, a[n-1] + sum - i < a[n]
	// res[i] = n
	//sum := int64(1+n) * int64(n) / 2

	ps := make([]Pair, n-1)

	for i := 0; i+1 < n; i++ {
		ps[i] = Pair{a[i] - a[i+1], i + 1}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first > ps[j].first
	})

	j := n
	for _, p := range ps {
		if p.first < 0 {
			break
		}
		// p.first > 0
		res[j] = p.second + 1
		j--
	}

	for j > 0 {
		res[j] = 1
		j--
	}

	return res[1:]
}

type Pair struct {
	first  int
	second int
}
