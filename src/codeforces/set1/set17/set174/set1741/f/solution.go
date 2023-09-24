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
		segments := make([][]int, n)
		for i := 0; i < n; i++ {
			segments[i] = readNNums(reader, 3)
		}
		res := solve(segments)
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

const inf = 1000000009

func solve(segments [][]int) []int {
	n := len(segments)
	a := calc(segments)
	for i := 0; i < n; i++ {
		segments[i][0], segments[i][1] = inf-segments[i][1], inf-segments[i][0]
	}
	b := calc(segments)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = min(a[i], b[i])
	}
	return ans
}

func calc(segs [][]int) []int {
	n := len(segs)
	l := make([]Pair, n)
	r := make([]Pair, n)
	for i := 0; i < n; i++ {
		l[i] = Pair{segs[i][0], i}
		r[i] = Pair{segs[i][1], i}
	}
	sort.Sort(Pairs(l))
	sort.Sort(Pairs(r))
	suf := make([][]Pair, n)
	var cur []Pair
	for j := n - 1; j >= 0; j-- {
		//xr := r[j].first
		i := r[j].second
		xl := segs[i][0]
		cl := segs[i][2]
		if len(cur) == 0 {
			cur = append(cur, Pair{xl, cl})
		} else if len(cur) == 1 {
			if cur[0].second == cl {
				cur[0] = Pair{min(cur[0].first, xl), cl}
			} else {
				cur = append(cur, Pair{xl, cl})
			}
		} else {
			// if len(cur) == 2
			if cur[0].second == cl {
				cur[0] = Pair{min(cur[0].first, xl), cl}
			} else if cur[1].second == cl {
				cur[1] = Pair{min(cur[1].first, xl), cl}
			} else {
				cur = append(cur, Pair{xl, cl})
			}
		}

		sort.Sort(Pairs(cur))
		if len(cur) == 3 {
			cur = cur[:2]
		}
		tmp := make([]Pair, len(cur))
		copy(tmp, cur)
		suf[j] = tmp
	}

	ans := make([]int, n)

	for j, k := 0, 0; k < n; k++ {
		xl := l[k].first
		i := l[k].second
		ans[i] = inf
		xr := segs[i][1]
		cl := segs[i][2]
		for j < n && r[j].first < xl {
			j++
		}
		if j < n {
			s := suf[j]
			for _, x := range s {
				if x.second != cl {
					ans[i] = min(ans[i], max(0, x.first-xr))
				}
			}
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
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
