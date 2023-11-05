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
		b := readNNums(reader, n)
		m := readNum(reader)
		x := readNNums(reader, m)
		res := solve(a, b, x)
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
func solve(a []int, b []int, x []int) bool {
	// 如果 b[i] > a[i] => false
	// b[i] <= a[i]
	// 假设以 [l...r] 为一段
	//   必须有一个x = max(b[l...r])
	//  假设这个位置是k， 然后两边处理

	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return false
		}
	}
	tr := NewSegTree(b)

	cnt := make(map[int]int)
	for _, num := range x {
		cnt[num]++
	}
	// 要处理l...r， 是否能让 a[l...r] == b[l...r]
	var dfs func(l int, r int) bool

	dfs = func(l int, r int) bool {
		wb := tr.Get(l, r+1)

		i := wb.second
		cut := a[i] > b[i]
		// else a[i] == b[i]
		// split the whole range
		var segs []Pair
		if l < i {
			segs = append(segs, Pair{l, i - 1})
		}
		pos := i
		for pos+1 <= r {
			wb = tr.Get(pos+1, r+1)
			if wb.first < b[i] {
				segs = append(segs, Pair{pos + 1, r})
				break
			}

			// cur.first == b[i], no need to cut, just process seg between it
			if pos+1 <= wb.second-1 {
				segs = append(segs, Pair{pos + 1, wb.second - 1})
			}
			pos = wb.second
			if a[pos] > b[i] {
				cut = true
			}
		}

		if cut {
			if cnt[b[i]] <= 0 {
				return false
			}
			cnt[b[i]]--
		}

		for _, s := range segs {
			if !dfs(s.first, s.second) {
				return false
			}
		}
		return true
	}

	return dfs(0, len(a)-1)
}

type Pair struct {
	first  int
	second int
}

func max_pair(a, b Pair) Pair {
	if a.first > b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

type SegTree struct {
	arr []Pair
	n   int
}

func NewSegTree(nums []int) *SegTree {
	tr := new(SegTree)

	n := len(nums)
	tr.arr = make([]Pair, 2*n)
	tr.n = n
	for i := n; i < 2*n; i++ {
		tr.arr[i] = Pair{nums[i-n], i - n}
	}
	for i := n - 1; i > 0; i-- {
		tr.arr[i] = max_pair(tr.arr[i*2], tr.arr[i*2+1])
	}

	return tr
}

func (tr *SegTree) Get(l int, r int) Pair {
	l += tr.n
	r += tr.n
	res := Pair{0, -1}
	for l < r {
		if l&1 == 1 {
			res = max_pair(res, tr.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max_pair(res, tr.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
