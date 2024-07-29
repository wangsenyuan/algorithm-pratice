package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	num := readString(reader)

	res := solve(num)

	fmt.Println(res)
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

func solve(num string) string {
	n := len(num)

	if n&1 == 1 {
		// must be even length
		return last(n)
	}
	// n is even

	// 尽量的放置4，在无法放置4的时候，放置7
	// 假设到i为止，res和num都是一致的
	// 第i位时，要放置x(4, 7), 那么必须考察
	// 是否能放置x = 还有至少一个位置放置给x，且放置后，res >= num
	// 假设 res[i] = num[i] = 4 那么后面必须满足 放置 n/2 - cnt[1] 个7 + n/2 - cnt[0] 个4，res[i+1:] >= num[i+1:]
	// 	否则必须放置一个7，如果7都不可以，只能放置更大的数

	buf := []byte(num)
	cnt := make([]int, 2)

	fill := func(i int) {
		for j := i; j < n; j++ {
			if cnt[0] < n/2 {
				buf[j] = '4'
				cnt[0]++
			} else {
				buf[j] = '7'
				cnt[1]++
			}
		}
	}

	p4 := NewSegTree(n, n, min)
	p7 := NewSegTree(n, n, min)

	for i := n - 1; i >= 0; i-- {
		if num[i] != '4' {
			p4.Update(i, i)
		}
		if num[i] != '7' {
			p7.Update(i, i)
		}
	}
	check := func(i int) bool {
		// 是否能够使用剩余的 n/2-cnt[1]个7，和n/2-cnt[0]个4，组成一个不比num[i:]小的数
		// 那么就应该是右边最近的数不是7的数，这个数 < 7
		// 否则的话，这个数没法被降低成7
		if cnt[1] < n/2 {
			j := p7.Get(i, i+n/2-cnt[1])
			// a是最近的不能7的数
			if j < n {
				// 如果 j < n， 要么大于7, 要么小于7
				return num[j] < '7'
			}
			// j == n
			// 全部是7
			i += n/2 - cnt[1]
		}
		j := p4.Get(i, n)
		return j == n || num[j] < '4'
	}

	if !check(0) {
		return last(n + 1)
	}

	for i := 0; i < n; i++ {
		x := int(num[i] - '0')
		if x < 4 && cnt[0] < n/2 {
			fill(i)
			return string(buf)
		}
		if x < 7 {
			if x == 4 && cnt[0] < n/2 {
				// 是否可以放置4
				if check(i + 1) {
					buf[i] = '4'
					cnt[0]++
					continue
				}
			}
			// 肯定能否放置7
			buf[i] = '7'
			cnt[1]++
			fill(i + 1)
			return string(buf)
		}
		// 这里肯定能放置7，否则第一个检查就不能通过
		buf[i] = '7'
		cnt[1]++
	}

	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func last(n int) string {
	res := make([]byte, n+(n&1))
	for i := 0; i < len(res); i++ {
		if i < (n+1)/2 {
			res[i] = '4'
		} else {
			res[i] = '7'
		}
	}
	return string(res)
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = seg.op(v, seg.arr[p])
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
