package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	v := readNNums(reader, n)
	c := readNNums(reader, n)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}

	res := solve(v, c, queries)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

const inf = 1 << 50

type Pair struct {
	first  int
	second int
}

func solve(v []int, c []int, queries [][]int) []int {
	n := len(v)
	for i := 0; i < n; i++ {
		c[i]--
	}

	dp := make([]int, n)
	get := func(a, b int) int {
		for i := 0; i < n; i++ {
			dp[i] = -inf
		}

		dp_max := []Pair{{-1, 0}, {-1, 0}}

		for i := 0; i < n; i++ {
			ans1 := dp_max[0].second
			if dp_max[0].first == c[i] {
				ans1 = dp_max[1].second
			}
			ans1 = max(ans1, 0)
			ans2 := dp[c[i]]

			ans := max(ans1+b*v[i], ans2+a*v[i])

			if ans > ans2 {
				dp[c[i]] = ans
				if ans >= dp_max[0].second {
					if c[i] != dp_max[0].first {
						dp_max[1] = dp_max[0]
					}
					dp_max[0] = Pair{c[i], ans}
				} else if ans >= dp_max[1].second {
					dp_max[1] = Pair{c[i], ans}
				}
			}

		}

		return max(0, dp_max[0].second)
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = get(cur[0], cur[1])
	}

	return ans
}

func solve1(v []int, c []int, queries [][]int) []int {
	n := len(v)
	for i := 0; i < n; i++ {
		c[i]--
	}

	// for color
	tr := NewSegTree(n, -inf)

	get := func(a, b int) int {
		for i := 0; i < n; i++ {
			tr.Update(i, -inf)
		}

		for i := 0; i < n; i++ {
			ans1 := max(tr.Query(0, c[i]), tr.Query(c[i]+1, n), 0)
			ans2 := tr.Query(c[i], c[i]+1)

			ans := max(ans1+b*v[i], ans2+a*v[i])
			if ans > ans2 {
				tr.Update(c[i], ans)
			}
		}

		return max(0, tr.Query(0, n))
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = get(cur[0], cur[1])
	}

	return ans
}

type SegTree struct {
	arr []int
	iv  int
	n   int
}

func NewSegTree(n int, iv int) *SegTree {
	arr := make([]int, n*2)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &SegTree{arr, iv, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.n
	t.arr[p] = v

	for p > 1 {
		t.arr[p>>1] = max(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Query(l int, r int) int {
	res := t.iv
	l += t.n
	r += t.n
	for l < r {
		if l&1 == 1 {
			res = max(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
