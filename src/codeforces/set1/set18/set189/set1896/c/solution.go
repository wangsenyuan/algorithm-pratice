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
		n, x := readTwoNums(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(x, a, b)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			s := fmt.Sprintf("%v", res)
			s = s[1 : len(s)-1]
			buf.WriteString(fmt.Sprintf("%s\n", s))
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

func solve(x int, a []int, b []int) []int {
	n := len(a)

	pa := make([]Pair, n)
	pb := make([]Pair, n)

	for i := 0; i < n; i++ {
		pa[i] = Pair{a[i], i}
		pb[i] = Pair{b[i], i}
	}

	sort.Slice(pa, func(i, j int) bool {
		return pa[i].first < pa[j].first
	})

	sort.Slice(pb, func(i, j int) bool {
		return pb[i].first < pb[j].first
	})
	res := make([]Pair, n)

	for i := 0; i < x; i++ {
		if pb[i].first >= pa[n-x+i].first {
			// no solution
			return nil
		}
		res[i] = Pair{pa[n-x+i].second, pb[i].second}
	}
	for i := x; i < n; i++ {
		if pb[i].first < pa[i-x].first {
			return nil
		}
		res[i] = Pair{pa[i-x].second, pb[i].second}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].first < res[j].first
	})

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = b[res[i].second]
	}

	return ans
}

type Pair struct {
	first  int
	second int
}
