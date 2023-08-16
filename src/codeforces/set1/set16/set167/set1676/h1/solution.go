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
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(A []int) int64 {
	// 对于i，需要知道i后面有多少个 a[j] <= a[i]的
	type Num struct {
		val int
		id  int
	}

	n := len(A)
	nums := make([]Num, n)

	for i := 0; i < n; i++ {
		nums[i] = Num{A[i], i}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i].val > nums[j].val || nums[i].val == nums[j].val && nums[i].id < nums[j].id
	})

	var ans int64

	t := NewSegTree(n)

	for i, num := range nums {
		// 在i的后面有 n - i - 1个数，
		// 其中有一部分排在它的前面
		cnt := n - 1 - i
		cnt -= t.Get(num.id+1, n)
		ans += int64(cnt)
		t.Add(num.id, 1)
	}

	return ans
}

type SegTree struct {
	arr []int
	n   int
}

func NewSegTree(n int) *SegTree {
	t := new(SegTree)
	t.arr = make([]int, 2*n)
	t.n = n
	return t
}

func (t *SegTree) Add(p int, v int) {
	for i := p + t.n; i > 0; i >>= 1 {
		t.arr[i] += v
	}
}

func (t *SegTree) Get(l int, r int) int {
	l += t.n
	r += t.n
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
