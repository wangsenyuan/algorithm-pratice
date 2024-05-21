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
		readNum(reader)
		s := readString(reader)
		res := solve(s)
		s = fmt.Sprintf("%v", res)

		buf.WriteString(fmt.Sprintf("%s\n", s[1:len(s)-1]))
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

func solve(s string) []int {
	n := len(s)
	// suf[i] 表示i的后面有多少个 < 向左的符号
	// 还需要知道前a个符号的下标和
	suf := make([]int, n+2)
	for i := n; i > 0; i-- {
		suf[i] = suf[i+1]
		if s[i-1] == '<' {
			suf[i]++
		}
	}
	
	suf[0] = suf[1]
	R := NewSegTree(suf[0])
	for i := 1; i <= n; i++ {
		if s[i-1] == '<' {
			R.Update(suf[0]-suf[i], i)
		}
	}

	pref := make([]int, n+2)
	for i := 1; i <= n; i++ {
		pref[i] = pref[i-1]
		if s[i-1] == '>' {
			pref[i]++
		}
	}
	L := NewSegTree(pref[n])

	for i := 1; i <= n; i++ {
		if s[i-1] == '>' {
			// 位置从0开始算
			L.Update(pref[i]-1, i)
		}
	}

	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if s[i-1] == '>' {
			// 先判断它从那边落地
			a := pref[i]
			b := suf[i]
			if b <= a-1 {
				// 从右边落地
				sum := R.Get(suf[0]-suf[i], suf[0]-suf[i]+b) * 2
				sum += n + 1
				sum -= L.Get(a-b-1, a-1) * 2
				sum -= i
				ans[i] = sum
			} else {
				// 从左边落地
				sum := R.Get(suf[0]-suf[i], suf[0]-suf[i]+a) * 2
				sum -= L.Get(0, a-1) * 2
				sum -= i
				ans[i] = sum
			}
		} else {
			// s[i] = <
			a := pref[i]
			b := suf[i]
			if a <= b-1 {
				// 从左边落地
				sum := R.Get(suf[0]-suf[i]+1, suf[0]-suf[i]+a+1) * 2
				sum -= L.Get(0, a) * 2
				sum += i
				ans[i] = sum
			} else {
				// 从右边落地
				sum := R.Get(suf[0]-suf[i]+1, suf[0]) * 2
				sum -= L.Get(pref[i]-b, pref[i]) * 2
				sum += n + 1
				sum += i
				ans[i] = sum
			}
		}
	}

	return ans[1:]
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] += v
	for p > 1 {
		t.arr[p>>1] = t.arr[p] + t.arr[p^1]
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	l += t.sz
	r += t.sz
	var res int
	for l < r {
		if l&1 == 1 {
			res += t.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += t.arr[r]
		}
		l >>= 1
		r >>= 1
	}

	return res
}
